//go:build windows

package inputcontrol

import (
	"syscall"
	"time"
	"unsafe"
)

// Déclaration des DLLs et fonctions Windows nécessaires
var (
	// Référence vers la DLL kernel32.dll qui contient les fonctions de gestion de la console
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	// Fonction pour récupérer le mode actuel de la console
	getConsoleMode = kernel32.NewProc("GetConsoleMode")
	// Fonction pour définir le mode de la console
	setConsoleMode = kernel32.NewProc("SetConsoleMode")
	// Fonction pour vider le buffer d'entrée de la console
	flushConsoleBuffer = kernel32.NewProc("FlushConsoleInputBufferW")
	// Fonction pour obtenir un handle standard (stdin, stdout, stderr)
	getStdHandle = kernel32.NewProc("GetStdHandle")
	// Fonction pour lire les événements d'entrée de la console
	readConsoleInput = kernel32.NewProc("ReadConsoleInputW")
	// Fonction pour obtenir le nombre d'événements en attente dans le buffer d'entrée
	getNumberOfEvents = kernel32.NewProc("GetNumberOfConsoleInputEventsW")
	// Variables globales pour sauvegarder l'état de la console
	originalConsoleMode uint32                 // Mode de console original à restaurer
	consoleModesSaved   bool           = false // Flag pour savoir si on a déjà sauvegardé le mode
	stdinHandle         syscall.Handle         // Handle vers l'entrée standard (stdin)
)

// Constantes Windows
const (
	// Constante pour obtenir le handle de l'entrée standard
	// ^uint32(10) est equivalent à -11 en complément à deux
	STD_INPUT_HANDLE = ^uint32(10)
)

// Point d'entrée unifié pour Windows - désactiver les entrées
func disableInputPlatform() error {
	return disableInputWindows()
}

// Point d'entrée unifié pour Windows - réactiver les entrées
func enableInputPlatform() error {
	return enableInputWindows()
}

// Fonction spécifique à Windows pour désactiver les entrées utilisateur
func disableInputWindows() error {
	// Protection contre les accès concurrents
	inputMutex.Lock()
	defer inputMutex.Unlock()
	// Si les entrées sont déjà désactivées, on ne fait rien
	if inputDisabled {
		return nil
	}
	// Obtient le handle de stdin (entrée standard)
	handle, _, err := getStdHandle.Call(uintptr(STD_INPUT_HANDLE))
	if handle == uintptr(syscall.InvalidHandle) {
		return err
	}
	stdinHandle = syscall.Handle(handle)
	// Sauvegarde le mode de console actuel (une seule fois)
	if !consoleModesSaved {
		// Appel de GetConsoleMode pour récupérer le mode actuel
		ret, _, _ := getConsoleMode.Call(
			uintptr(stdinHandle),
			uintptr(unsafe.Pointer(&originalConsoleMode)),
		)
		// Si l'appel échoue (ret == 0), on retourne l'erreur
		if ret == 0 {
			return syscall.GetLastError()
		}
		consoleModesSaved = true
	}
	// Vide complètement le buffer d'entrée pour éliminer toute entrée en attente
	flushConsoleBuffer.Call(uintptr(stdinHandle))
	// Désactive TOUS les modes d'input en mettant le mode à 0
	// Cela désactive : echo, line input, processed input, etc.
	newMode := uint32(0)
	ret, _, err := setConsoleMode.Call(uintptr(stdinHandle), uintptr(newMode))
	if ret == 0 {
		return err
	}
	// Marque les entrées comme désactivées
	inputDisabled = true
	// Lance une goroutine pour consommer les entrées en arrière-plan
	// Cela évite que le buffer se remplisse et bloque l'application
	go consumeInputsWindows()
	return nil
}

// Fonction spécifique à Windows pour réactiver les entrées utilisateur
func enableInputWindows() error {
	// Protection contre les accès concurrents
	inputMutex.Lock()
	defer inputMutex.Unlock()
	// Si les entrées ne sont pas désactivées ou si on n'a pas sauvegardé le mode,
	// on ne fait rien
	if !inputDisabled || !consoleModesSaved {
		return nil
	}
	// Vide le buffer avant de restaurer le mode normal
	// Cela évite que des entrées "fantômes" apparaissent
	flushConsoleBuffer.Call(uintptr(stdinHandle))
	// Restaure le mode de console original
	ret, _, err := setConsoleMode.Call(uintptr(stdinHandle), uintptr(originalConsoleMode))
	if ret == 0 {
		return err
	}
	// Marque les entrées comme réactivées
	inputDisabled = false
	return nil
}

// Goroutine qui consomme en continu les événements d'entrée
// Cette fonction tourne en arrière-plan tant que les entrées sont désactivées
func consumeInputsWindows() {
	// Buffer pour stocker les événements lus
	// Chaque événement INPUT_RECORD fait 20 bytes sur Windows
	inputBuffer := make([]byte, 256)
	var eventsRead uint32 // Nombre d'événements effectivement lus
	// Boucle tant que les entrées sont désactivées
	for IsInputDisabled() {
		var numberOfEvents uint32
		// Vérifie s'il y a des événements en attente dans le buffer
		getNumberOfEvents.Call(
			uintptr(stdinHandle),
			uintptr(unsafe.Pointer(&numberOfEvents)),
		)
		// S'il y a des événements, on les lit (et les "consomme")
		if numberOfEvents > 0 {
			readConsoleInput.Call(
				uintptr(stdinHandle),                     // Handle de stdin
				uintptr(unsafe.Pointer(&inputBuffer[0])), // Buffer de destination
				uintptr(len(inputBuffer)/20),             // Nombre max d'événements à lire
				uintptr(unsafe.Pointer(&eventsRead)),     // Nombre d'événements lus
			)
		}
		// Pause de 10ms pour éviter une utilisation CPU excessive
		// C'est un compromis entre réactivité et performance
		time.Sleep(10 * time.Millisecond)
	}
}
