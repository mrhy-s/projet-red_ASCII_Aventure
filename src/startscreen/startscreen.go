package startscreen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
\033[<paramètres><commande>

- \033 ou \x1b = caractère d'échappement ESC
- [ = début de la séquence
- <paramètres> = nombres séparés par ;
- <commande> = lettre finale
*/

// couleurs
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

// ClearScreen efface l'écran actuel pour avoir un écran vide dans le terminal
func ClearScreen() {
	fmt.Print("\033[2J") // efface tout le contenu de l'écran
	fmt.Print("\033[H")  // met le curseur en haut à gauche (x 1, y 1)
}

var logoASCII = []string{
	"  █████╗ ███████╗ ██████╗██╗██╗",
	" ██╔══██╗██╔════╝██╔════╝██║██║",
	" ███████║███████╗██║     ██║██║",
	" ██╔══██║╚════██║██║     ██║██║",
	" ██║  ██║███████║╚██████╗██║██║",
	" ╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝",
}

var logoAVENTURE = []string{
	"    █████╗ ██╗   ██╗███████╗███╗   ██╗████████╗██╗   ██╗██████╗ ███████╗",
	"   ██╔══██╗██║   ██║██╔════╝████╗  ██║╚══██╔══╝██║   ██║██╔══██╗██╔════╝",
	"   ███████║██║   ██║█████╗  ██╔██╗ ██║   ██║   ██║   ██║██████╔╝█████╗  ",
	"   ██╔══██║╚██╗ ██╔╝██╔══╝  ██║╚██╗██║   ██║   ██║   ██║██╔══██╗██╔══╝  ",
	"   ██║  ██║ ╚████╔╝ ███████╗██║ ╚████║   ██║   ╚██████╔╝██║  ██║███████╗",
	"   ╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═══╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚══════╝",
}

// Animation hardcodée du logo ASCII lettre par lettre
func printASCII() {
	ClearScreen()
	// Centrage vertical
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("██╗")
	fmt.Println("██║")
	fmt.Println("██║")
	fmt.Println("██║")
	fmt.Println("██║")
	fmt.Print("╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("██╗██╗")
	fmt.Println("██║██║")
	fmt.Println("██║██║")
	fmt.Println("██║██║")
	fmt.Println("██║██║")
	fmt.Print("╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("██╗██╗██╗")
	fmt.Println("══╝██║██║")
	fmt.Println("   ██║██║")
	fmt.Println("   ██║██║")
	fmt.Println("██╗██║██║")
	fmt.Print("══╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println(" ██████╗██╗██╗")
	fmt.Println("██╔════╝██║██║")
	fmt.Println("██║     ██║██║")
	fmt.Println("██║     ██║██║")
	fmt.Println("╚██████╗██║██║")
	fmt.Print(" ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("██╗ ██████╗██╗██╗")
	fmt.Println("══╝██╔════╝██║██║")
	fmt.Println("██╗██║     ██║██║")
	fmt.Println("██║██║     ██║██║")
	fmt.Println("██║╚██████╗██║██║")
	fmt.Print("══╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("█████╗ ██████╗██╗██╗")
	fmt.Println("╔════╝██╔════╝██║██║")
	fmt.Println("█████╗██║     ██║██║")
	fmt.Println("═══██║██║     ██║██║")
	fmt.Println("█████║╚██████╗██║██║")
	fmt.Print("═════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("███████╗ ██████╗██╗██╗")
	fmt.Println("██╔════╝██╔════╝██║██║")
	fmt.Println("███████╗██║     ██║██║")
	fmt.Println("╚════██║██║     ██║██║")
	fmt.Println("███████║╚██████╗██║██║")
	fmt.Print("╚══════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("╗ ███████╗ ██████╗██╗██╗")
	fmt.Println("█╗██╔════╝██╔════╝██║██║")
	fmt.Println("█║███████╗██║     ██║██║")
	fmt.Println("█║╚════██║██║     ██║██║")
	fmt.Println("█║███████║╚██████╗██║██║")
	fmt.Print("═╝╚══════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("██╗ ███████╗ ██████╗██╗██╗")
	fmt.Println("═██╗██╔════╝██╔════╝██║██║")
	fmt.Println("███║███████╗██║     ██║██║")
	fmt.Println("═██║╚════██║██║     ██║██║")
	fmt.Println(" ██║███████║╚██████╗██║██║")
	fmt.Print(" ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("████╗ ███████╗ ██████╗██╗██╗")
	fmt.Println("╔══██╗██╔════╝██╔════╝██║██║")
	fmt.Println("█████║███████╗██║     ██║██║")
	fmt.Println("╔══██║╚════██║██║     ██║██║")
	fmt.Println("║  ██║███████║╚██████╗██║██║")
	fmt.Print("╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(100 * time.Millisecond)
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	fmt.Print(Red + Bold)
	fmt.Println("  █████╗ ███████╗ ██████╗██╗██╗")
	fmt.Println(" ██╔══██╗██╔════╝██╔════╝██║██║")
	fmt.Println(" ███████║███████╗██║     ██║██║")
	fmt.Println(" ██╔══██║╚════██║██║     ██║██║")
	fmt.Println(" ██║  ██║███████║╚██████╗██║██║")
	fmt.Print(" ╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝╚═╝" + Reset)
	time.Sleep(500 * time.Millisecond)
}

/*
cette fonction affiche d'abord le logo AVENTURE ligne par ligne,
puis utilise l'animation hardcodée (PARCE QUE QUAND ON FAIT UNE BOUCLE LES LIGNES SONT PAS SYNCHRONISÉES, J'AI PAS DU TOUT LE SEUM) pour ASCII
*/
func logoAnimation() {
	ClearScreen()
	// ici l'animation hardcodée d'ASCII qui apparaît
	printASCII()
	// ASCII qui pousse pousse AVENTURE
	nombreEtapesPush := 20
	for etape := 0; etape <= nombreEtapesPush; etape++ {
		ClearScreen()
		for range 5 {
			fmt.Println()
		}
		// on calcule l'espacement qui diminue entre avanture et ASCII
		espacementEntreLogos := 20 - etape
		if espacementEntreLogos < 3 {
			espacementEntreLogos = 3 // minimum 3 espaces
		}
		// on affiche les 6 lignes des logos côte à côte
		for numeroLigne := range 6 {
			ligneComplete := Red + Bold + logoASCII[numeroLigne] + Reset
			for range espacementEntreLogos {
				ligneComplete += " "
			}
			ligneComplete += Green + Bold + logoAVENTURE[numeroLigne] + Reset
			// on veut centrer toute la ligne
			largeurEcran := 120
			largeurTotale := len(logoASCII[numeroLigne]) + espacementEntreLogos + len(logoAVENTURE[numeroLigne])
			espacesAGauche := (largeurEcran - largeurTotale) / 2
			if espacesAGauche < 0 {
				espacesAGauche = 0
			}
			fmt.Printf("%*s%s\n", espacesAGauche, "", ligneComplete)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// ici, l'animation de chargement avec barre de progression
func loadingAnimation() {
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	for i := range 6 {
		fmt.Print(Red + Bold + logoASCII[i] + Reset)
		fmt.Printf("%3s", "")
		fmt.Print(Green + Bold + logoAVENTURE[i] + Reset)
		fmt.Println()
	}
	// toujours l'animation de chargement
	texteChargement := "Chargement de l'aventure"
	largeurEcran := 120
	espacesAGauche := (largeurEcran - len(texteChargement) - 10) / 2
	if espacesAGauche < 0 {
		espacesAGauche = 0
	}
	for i := range 20 {
		fmt.Print("\033[14;1H")
		nombrePoints := i % 4
		fmt.Printf("%*s%s%s%s", espacesAGauche, "", Cyan+Bold, texteChargement, Reset)
		fmt.Printf("%.*s%*s", nombrePoints, "...", 3-nombrePoints, "")
		fmt.Printf("\n%*s", espacesAGauche, "")
		pourcentage := i * 5
		barresRemplies := pourcentage / 5
		barresVides := 20 - barresRemplies
		fmt.Printf("%s[%s", Cyan, Yellow)
		fmt.Printf("%.*s", barresRemplies, "████████████████████")
		fmt.Printf("%s", White)
		fmt.Printf("%.*s", barresVides, "░░░░░░░░░░░░░░░░░░░░")
		fmt.Printf("%s] %d%%%s", Cyan, pourcentage, Reset)
		time.Sleep(100 * time.Millisecond)
	}
}

// Maintenant voici la fonction qui permet de mettre ensemble tout ce boulgiboulga. Elle devra être appelée dans le main.go
func StartScreen() {
	fmt.Print("\033[?25l") // cacher le curseur
	logoAnimation()        // animation d'apparition des logos
	time.Sleep(1 * time.Second)
	loadingAnimation() // animation de chargement
	time.Sleep(500 * time.Millisecond)
	// Affichage final
	ClearScreen()
	for range 5 {
		fmt.Println()
	}
	for i := range 6 {
		fmt.Print(Red + Bold + logoASCII[i] + Reset)
		fmt.Printf("%3s", "")
		fmt.Print(Green + Bold + logoAVENTURE[i] + Reset)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	messageFinal := "Appuyez sur ENTRÉE pour commencer votre aventure..."
	espacesAGauche := (120 - len(messageFinal)) / 2
	if espacesAGauche < 0 {
		espacesAGauche = 0
	}
	fmt.Printf("%*s%s%s%s%s\n", espacesAGauche, "", Cyan, Bold, messageFinal, Reset)

	input := readInput()
	if input != "" {
		fmt.Printf("\n%sÇa commence bien... On te dit d'appuyer sur la touche ENTRÉE et toi tu tapes '%s'%s\n",
			Red, input, Reset)
		fmt.Printf("%sAppuyez juste sur ENTRÉE cette fois...%s\n", Cyan, Reset)
		readInput()
	}
	fmt.Print("\033[?25h") // remonter le curseur
	ClearScreen()
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
