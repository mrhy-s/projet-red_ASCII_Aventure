//go:build linux || darwin || freebsd || netbsd || openbsd

package inputcontrol

import (
	"os"
	"syscall"
	"time"
	"unsafe"
)

// Variables globales pour sauvegarder l'état du terminal Unix
var (
	// Sauvegarde complète des attributs du terminal (structure termios)
	originalTermios []byte
	// Flag pour savoir si on a déjà sauvegardé les attributs du terminal
	termiosSaved bool = false
)

// Point d'entrée unifié pour les plateformes Unix - désactiver les entrées
func disableInputPlatform() error {
	return disableInputUnix()
}

// Point d'entrée unifié pour les plateformes Unix - réactiver les entrées
func enableInputPlatform() error {
	return enableInputUnix()
}

// Fonction spécifique aux systèmes Unix pour désactiver les entrées utilisateur
func disableInputUnix() error {
	// Protection contre les accès concurrents
	inputMutex.Lock()
	defer inputMutex.Unlock()

	// Si les entrées sont déjà désactivées, on ne fait rien
	if inputDisabled {
		return nil
	}

	// Récupère le file descriptor de stdin (entrée standard)
	fd := int(os.Stdin.Fd())

	// Sauvegarde les attributs actuels du terminal (une seule fois)
	if !termiosSaved {
		// Alloue un buffer pour la structure termios
		// 60 bytes est une approximation de la taille de la structure termios
		// sur la plupart des systèmes Unix
		termios := make([]byte, 60)
		// Appel système ioctl avec TCGETS pour récupérer les attributs du terminal
		_, _, errno := syscall.Syscall(
			syscall.SYS_IOCTL,                    // Numéro de syscall pour ioctl
			uintptr(fd),                          // File descriptor de stdin
			0x5401,                               // TCGETS - constante pour récupérer termios
			uintptr(unsafe.Pointer(&termios[0])), // Pointeur vers le buffer termios
		)
		// Si l'appel système échoue, retourner l'erreur
		if errno != 0 {
			return errno
		}
		// Fait une copie complète des attributs originaux pour la restauration
		originalTermios = make([]byte, len(termios))
		copy(originalTermios, termios)
		termiosSaved = true
	}
	// Vide le buffer d'entrée du terminal (TCFLSH avec TCIFLUSH)
	// Cela élimine toutes les entrées en attente
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		0x540B, // TCFLSH - constante pour vider les buffers
		0,      // TCIFLUSH - vider le buffer d'entrée
	)
	// Configure le terminal avec de nouveaux paramètres
	// Crée une copie des attributs originaux pour les modifier
	newTermios := make([]byte, len(originalTermios))
	copy(newTermios, originalTermios)
	// Modifie les local flags (lflag) dans la structure termios
	// Les local flags sont situés à l'offset 12 dans la structure termios
	lflag := (*uint32)(unsafe.Pointer(&newTermios[12]))
	// Désactive les flags ICANON (mode canonique) et ECHO
	// ICANON (0x02) : mode ligne par ligne -> mode caractère par caractère
	// ECHO (0x08) : affichage des caractères tapés
	// L'opération &^= fait un AND avec le complément (désactive les bits)
	*lflag &^= 0x0A // 0x0A = ICANON | ECHO (0x02 | 0x08)
	// Configure VMIN et VTIME pour un comportement non-bloquant
	// VMIN (index 52) : nombre minimum de caractères à lire (0 = non-bloquant)
	// VTIME (index 53) : timeout en dixièmes de seconde (0 = pas de timeout)
	newTermios[52] = 0 // VMIN = 0
	newTermios[53] = 0 // VTIME = 0
	// Applique les nouveaux attributs au terminal avec TCSETS
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		0x5402,                                  // TCSETS - constante pour définir termios
		uintptr(unsafe.Pointer(&newTermios[0])), // Pointeur vers les nouveaux attributs
	)
	// Marque les entrées comme désactivées
	inputDisabled = true
	// Lance une goroutine pour consommer les entrées en arrière-plan
	go consumeInputsUnix()
	return nil
}

// Fonction spécifique aux systèmes Unix pour réactiver les entrées utilisateur
func enableInputUnix() error {
	// Protection contre les accès concurrents
	inputMutex.Lock()
	defer inputMutex.Unlock()
	// Si les entrées ne sont pas désactivées ou si on n'a pas sauvegardé termios,
	// on ne fait rien
	if !inputDisabled || !termiosSaved {
		return nil
	}
	// Récupère le file descriptor de stdin
	fd := int(os.Stdin.Fd())
	// Vide le buffer d'entrée avant de restaurer les paramètres
	// Cela évite que des caractères "fantômes" apparaissent
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		0x540B, // TCFLSH
		0,      // TCIFLUSH
	)
	// Restaure les attributs originaux du terminal
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		0x5402, // TCSETS
		uintptr(unsafe.Pointer(&originalTermios[0])), // Pointeur vers les attributs originaux
	)
	// Marque les entrées comme réactivées
	inputDisabled = false
	return nil
}

// Goroutine qui consomme en continu les entrées du terminal
// Cette fonction tourne en arrière-plan tant que les entrées sont désactivées
func consumeInputsUnix() {
	// Récupère le file descriptor de stdin
	fd := int(os.Stdin.Fd())
	// Buffer pour stocker les caractères lus
	buffer := make([]byte, 1024)
	// Boucle tant que les entrées sont désactivées
	for IsInputDisabled() {
		// Lit les données disponibles sur stdin (non-bloquant grâce à VMIN=0)
		// Les données lues sont "consommées" et disparaissent du buffer
		syscall.Read(fd, buffer)
		// Pause de 10ms pour éviter une utilisation CPU excessive
		// C'est un compromis entre réactivité et performance
		time.Sleep(10 * time.Millisecond)
	}
}
