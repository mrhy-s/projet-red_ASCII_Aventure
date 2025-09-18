package inputcontrol

import (
	"os"
	"sync"
	"syscall"
)

// Constantes pour les appels système ioctl() sous Linux
const (
	// IOCTL_TCFLSH est la constante système pour l'opération TCFLSH
	// Elle permet de vider les buffers du terminal (entrée/sortie)
	// Valeur hexadécimale : 0x540B (correspond à TCFLSH sous Linux)
	IOCTL_TCFLSH = 0x540B

	// TCIFLUSH_OPERATION indique qu'on veut vider uniquement le buffer d'ENTRÉE
	// 0 = TCIFLUSH - vide le buffer d'entrée (ce que tape l'utilisateur)
	// 1 = TCOFLUSH - viderait le buffer de sortie (ce qui s'affiche)
	// 2 = TCIOFLUSH - viderait les deux buffers
	TCIFLUSH_OPERATION = 0
)

// ===== VARIABLES GLOBALES =====

var (
	// terminalInputBufferMutex protège l'accès concurrent aux opérations
	// de manipulation du buffer d'entrée du terminal.
	//
	// POURQUOI UN MUTEX ?
	// - Plusieurs goroutines pourraient appeler ClearInputBuffer() simultanément
	// - Les opérations ioctl() sur le même file descriptor peuvent entrer en conflit
	// - Le mutex garantit qu'une seule opération de nettoyage se fait à la fois
	terminalInputBufferMutex sync.Mutex
)

// ClearTerminalBuffer vide complètement le buffer d'entrée du terminal.
//
// QU'EST-CE QUE ÇA FAIT CONCRÈTEMENT ?
// Imagine que l'utilisateur tape très vite "abcdef" pendant qu'un programme
// traite autre chose. Ces caractères restent "en attente" dans un buffer système.
// Cette fonction supprime tous ces caractères en attente.

func ClearInputBuffer() error {
	// ÉTAPE 1 : Acquisition du verrou
	// On s'assure qu'une seule goroutine peut exécuter cette fonction à la fois
	terminalInputBufferMutex.Lock()
	// defer garantit que le verrou sera libéré même si la fonction panique
	defer terminalInputBufferMutex.Unlock()

	// ÉTAPE 2 : Obtention du file descriptor
	// os.Stdin représente l'entrée standard (généralement le clavier)
	// Fd() retourne le numéro du file descriptor Unix associé
	// Sous Linux, c'est généralement 0 pour stdin
	stdinFileDescriptor := int(os.Stdin.Fd())

	// ÉTAPE 3 : Appel système ioctl()
	// ioctl() est un appel système Unix qui permet de contrôler les périphériques
	//
	// PARAMÈTRES DE syscall.Syscall :
	// - syscall.SYS_IOCTL : numéro de l'appel système ioctl
	// - uintptr(stdinFileDescriptor) : le fichier à contrôler (stdin)
	// - IOCTL_TCFLSH : l'opération à effectuer (vider des buffers)
	// - TCIFLUSH_OPERATION : le type de vidage (buffer d'entrée uniquement)
	//
	// VALEURS DE RETOUR :
	// - syscallReturn1 : première valeur de retour (pas utilisée ici)
	// - syscallReturn2 : deuxième valeur de retour (pas utilisée ici)
	// - systemErrorCode : code d'erreur (0 = succès, autre = erreur)
	syscallReturn1, syscallReturn2, systemErrorCode := syscall.Syscall(
		syscall.SYS_IOCTL,            // Appel système : ioctl
		uintptr(stdinFileDescriptor), // File descriptor à manipuler
		IOCTL_TCFLSH,                 // Commande : vider les buffers du terminal
		TCIFLUSH_OPERATION,           // Paramètre : vider seulement le buffer d'entrée
	)

	// On ignore les valeurs de retour 1 et 2 car TCFLSH ne les utilise pas
	_ = syscallReturn1
	_ = syscallReturn2

	// ÉTAPE 4 : Vérification du succès
	// Si systemErrorCode != 0, il y a eu une erreur
	if systemErrorCode != 0 {
		return systemErrorCode
	}

	// ÉTAPE 5 : Succès
	return nil
}
