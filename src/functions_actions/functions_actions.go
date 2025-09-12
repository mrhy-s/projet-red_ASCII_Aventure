package functionsactions

import (
	characters "ASCII_Aventure/characters"
	functionshelper "ASCII_Aventure/functions_helper"
	"fmt"
	"strings"
)

var Tour int

// ================
// fonction actions
// ================

func DisplayInfo(s string) {
	s = strings.TrimSpace(s)
	if strings.Contains(strings.ToLower(s), "liste") {
		fmt.Println("\n=== LISTE DES PERSONNAGES ===")
		fmt.Println("\nPersonnage 1 :")
		characters.DisplayCharacterTable(*characters.C1)
		if characters.C2_b && characters.C2 != nil {
			fmt.Println("\nPersonnage 2 :")
			characters.DisplayCharacterTable(*characters.C2)
		}
		return
	}
	if s == characters.C1.Nom {
		characters.DisplayCharacterTable(*characters.C1)
		return
	}
	if characters.C2_b && characters.C2 != nil && s == characters.C2.Nom {
		characters.DisplayCharacterTable(*characters.C2)
		return
	}
	fmt.Println("Personnage non trouvé")
}

func AccessInventory(s string) {
	var character *characters.Character
	if characters.C2_b {
		switch s {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = characters.C1
	}
	fmt.Printf("\n┌─────────────────────────────────────────────────┐\n")
	fmt.Printf("│ Inventaire de %-33s │\n", character.Nom)
	fmt.Printf("├─────────────────────────────────────────────────┤\n")
	if len(character.Inventaire) == 0 {
		fmt.Printf("│ Inventaire vide                                 │\n")
	} else {
		for i, item := range character.Inventaire {
			fmt.Printf("│ %d. %-44s │\n", i+1, item)
		}
	}
	fmt.Printf("└─────────────────────────────────────────────────┘\n")
}

func TakePot(characterName string) {
	var character *characters.Character
	if characters.C2_b {
		switch characterName {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = characters.C1
	}
	potionIndex := -1
	potionType := ""
	for i, item := range character.Inventaire {
		if item == "potions de soin" || item == "potion de vie" {
			potionIndex = i
			potionType = item
			break
		}
	}
	if potionIndex == -1 {
		fmt.Printf("\n%s n'a pas de potion dans son inventaire.\n", character.Nom)
		return
	}
	character.PointsDeVieActuels += 50
	if character.PointsDeVieActuels > character.PointsDeVieMaximum {
		character.PointsDeVieActuels = character.PointsDeVieMaximum
	}
	character.Inventaire = functionshelper.RemoveItem(character.Inventaire, potionIndex)
	fmt.Printf("\n%s a utilisé une %s ! (+50 PV)\n", character.Nom, potionType)
	fmt.Printf("Points de vie: %d/%d\n", character.PointsDeVieActuels, character.PointsDeVieMaximum)
}

func Marchand(tour int) {
	fmt.Print("\n=== BOUTIQUE DU MARCHAND ===\n")
	fmt.Print("Bienvenue dans ma boutique\n")
	fmt.Print("Voici les items disponibles :\n\n")
	fmt.Print("1. Potion de vie - GRATUIT\n")
	fmt.Print("2. Retourner au menu principal\n")
	fmt.Print("\nVotre choix : ")

	choix := functionshelper.ReadInput()

	switch choix {
	case "1", "1.":
		var characterName string
		if characters.C2_b && characters.C2 != nil {
			fmt.Print("\nQuel personnage souhaite prendre la potion ?\n")
			fmt.Printf("1. %s\n", characters.C1.Nom)
			fmt.Printf("2. %s\n", characters.C2.Nom)
			fmt.Print("Votre choix : ")
			choixPerso := functionshelper.ReadInput()
			switch choixPerso {
			case "1", "1.":
				characterName = characters.C1.Nom
			case "2", "2.":
				characterName = characters.C2.Nom
			default:
				fmt.Println("Choix invalide")
				return
			}
		} else {
			characterName = characters.C1.Nom
		}
		functionshelper.AddInventory(characterName, "potion de vie")
		fmt.Printf("\nPotion de vie ajoutée à l'inventaire de %s !\n", characterName)
		fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
		continuer := functionshelper.ReadInput()
		if continuer == "o" || continuer == "oui" || continuer == "O" || continuer == "Oui" {
			Marchand(tour)
		}
	case "2", "2.":
		return
	default:
		fmt.Println("Choix invalide")
		Marchand(tour)
	}
}
