package menus

import (
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/items"
	"fmt"
	"strings"
)

var Explorée string = "◦"
var Inexplorée string = "?"
var Chemin string = "·"
var Bâtiment string = "■"
var ZoneCombat string = "⚔"
var Joueur string = "⋄"

// PNJ représente un personnage non-joueur
type PNJ struct {
	Nom        string
	Dialogue   string
	Quete      string
	Recompense string
}

// Zone représente une sous-zone de la map
type Zone struct {
	Nom         string
	Description string
	Ressources  []items.Item
	Monstres    []string
	PNJs        []PNJ
	Visitee     bool
}

// Position du joueur sur la map
type Position struct {
	X, Y int
}

// Variables globales pour la position du joueur et la map des zones
var (
	JoueurPosition Position
	ZonesMap       map[string]Zone
	TourG          int
)

var (
	Etat_1_1, Etat_1_2, Etat_1_3, Etat_1_4, Etat_1_5, Etat_1_6, Etat_1_7, Etat_1_8, Etat_1_9, Etat_1_10, Etat_1_11                                                                                                                                                                                                                                                                                             string
	Etat_2_1, Etat_2_2, Etat_2_3, Etat_2_4, Etat_2_5, Etat_2_6, Etat_2_7, Etat_2_8, Etat_2_9, Etat_2_10, Etat_2_11                                                                                                                                                                                                                                                                                             string
	Etat_3_1, Etat_3_2, Etat_3_3, Etat_3_4, Etat_3_5, Etat_3_6, Etat_3_7, Etat_3_8, Etat_3_9, Etat_3_10, Etat_3_11                                                                                                                                                                                                                                                                                             string
	Etat_4_1, Etat_4_2, Etat_4_3, Etat_4_4, Etat_4_5, Etat_4_6, Etat_4_7, Etat_4_8, Etat_4_9, Etat_4_10, Etat_4_11                                                                                                                                                                                                                                                                                             string
	Etat_5_1, Etat_5_2, Etat_5_3, Etat_5_4, Etat_5_5, Etat_5_6, Etat_5_7, Etat_5_8, Etat_5_9, Etat_5_10, Etat_5_11                                                                                                                                                                                                                                                                                             string
	Etat_6_1, Etat_6_2, Etat_6_3, Etat_6_4, Etat_6_5, Etat_6_6, Etat_6_7, Etat_6_8, Etat_6_9, Etat_6_10, Etat_6_11                                                                                                                                                                                                                                                                                             string
	Etat_7_1, Etat_7_2, Etat_7_3, Etat_7_4, Etat_7_5, Etat_7_6, Etat_7_7, Etat_7_8, Etat_7_9, Etat_7_10, Etat_7_11                                                                                                                                                                                                                                                                                             string
	Etat_8_1, Etat_8_2, Etat_8_3, Etat_8_4, Etat_8_5, Etat_8_6, Etat_8_7, Etat_8_8, Etat_8_9, Etat_8_10, Etat_8_11                                                                                                                                                                                                                                                                                             string
	Etat_9_1, Etat_9_2, Etat_9_3, Etat_9_4, Etat_9_5, Etat_9_6, Etat_9_7, Etat_9_8, Etat_9_9, Etat_9_10, Etat_9_11                                                                                                                                                                                                                                                                                             string
	Etat_10_1, Etat_10_2, Etat_10_3, Etat_10_4, Etat_10_5, Etat_10_6, Etat_10_7, Etat_10_8, Etat_10_9, Etat_10_10, Etat_10_11                                                                                                                                                                                                                                                                                  string
	Etat_11_1, Etat_11_2, Etat_11_3, Etat_11_4, Etat_11_5, Etat_11_6, Etat_11_7, Etat_11_8, Etat_11_9, Etat_11_10, Etat_11_11                                                                                                                                                                                                                                                                                  string
	Ligne1, Ligne2, Ligne3, Ligne4, Ligne5, Ligne6, Ligne7, Ligne8, Ligne9, Ligne10, Ligne11, Ligne12, Ligne13, Ligne14, Ligne15, Ligne16, Ligne17, Ligne18, Ligne19, Ligne20, Ligne21, Ligne22, Ligne23, Ligne24, Ligne25, Ligne26, Ligne27, Ligne28, Ligne29, Ligne30, Ligne31, Ligne32, Ligne33, Ligne34, Ligne35, Ligne36, Ligne37, Ligne38, Ligne39, Ligne40, Ligne41, Ligne42, Ligne43, Ligne44, Ligne45 string
	espacementDroite                                                                                                                                                                                                                                                                                                                                                                                           string
)

