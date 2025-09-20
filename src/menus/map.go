package menus

import (
	"ASCII_Aventure/items"
	"fmt"
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
)

var (
	Etat_1_1, Etat_1_2, Etat_1_3, Etat_1_4, Etat_1_5, Etat_1_6, Etat_1_7, Etat_1_8, Etat_1_9, Etat_1_10, Etat_1_11            string
	Etat_2_1, Etat_2_2, Etat_2_3, Etat_2_4, Etat_2_5, Etat_2_6, Etat_2_7, Etat_2_8, Etat_2_9, Etat_2_10, Etat_2_11            string
	Etat_3_1, Etat_3_2, Etat_3_3, Etat_3_4, Etat_3_5, Etat_3_6, Etat_3_7, Etat_3_8, Etat_3_9, Etat_3_10, Etat_3_11            string
	Etat_4_1, Etat_4_2, Etat_4_3, Etat_4_4, Etat_4_5, Etat_4_6, Etat_4_7, Etat_4_8, Etat_4_9, Etat_4_10, Etat_4_11            string
	Etat_5_1, Etat_5_2, Etat_5_3, Etat_5_4, Etat_5_5, Etat_5_6, Etat_5_7, Etat_5_8, Etat_5_9, Etat_5_10, Etat_5_11            string
	Etat_6_1, Etat_6_2, Etat_6_3, Etat_6_4, Etat_6_5, Etat_6_6, Etat_6_7, Etat_6_8, Etat_6_9, Etat_6_10, Etat_6_11            string
	Etat_7_1, Etat_7_2, Etat_7_3, Etat_7_4, Etat_7_5, Etat_7_6, Etat_7_7, Etat_7_8, Etat_7_9, Etat_7_10, Etat_7_11            string
	Etat_8_1, Etat_8_2, Etat_8_3, Etat_8_4, Etat_8_5, Etat_8_6, Etat_8_7, Etat_8_8, Etat_8_9, Etat_8_10, Etat_8_11            string
	Etat_9_1, Etat_9_2, Etat_9_3, Etat_9_4, Etat_9_5, Etat_9_6, Etat_9_7, Etat_9_8, Etat_9_9, Etat_9_10, Etat_9_11            string
	Etat_10_1, Etat_10_2, Etat_10_3, Etat_10_4, Etat_10_5, Etat_10_6, Etat_10_7, Etat_10_8, Etat_10_9, Etat_10_10, Etat_10_11 string
	Etat_11_1, Etat_11_2, Etat_11_3, Etat_11_4, Etat_11_5, Etat_11_6, Etat_11_7, Etat_11_8, Etat_11_9, Etat_11_10, Etat_11_11 string
)

func Map() {
	fmt.Printf("┌───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┬───────┐\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_1_1, Etat_1_2, Etat_1_3, Etat_1_4, Etat_1_5, Etat_1_6, Etat_1_7, Etat_1_8, Etat_1_9, Etat_1_10, Etat_1_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_2_1, Etat_2_2, Etat_2_3, Etat_2_4, Etat_2_5, Etat_2_6, Etat_2_7, Etat_2_8, Etat_2_9, Etat_2_10, Etat_2_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_3_1, Etat_3_2, Etat_3_3, Etat_3_4, Etat_3_5, Etat_3_6, Etat_3_7, Etat_3_8, Etat_3_9, Etat_3_10, Etat_3_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_4_1, Etat_4_2, Etat_4_3, Etat_4_4, Etat_4_5, Etat_4_6, Etat_4_7, Etat_4_8, Etat_4_9, Etat_4_10, Etat_4_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_5_1, Etat_5_2, Etat_5_3, Etat_5_4, Etat_5_5, Etat_5_6, Etat_5_7, Etat_5_8, Etat_5_9, Etat_5_10, Etat_5_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_6_1, Etat_6_2, Etat_6_3, Etat_6_4, Etat_6_5, Etat_6_6, Etat_6_7, Etat_6_8, Etat_6_9, Etat_6_10, Etat_6_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_7_1, Etat_7_2, Etat_7_3, Etat_7_4, Etat_7_5, Etat_7_6, Etat_7_7, Etat_7_8, Etat_7_9, Etat_7_10, Etat_7_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_8_1, Etat_8_2, Etat_8_3, Etat_8_4, Etat_8_5, Etat_8_6, Etat_8_7, Etat_8_8, Etat_8_9, Etat_8_10, Etat_8_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_9_1, Etat_9_2, Etat_9_3, Etat_9_4, Etat_9_5, Etat_9_6, Etat_9_7, Etat_9_8, Etat_9_9, Etat_9_10, Etat_9_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_10_1, Etat_10_2, Etat_10_3, Etat_10_4, Etat_10_5, Etat_10_6, Etat_10_7, Etat_10_8, Etat_10_9, Etat_10_10, Etat_10_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("├───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┼───────┤\n")
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("│   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │   %s   │\n", Etat_11_1, Etat_11_2, Etat_11_3, Etat_11_4, Etat_11_5, Etat_11_6, Etat_11_7, Etat_11_8, Etat_11_9, Etat_11_10, Etat_11_11)
	fmt.Printf("│       │       │       │       │       │       │       │       │       │       │       │\n")
	fmt.Printf("└───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┴───────┘\n")
}

func ExplorationAffichage(tour int) {
	fmt.Printf("=== Tour %d ===\n", tour)
	fmt.Printf("Position actuelle: (%d,%d)\n", JoueurPosition.X, JoueurPosition.Y)
	Map()
}

func NewMap() {
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
