package inputcontrol

import (
	"os"
	"sync"
	"syscall"
)

// Variables globales pour gérer l'état des entrées utilisateur
var (
	// Mutex pour protéger les accès concurrents aux variables partagées
	// Nécessaire car plusieurs goroutines peuvent appeler ces fonctions simultanément
	inputMutex sync.Mutex
)

func ClearInputBuffer() error {
	inputMutex.Lock()
	defer inputMutex.Unlock()

	fd := int(os.Stdin.Fd())

	// TCFLSH avec TCIFLUSH = vide le buffer d'entrée
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		0x540B, // TCFLSH
		0,      // TCIFLUSH - vide le buffer d'entrée
	)

	if errno != 0 {
		return errno
	}

	return nil
}
