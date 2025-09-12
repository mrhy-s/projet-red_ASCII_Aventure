package functionsactions

import (
	characters "ASCII_Aventure/characters"
	functionshelper "ASCII_Aventure/functions_helper"
	"fmt"
	"strings"
)

var Tour int
var PotionGratuite bool

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
			itemName := fmt.Sprintf("%v", item)
			fmt.Printf("│ %d. %-44s │\n", i+1, itemName)
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
		itemStr := fmt.Sprintf("%v", item)
		if itemStr == "potion de soin" || itemStr == "potion de vie" {
			potionIndex = i
			potionType = itemStr
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
	const prixBase = 10
	for {
		fmt.Println("┌───────────────────────────────────┐")
		fmt.Print("│       BOUTIQUE DU MARCHAND        │")
		fmt.Print("\n├───────────────────────────────────┤\n")
		fmt.Print("│ Bienvenue dans ma boutique        │\n")
		fmt.Print("│ Voici les items disponibles :     │\n")
		fmt.Print("│                                   │\n")
		if PotionGratuite {
			fmt.Print("│ 1. Potion de vie - GRATUIT        │\n")
		} else {
			prixActuel := prixBase * tour
			if prixActuel == 0 {
				prixActuel = 10
			}
			fmt.Printf("│ 1. Potion de vie - %-3d pièces d'or│\n", prixActuel)
		}
		fmt.Print("│ 2. Retourner au menu principal    │\n")
		fmt.Print("└───────────────────────────────────┘\n\nQue souhaitez vous faire ?\n")
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
					continue
				}
			} else {
				characterName = characters.C1.Nom
			}
			functionshelper.AddInventory(characterName, "potion de vie")
			if PotionGratuite {
				fmt.Printf("\nPotion de vie GRATUITE ajoutée à l'inventaire de %s !\n", characterName)
				PotionGratuite = false
			} else {
				fmt.Printf("\nPotion de vie ajoutée à l'inventaire de %s !\n", characterName)
			}
			fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "2", "2.":
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func ItemView(s string) {
	fmt.Print("Souhaitez-vous avoir le détail d'un objet ? (Oui/Non): ")
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
	input := functionshelper.ReadInput()
	input = strings.ToLower(strings.TrimSpace(input))
	if input == "oui" || input == "oui." || input == "o" || input == "yes" {
		ItemViewOui(character)
	} else {
		return
	}
}

func ItemViewOui(character *characters.Character) {
	displayInventoryForSelection(character)
	fmt.Print("De quel objet souhaitez-vous avoir le détail ? \n(Veuillez mettre le NOM de l'objet) \n") // Sélection de l'item
	item := strings.TrimSpace(functionshelper.ReadInput())
	if item == "" {
		fmt.Print("Aucun item saisi.\n")
		return
	}
	itemFound := false
	for _, inventoryItem := range character.Inventaire { // Recherche de l'item dans l'inventaire
		itemStr := fmt.Sprintf("%v", inventoryItem)
		if itemStr == item {
			functionshelper.DisplayItemDetails(itemStr)
			itemFound = true
			break
		}
	}

	if !itemFound {
		fmt.Printf("L'item '%s' n'est pas dans l'inventaire.\n", item)
	}
}

// Fonction pour afficher l'inventaire avant sélection
func displayInventoryForSelection(character *characters.Character) {
	if character == nil {
		fmt.Println("Erreur : joueur non initialisé")
		return
	}
	fmt.Printf("\n┌─────────────────────────────────────────────────┐\n")
	fmt.Printf("│ Inventaire de %-33s │\n", character.Nom)
	fmt.Printf("├─────────────────────────────────────────────────┤\n")

	if len(character.Inventaire) == 0 {
		fmt.Printf("│ Inventaire vide                                 │\n")
	} else {
		for i, item := range character.Inventaire {
			itemStr := fmt.Sprintf("%v", item)
			fmt.Printf("│ %d. %-44s │\n", i+1, itemStr)
		}
	}

	fmt.Printf("└─────────────────────────────────────────────────┘\n")
}
