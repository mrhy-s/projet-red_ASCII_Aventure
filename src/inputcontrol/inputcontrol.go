package inputcontrol

import (
	"fmt"
	"runtime"
	"sync"
)

// Variables globales pour gérer l'état des entrées utilisateur
var (
	// Mutex pour protéger les accès concurrents aux variables partagées
	// Nécessaire car plusieurs goroutines peuvent appeler ces fonctions simultanément
	inputMutex sync.Mutex
	// Flag global indiquant si les entrées utilisateur sont actuellement désactivées
	// Protected par inputMutex pour éviter les race conditions
	inputDisabled bool = false
)

// IsInputDisabled retourne l'état actuel des entrées utilisateur
// Cette fonction est thread-safe grâce au mutex
func IsInputDisabled() bool {
	// Verrouille l'accès pour lecture atomique
	inputMutex.Lock()
	defer inputMutex.Unlock()
	// Retourne l'état actuel (true = désactivées, false = activées)
	return inputDisabled
}

// DisableInput désactive les entrées utilisateur de manière cross-plateforme
// Cette fonction détecte automatiquement l'OS et appelle la fonction appropriée
func DisableInput() error {
	// Détection de l'OS au runtime pour choisir l'implémentation appropriée
	switch runtime.GOOS {
	case "windows":
		// Utilise l'implémentation Windows (via kernel32.dll et API Console)
		return disableInputPlatform()
	case "linux", "darwin", "freebsd", "netbsd", "openbsd":
		// Utilise l'implémentation Unix (via termios et ioctl)
		// darwin = macOS, les autres sont des variantes de BSD/Linux
		return disableInputPlatform()
	default:
		// OS non supporté (par exemple : plan9, js/wasm, etc.)
		// Affiche un message d'erreur mais ne fait pas planter l'application
		fmt.Print("Erreur : aucun OS n'est supporté")
		return nil // Retourne nil pour éviter de casser le flow de l'application
	}
}

// EnableInput réactive les entrées utilisateur de manière cross-plateforme
// Suit le même pattern que DisableInput pour la détection d'OS
func EnableInput() error {
	// Même logique de détection d'OS que DisableInput
	switch runtime.GOOS {
	case "windows":
		// Restaure les paramètres Windows originaux
		return enableInputPlatform()
	case "linux", "darwin", "freebsd", "netbsd", "openbsd":
		// Restaure les attributs termios originaux sur Unix
		return enableInputPlatform()
	default:
		// Gestion gracieuse des OS non supportés
		fmt.Print("Erreur : aucun OS n'est supporté")
		return nil // Fail silently pour maintenir la compatibilité
	}
}