func Map() {
	fmt.Printf("┌───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┐ %s\n", Ligne1)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne2)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_1_1, Etat_1_2, Etat_1_3, Etat_1_4, Etat_1_5, Etat_1_6, Etat_1_7, Etat_1_8, Etat_1_9, Etat_1_10, Etat_1_11, Ligne3)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne4)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne5)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne6)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_2_1, Etat_2_2, Etat_2_3, Etat_2_4, Etat_2_5, Etat_2_6, Etat_2_7, Etat_2_8, Etat_2_9, Etat_2_10, Etat_2_11, Ligne7)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne8)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne9)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne10)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_3_1, Etat_3_2, Etat_3_3, Etat_3_4, Etat_3_5, Etat_3_6, Etat_3_7, Etat_3_8, Etat_3_9, Etat_3_10, Etat_3_11, Ligne11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne12)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne13)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne14)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_4_1, Etat_4_2, Etat_4_3, Etat_4_4, Etat_4_5, Etat_4_6, Etat_4_7, Etat_4_8, Etat_4_9, Etat_4_10, Etat_4_11, Ligne15)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne16)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne17)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne18)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_5_1, Etat_5_2, Etat_5_3, Etat_5_4, Etat_5_5, Etat_5_6, Etat_5_7, Etat_5_8, Etat_5_9, Etat_5_10, Etat_5_11, Ligne19)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne20)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne21)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne22)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_6_1, Etat_6_2, Etat_6_3, Etat_6_4, Etat_6_5, Etat_6_6, Etat_6_7, Etat_6_8, Etat_6_9, Etat_6_10, Etat_6_11, Ligne23)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne24)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne25)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne26)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_7_1, Etat_7_2, Etat_7_3, Etat_7_4, Etat_7_5, Etat_7_6, Etat_7_7, Etat_7_8, Etat_7_9, Etat_7_10, Etat_7_11, Ligne27)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne28)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne29)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne30)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_8_1, Etat_8_2, Etat_8_3, Etat_8_4, Etat_8_5, Etat_8_6, Etat_8_7, Etat_8_8, Etat_8_9, Etat_8_10, Etat_8_11, Ligne31)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne32)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne33)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne34)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_9_1, Etat_9_2, Etat_9_3, Etat_9_4, Etat_9_5, Etat_9_6, Etat_9_7, Etat_9_8, Etat_9_9, Etat_9_10, Etat_9_11, Ligne35)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne36)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne37)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne38)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_10_1, Etat_10_2, Etat_10_3, Etat_10_4, Etat_10_5, Etat_10_6, Etat_10_7, Etat_10_8, Etat_10_9, Etat_10_10, Etat_10_11, Ligne39)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne40)
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤ %s\n", Ligne41)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne42)
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │ %s\n", Etat_11_1, Etat_11_2, Etat_11_3, Etat_11_4, Etat_11_5, Etat_11_6, Etat_11_7, Etat_11_8, Etat_11_9, Etat_11_10, Etat_11_11, Ligne43)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │ %s\n", Ligne44)
	fmt.Printf("└───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┘ %s\n", Ligne45)
}

