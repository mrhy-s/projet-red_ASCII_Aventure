package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

type Character struct {
	nom                   string
	classe                string
	niveau                int
	points_de_vie_maximum int
	points_de_vie_actuels int
	inventaire            []string
}

var c1 *Character
var c2 *Character
var c2_b bool

var c_temp_name string
var c_temp_classe string
var c_temp_niveau int
var c_temp_points_de_vie_maximum int
var c_temp_points_de_vie_actuels int
var c_temp_inventaire []string

var tour int

func main() {
	IsFirst := true
	c1 = initCharacter("Zinrel", "Elfe", 1, 100, 40, []string{"potions de soin", "potions de soin", "potions de soin"})

	if IsFirst {
		min_pt_vie_rand, max_pt_vie_rand := 95, 105
		fmt.Print("Voulez vous utiliser le personnage par défaut ? Ou voulez vous créer un nouveau personnage ?\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'\n\n")
		displayCharacterTable(*c1)
		fmt.Print("\nVotre réponse : ")
		input := readInput()
		if input == "Oui" || input == "Oui." || input == "oui" || input == "oui." {
			fmt.Print("\nVeuillez entrer le nom du personnage : ")
			c_temp_name = readInput()
			fmt.Print("\nVeuillez entrer la classe du personnage : ")
			c_temp_classe = readInput()
			fmt.Printf("\nVotre personnage s'appelle : %s\nLa classe de votre personnage est : %s", c_temp_name, c_temp_classe)
			fmt.Print("\nVotre personnage commence au niveau 0.")
			c_temp_niveau = 0
			c_temp_points_de_vie_maximum = randomBetween(min_pt_vie_rand, max_pt_vie_rand)
			c_temp_points_de_vie_actuels = c_temp_points_de_vie_maximum - randomBetween(20, 70)
			fmt.Printf("\nVotre personnage avec %d/%d points de vie.", c_temp_points_de_vie_actuels, c_temp_points_de_vie_maximum)
			c_temp_inventaire = []string{"potions de soin", "potions de soin", "potions de soin"}
			fmt.Print("\nvous commencez avec 'Potion de soin x3'\n")
			c2_b = true
			c2 = initCharacter(c_temp_name, c_temp_classe, c_temp_niveau, c_temp_points_de_vie_maximum, c_temp_points_de_vie_actuels, c_temp_inventaire)
			fmt.Println("\nVoici votre nouveau personnage :")
			displayCharacterTable(*c2)
		}
		IsFirst = false
	}
	menu()
}

// ================
// fonction actions
// ================

func initCharacter(nom string, classe string, niveau int, points_de_vie_maximum int, points_de_vie_actuels int, inventaire []string) *Character {
	character_template := &Character{
		nom:                   nom,
		classe:                classe,
		niveau:                niveau,
		points_de_vie_maximum: points_de_vie_maximum,
		points_de_vie_actuels: points_de_vie_actuels,
		inventaire:            inventaire,
	}
	return character_template
}

func displayInfo(s string) {
	s = strings.TrimSpace(s)
	if strings.Contains(strings.ToLower(s), "liste") {
		fmt.Println("\n=== LISTE DES PERSONNAGES ===")
		fmt.Println("\nPersonnage 1 :")
		displayCharacterTable(*c1)
		if c2_b && c2 != nil {
			fmt.Println("\nPersonnage 2 :")
			displayCharacterTable(*c2)
		}
		return
	}
	if s == c1.nom {
		displayCharacterTable(*c1)
		return
	}
	if c2_b && c2 != nil && s == c2.nom {
		displayCharacterTable(*c2)
		return
	}
	fmt.Println("Personnage non trouvé")
}