func ExplorationAffichage(tour int) {
	Map()
	Ligne1 = fmt.Sprintf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐", espacementDroite)
	Ligne2 = fmt.Sprintf("%s│ Position actuelle: (X:%-2d, Y:%-2d)%54s", espacementDroite, JoueurPosition.X, JoueurPosition.Y, "│")
	Ligne3 = fmt.Sprintf("%s│ Légende: Joueur = ⋄ ; Zone de combat = ⚔ ; Bâtiments = ■ ; Chemin = · ;             │", espacementDroite)
	Ligne4 = fmt.Sprintf("%s│          Zone explorée = ◦ ; Zone inexplorée = ?                                    │", espacementDroite)
	Ligne5 = fmt.Sprintf("%s├─────────────────────────────────────────────────────────────────────────────────────┤", espacementDroite)

	if TourG == 0 {
		dialogue_intro := "Tu te trouve dans un monde où les villages modestes côtoient des ruines oubliées. Ici, la route n'est jamais tracée d'avance: on apprend en marchant, on grandit en choisissant. Tu démarres avec peu, mais chaque pas compte. Parle aux gens, écoute le vent, observe les traces sur la terre. Certaines portes s'ouvrent avec du courage, d'autres avec de la ruse. Si tu te perds, cherche un repère, une lumière, un sentier. Prends le temps, sois curieux, et fais de cette terre ton histoire."
		dialogue_intro2 := "Te voici à la capitale. Les rues bourdonnent et les gardes vont et viennent. Tu peux rejoindre le marchand pour vendre ou acheter ce qu'il te manque. Le forgeron répare ton équipement et peut améliorer tes armes. Des PNJ proposent des quêtes secondaires sur la place et dans les ruelles. Pour l'histoire principale, adresse toi a Sebastiano près de l'avenue centrale. Explore sans te presser, observe les enseignes, et choisis ta prochaine étape."
		maxWidth := 81

		// Premier dialogue
		lines := functionshelper.WrapText(dialogue_intro, maxWidth)

		Ligne6 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "Bienvenue, voyageur.")
		Ligne7 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "")

		// Attribution des lignes du premier dialogue
		lignesDialogue := []*string{&Ligne8, &Ligne9, &Ligne10, &Ligne11, &Ligne12, &Ligne13, &Ligne14, &Ligne15}
		for i, line := range lines {
			if i < len(lignesDialogue) {
				*lignesDialogue[i] = fmt.Sprintf("%s│ %-84s│", espacementDroite, line)
			}
		}

		// Remplir les lignes vides restantes du premier dialogue
		for i := len(lines); i < len(lignesDialogue); i++ {
			*lignesDialogue[i] = fmt.Sprintf("%s│ %-84s│", espacementDroite, "")
		}

		// Deuxième dialogue
		lines2 := functionshelper.WrapText(dialogue_intro2, maxWidth)

		Ligne17 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "")

		// Attribution des lignes du deuxième dialogue
		lignesDialogue2 := []*string{&Ligne16, &Ligne17, &Ligne18, &Ligne19, &Ligne20, &Ligne21, &Ligne22}
		for i, line := range lines2 {
			if i < len(lignesDialogue2) {
				*lignesDialogue2[i] = fmt.Sprintf("%s│ %-84s│", espacementDroite, line)
			}
		}

		// Remplir les lignes vides restantes du deuxième dialogue
		for i := len(lines2); i < len(lignesDialogue2); i++ {
			*lignesDialogue2[i] = fmt.Sprintf("%s│ %-84s│", espacementDroite, "")
		}

		Ligne23 = fmt.Sprintf("%s├─────────────────────────────────────────────────────────────────────────────────────┤", espacementDroite)
		Ligne24 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "Que souhaitez-vous faire ?")
		Ligne25 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "C. Entrer dans la capitale")
		Ligne26 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "Z. Aller au Nord")
		Ligne27 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "Q. Aller à l'Ouest")
		Ligne28 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "S. Aller au Sud")
		Ligne29 = fmt.Sprintf("%s│ %-84s│", espacementDroite, "D. Aller a l'Est")
		Ligne30 = fmt.Sprintf("%s└─────────────────────────────────────────────────────────────────────────────────────┘", espacementDroite)
	}
	input := functionshelper.ReadInput()
	switch input {
	case strings.ToLower("c"), strings.ToLower("c."):
		Capitale()
	case strings.ToLower("z"), strings.ToLower("z."):
		fmt.Print("Debug: Z")
	case strings.ToLower("q"), strings.ToLower("q."):
		fmt.Print("Debug: Q")
	case strings.ToLower("s"), strings.ToLower("s."):
		fmt.Print("Debug: S")
	case strings.ToLower("d"), strings.ToLower("d."):
		fmt.Print("Debug: D")
	}
}