func accessInventory(s string) {
	var character *Character
	if c2_b {
		switch s {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = c1
	}
	fmt.Printf("\n┌─────────────────────────────────────────────────┐\n")
	fmt.Printf("│ Inventaire de %-33s │\n", character.nom)
	fmt.Printf("├─────────────────────────────────────────────────┤\n")
	if len(character.inventaire) == 0 {
		fmt.Printf("│ Inventaire vide                                 │\n")
	} else {
		for i, item := range character.inventaire {
			fmt.Printf("│ %d. %-44s │\n", i+1, item)
		}
	}
	fmt.Printf("└─────────────────────────────────────────────────┘\n")
}

func takePot(characterName string) {
	var character *Character
	if c2_b {
		switch characterName {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = c1
	}
	potionIndex := -1
	potionType := ""
	for i, item := range character.inventaire {
		if item == "potions de soin" || item == "potion de vie" {
			potionIndex = i
			potionType = item
			break
		}
	}
	if potionIndex == -1 {
		fmt.Printf("\n%s n'a pas de potion dans son inventaire.\n", character.nom)
		return
	}
	character.points_de_vie_actuels += 50
	if character.points_de_vie_actuels > character.points_de_vie_maximum {
		character.points_de_vie_actuels = character.points_de_vie_maximum
	}
	character.inventaire = removeItem(character.inventaire, potionIndex)
	fmt.Printf("\n%s a utilisé une %s ! (+50 PV)\n", character.nom, potionType)
	fmt.Printf("Points de vie: %d/%d\n", character.points_de_vie_actuels, character.points_de_vie_maximum)
}

func marchand(tour int) {
	fmt.Print("\n=== BOUTIQUE DU MARCHAND ===\n")
	fmt.Print("Bienvenue dans ma boutique\n")
	fmt.Print("Voici les items disponibles :\n\n")
	fmt.Print("1. Potion de vie - GRATUIT\n")
	fmt.Print("2. Retourner au menu principal\n")
	fmt.Print("\nVotre choix : ")

	choix := readInput()

	switch choix {
	case "1", "1.":
		var characterName string
		if c2_b && c2 != nil {
			fmt.Print("\nQuel personnage souhaite prendre la potion ?\n")
			fmt.Printf("1. %s\n", c1.nom)
			fmt.Printf("2. %s\n", c2.nom)
			fmt.Print("Votre choix : ")
			choixPerso := readInput()
			switch choixPerso {
			case "1", "1.":
				characterName = c1.nom
			case "2", "2.":
				characterName = c2.nom
			default:
				fmt.Println("Choix invalide")
				return
			}
		} else {
			characterName = c1.nom
		}
		addInventory(characterName, "potion de vie")
		fmt.Printf("\nPotion de vie ajoutée à l'inventaire de %s !\n", characterName)
		fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
		continuer := readInput()
		if continuer == "o" || continuer == "oui" || continuer == "O" || continuer == "Oui" {
			marchand(tour)
		}
	case "2", "2.":
		return
	default:
		fmt.Println("Choix invalide")
		marchand(tour)
	}
}

// ==============
// fonction menus
// ==============

func menu() {
	fmt.Printf("\n")
	fmt.Printf("┌─────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│                                  MENU PRINCIPAL                                     │\n")
	fmt.Printf("├─────────────────────────────────────────────────────────────────────────────────────┤\n")
	fmt.Printf("│ 1. Afficher les informations du personnage                                          │\n")
	fmt.Printf("│ 2. Accéder au contenu de l'inventaire                                               │\n")
	fmt.Printf("│ 3. Utiliser une potion de soin                                                      │\n")
	fmt.Printf("│ 4. Boutique du Marchand                                                             │\n")
	fmt.Printf("│ 5. Quitter                                                                          │\n")
	fmt.Printf("└─────────────────────────────────────────────────────────────────────────────────────┘\n")
	fmt.Print("\nVotre choix : ")
	option := readInput()
	switch option {
	case "1", "1.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'\n")
		fmt.Print("\nVotre réponse : ")
		commande := readInput()
		displayInfo(commande)
		menu()
	case "2", "2.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher son inventaire :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'\n")
		fmt.Print("\nVotre réponse : ")
		commande := readInput()
		accessInventory(commande)
		menu()
	case "3", "3.":
		fmt.Print("Nom du personnage : ")
		characterName := readInput()
		takePot(characterName)
		menu()
	case "4", "4.":
		marchand(tour)
		menu()
	case "5", "5.":
		return
	default:
		fmt.Println("Option invalide")
		menu()
	}
}

func displayCharacterTable(character Character) {
	inventaireStr := strings.Join(character.inventaire, ", ")
	// Largeur totale du cadre : 85 caractères
	// Largeur utile pour le contenu : 85 - 4 (bordures et espaces) = 81 caractères
	// Largeur pour les valeurs après "   - XXX : " : environ 65-67 caractères selon le label
	const totalWidth = 85
	const contentWidth = 81
	inventaireMaxWidth := 62 // 81 - 18 - 1 = 62
	// Tronquer l'inventaire si nécessaire
	if utf8.RuneCountInString(inventaireStr) > inventaireMaxWidth {
		// Garder de la place pour "..."
		targetLength := inventaireMaxWidth - 3
		runes := []rune(inventaireStr)
		if len(runes) > targetLength {
			inventaireStr = string(runes[:targetLength]) + "..."
		}
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Printf("│%-85s│\n", "Caractéristiques du personnage :")
	fmt.Printf("│   - Nom : %-74s│\n", character.nom)
	fmt.Printf("│   - Classe : %-71s│\n", character.classe)
	fmt.Printf("│   - Niveau : %-71d│\n", character.niveau)
	fmt.Printf("│   - Points de vie maximum : %-56d│\n", character.points_de_vie_maximum)
	fmt.Printf("│   - Points de vie actuels : %-56d│\n", character.points_de_vie_actuels)
	fmt.Printf("│   - Inventaire : [%-65s]│\n", inventaireStr)
	fmt.Printf("│%-85s│\n", "")
	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
}

// ================
// fonction helpers
// ================

func randomBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func removeItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func removeInventory(characterName string, item string) bool {
	var character *Character
	if c2_b {
		switch characterName {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return false
		}
	} else {
		character = c1
	}

	for i, inventoryItem := range character.inventaire {
		if inventoryItem == item {
			character.inventaire = removeItem(character.inventaire, i)
			return true
		}
	}
	return false
}

func addInventory(characterName string, item string) {
	var character *Character
	if c2_b {
		switch characterName {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = c1
	}
	character.inventaire = append(character.inventaire, item)
}

func isdead(characterName string) {
	var character *Character
	if c2_b {
		switch characterName {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = c1
	}
	if character.points_de_vie_actuels <= 0 {
		print("Vous êtes mort ...")
		character.points_de_vie_actuels = character.points_de_vie_maximum / 2
	}
}

// faire de sorte a ce que le marchand ne propose la potion gratuite qu'une fois