func NewMap() {
	TourG = 0
	espacementDroite = "     "
	// Initialise la position du joueur au centre
	JoueurPosition = Position{X: 6, Y: 6}

	// initialise les variables pour définir quelle case est quoi

	// Ligne 1
	Etat_1_1 = Inexplorée
	Etat_1_2 = Inexplorée
	Etat_1_3 = Inexplorée
	Etat_1_4 = Inexplorée
	Etat_1_5 = Inexplorée
	Etat_1_6 = Inexplorée
	Etat_1_7 = Inexplorée
	Etat_1_8 = Inexplorée
	Etat_1_9 = Inexplorée
	Etat_1_10 = Inexplorée
	Etat_1_11 = Inexplorée

	// Ligne 2
	Etat_2_1 = Inexplorée
	Etat_2_2 = Inexplorée
	Etat_2_3 = Inexplorée
	Etat_2_4 = Inexplorée
	Etat_2_5 = Inexplorée
	Etat_2_6 = Inexplorée
	Etat_2_7 = Inexplorée
	Etat_2_8 = Inexplorée
	Etat_2_9 = Inexplorée
	Etat_2_10 = Inexplorée
	Etat_2_11 = Inexplorée

	// Ligne 3
	Etat_3_1 = Inexplorée
	Etat_3_2 = Inexplorée
	Etat_3_3 = Inexplorée
	Etat_3_4 = Inexplorée
	Etat_3_5 = Inexplorée
	Etat_3_6 = Inexplorée
	Etat_3_7 = Inexplorée
	Etat_3_8 = Inexplorée
	Etat_3_9 = Inexplorée
	Etat_3_10 = Inexplorée
	Etat_3_11 = Inexplorée

	// Ligne 4
	Etat_4_1 = Inexplorée
	Etat_4_2 = Inexplorée
	Etat_4_3 = Inexplorée
	Etat_4_4 = Inexplorée
	Etat_4_5 = Inexplorée
	Etat_4_6 = Inexplorée
	Etat_4_7 = Inexplorée
	Etat_4_8 = Inexplorée
	Etat_4_9 = Inexplorée
	Etat_4_10 = Inexplorée
	Etat_4_11 = Inexplorée

	// Ligne 5
	Etat_5_1 = Inexplorée
	Etat_5_2 = Inexplorée
	Etat_5_3 = Inexplorée
	Etat_5_4 = Inexplorée
	Etat_5_5 = Inexplorée
	Etat_5_6 = Inexplorée
	Etat_5_7 = Inexplorée
	Etat_5_8 = Inexplorée
	Etat_5_9 = Inexplorée
	Etat_5_10 = Inexplorée
	Etat_5_11 = Inexplorée

	// Ligne 6
	Etat_6_1 = Inexplorée
	Etat_6_2 = Inexplorée
	Etat_6_3 = Inexplorée
	Etat_6_4 = Inexplorée
	Etat_6_5 = Inexplorée
	Etat_6_6 = Joueur // Position de départ du joueur au centre
	Etat_6_7 = Inexplorée
	Etat_6_8 = Inexplorée
	Etat_6_9 = Inexplorée
	Etat_6_10 = Inexplorée
	Etat_6_11 = Inexplorée

	// Ligne 7
	Etat_7_1 = Inexplorée
	Etat_7_2 = Inexplorée
	Etat_7_3 = Inexplorée
	Etat_7_4 = Inexplorée
	Etat_7_5 = Inexplorée
	Etat_7_6 = Inexplorée
	Etat_7_7 = Inexplorée
	Etat_7_8 = Inexplorée
	Etat_7_9 = Inexplorée
	Etat_7_10 = Inexplorée
	Etat_7_11 = Inexplorée

	// Ligne 8
	Etat_8_1 = Inexplorée
	Etat_8_2 = Inexplorée
	Etat_8_3 = Inexplorée
	Etat_8_4 = Inexplorée
	Etat_8_5 = Inexplorée
	Etat_8_6 = Inexplorée
	Etat_8_7 = Inexplorée
	Etat_8_8 = Inexplorée
	Etat_8_9 = Inexplorée
	Etat_8_10 = Inexplorée
	Etat_8_11 = Inexplorée

	// Ligne 9
	Etat_9_1 = Inexplorée
	Etat_9_2 = Inexplorée
	Etat_9_3 = Inexplorée
	Etat_9_4 = Inexplorée
	Etat_9_5 = Inexplorée
	Etat_9_6 = Inexplorée
	Etat_9_7 = Inexplorée
	Etat_9_8 = Inexplorée
	Etat_9_9 = Inexplorée
	Etat_9_10 = Inexplorée
	Etat_9_11 = Inexplorée

	// Ligne 10
	Etat_10_1 = Inexplorée
	Etat_10_2 = Inexplorée
	Etat_10_3 = Inexplorée
	Etat_10_4 = Inexplorée
	Etat_10_5 = Inexplorée
	Etat_10_6 = Inexplorée
	Etat_10_7 = Inexplorée
	Etat_10_8 = Inexplorée
	Etat_10_9 = Inexplorée
	Etat_10_10 = Inexplorée
	Etat_10_11 = Inexplorée

	// Ligne 11
	Etat_11_1 = Inexplorée
	Etat_11_2 = Inexplorée
	Etat_11_3 = Inexplorée
	Etat_11_4 = Inexplorée
	Etat_11_5 = Inexplorée
	Etat_11_6 = Inexplorée
	Etat_11_7 = Inexplorée
	Etat_11_8 = Inexplorée
	Etat_11_9 = Inexplorée
	Etat_11_10 = Inexplorée
	Etat_11_11 = Inexplorée
}

func LoadMap() {
	// Initialise la map des zones
	ZonesMap = make(map[string]Zone)

	// Exemple de zones prédéfinies
	ZonesMap["1-1"] = Zone{
		Nom:         "Forêt Sombre",
		Description: "Une forêt dense et mystérieuse",
		Ressources:  []items.Item{},
		Monstres:    []string{"Loup", "Araignée"},
		PNJs:        []PNJ{},
		Visitee:     false,
	}

	ZonesMap["6-3"] = Zone{
		Nom:         "Village",
		Description: "Un petit village paisible",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Marchand", Dialogue: "Bienvenue étranger!", Quete: "", Recompense: ""}},
		Visitee:     false,
	}

	ZonesMap["11-11"] = Zone{
		Nom:         "Donjon",
		Description: "Un donjon sombre et dangereux",
		Ressources:  []items.Item{},
		Monstres:    []string{"Squelette", "Orc", "Dragon"},
		PNJs:        []PNJ{},
		Visitee:     false,
	}
}
