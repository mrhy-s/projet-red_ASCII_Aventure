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
	s = strings.ToLower(strings.TrimSpace(s))
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
	if s == strings.ToLower(characters.C1.Nom) {
		characters.DisplayCharacterTable(*characters.C1)
		return
	}
	if characters.C2_b && characters.C2 != nil && s == strings.ToLower(characters.C2.Nom) {
		characters.DisplayCharacterTable(*characters.C2)
		return
	}
	fmt.Println("Personnage non trouvé")
}

func AccessInventory(s string) {
	var character *characters.Character
	s = strings.TrimSpace(strings.ToLower(s))
	if characters.C2_b {
		switch s {
		case strings.ToLower(characters.C1.Nom):
			character = characters.C1
		case strings.ToLower(characters.C2.Nom):
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = characters.C1
	}
	inventoryCount := len(character.Inventaire)
	fmt.Printf("\n┌─────────────────────────────────────────────────┐\n")
	fmt.Printf("│ Inventaire de %-33s │\n", character.Nom)
	fmt.Printf("│ (%d/10 items)                                    │\n", inventoryCount)
	fmt.Printf("├─────────────────────────────────────────────────┤\n")
	if inventoryCount == 0 {
		fmt.Printf("│ Inventaire vide                                 │\n")
	} else {
		for i, item := range character.Inventaire {
			itemName := fmt.Sprintf("%v", item)
			fmt.Printf("│ %d. %-44s │\n", i+1, itemName)
		}
	}
	fmt.Printf("└─────────────────────────────────────────────────┘\n")
}

func CheckInventorySpace(characterName string) bool {
	var character *characters.Character
	if characters.C2_b {
		switch characterName {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return false
		}
	} else {
		character = characters.C1
	}
	nombreObjets := len(character.Inventaire)
	if nombreObjets < 10 {
		return true
	} else {
		return false
	}
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
	var prixActuel_PotionDeVie int
	prixActuel_PotionDePoison := (prixBase + 5) * tour
	if prixActuel_PotionDePoison == 0 {
		prixActuel_PotionDePoison = 15
	}
	prixActuel_SpellBookBouleDeFeu := (prixBase + 15) * tour
	if prixActuel_SpellBookBouleDeFeu == 0 {
		prixActuel_SpellBookBouleDeFeu = 25
	}
	for {
		fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
		fmt.Print("│                                BOUTIQUE DU MARCHAND                                 │")
		fmt.Print("\n├─────────────────────────────────────────────────────────────────────────────────────┤\n")
		fmt.Print("│ Bienvenue dans ma boutique                                                          │\n")
		fmt.Print("│ Voici les items disponibles :                                                       │\n")
		fmt.Print("│                                                                                     │\n")
		if PotionGratuite {
			fmt.Print("│ 1. Potion de soin - GRATUIT                                                         │\n")
		} else {
			prixActuel_PotionDeVie = prixBase * tour
			if prixActuel_PotionDeVie == 0 {
				prixActuel_PotionDeVie = 10
			}
			fmt.Printf("│ 1. Potion de soin - %-3d pièces d'or                                                 │\n", prixActuel_PotionDeVie)
		}
		fmt.Printf("│ 2. Potion de poison - %-3d pièces d'or                                               │\n", prixActuel_PotionDePoison)
		fmt.Printf("│ 3. [Spell book] > Boule de feu - %-3d pièces d'or                                    │\n", prixActuel_SpellBookBouleDeFeu)
		fmt.Print("│ 4. Retourner au menu principal                                                      │\n")
		fmt.Print("└─────────────────────────────────────────────────────────────────────────────────────┘\n\nQue souhaitez vous faire ?\n")
		choix := functionshelper.ReadInput()
		switch choix {
		case "1", "1.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Print("\nQuel personnage souhaite acheter la potion ?\n")
				fmt.Printf("1. %s\n", characters.C1.Nom)
				fmt.Printf("2. %s\n", characters.C2.Nom)
				fmt.Print("Votre choix : ")
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Println("Choix invalide")
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\nL'inventaire de %s est plein ! (10/10 items)\n", characterName)
				fmt.Print("Vous devez libérer de l'espace avant d'acheter un nouvel item.\n")
				continue
			}
			if PotionGratuite {
				fmt.Printf("\nPotion de vie GRATUITE ajoutée à l'inventaire de %s !\n", characterName)
				functionshelper.AddInventory(characterName, "potion de soin")
				PotionGratuite = false
			} else {
				if character.PiècesDOr >= prixActuel_PotionDeVie {
					fmt.Printf("\nPotion de vie ajoutée à l'inventaire de %s !\n", characterName)
					functionshelper.AddInventory(characterName, "potion de soin")
					character.PiècesDOr -= prixActuel_PotionDeVie
					fmt.Printf("Votre nouveau solde : %d pièces d'or\n", character.PiècesDOr)
				} else {
					manque := prixActuel_PotionDeVie - character.PiècesDOr
					fmt.Printf("\nVous n'avez pas assez de pièces d'or (il vous manque %d pièces)\n", manque)
				}
			}
			fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "2", "2.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Print("\nQuel personnage souhaite acheter la potion ?\n")
				fmt.Printf("1. %s\n", characters.C1.Nom)
				fmt.Printf("2. %s\n", characters.C2.Nom)
				fmt.Print("Votre choix : ")
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Println("Choix invalide")
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\nL'inventaire de %s est plein ! (10/10 items)\n", characterName)
				fmt.Print("Vous devez libérer de l'espace avant d'acheter un nouvel item.\n")
				continue
			}
			if character.PiècesDOr >= prixActuel_PotionDePoison {
				fmt.Printf("\nPotion de poison ajoutée à l'inventaire de %s !\n", characterName)
				functionshelper.AddInventory(characterName, "potion de poison")
				character.PiècesDOr -= prixActuel_PotionDePoison
				fmt.Printf("Votre nouveau solde : %d pièces d'or\n", character.PiècesDOr)
			} else {
				manque := prixActuel_PotionDePoison - character.PiècesDOr
				fmt.Printf("\nVous n'avez pas assez de pièces d'or (il vous manque %d pièces)\n", manque)
			}
			fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "3", "3.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Print("\nQuel personnage souhaite acheter le livre de sort ?\n")
				fmt.Printf("1. %s\n", characters.C1.Nom)
				fmt.Printf("2. %s\n", characters.C2.Nom)
				fmt.Print("Votre choix : ")
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Println("Choix invalide")
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\nL'inventaire de %s est plein ! (10/10 items)\n", characterName)
				fmt.Print("Vous devez libérer de l'espace avant d'acheter un nouvel item.\n")
				continue
			}
			if character.PiècesDOr >= prixActuel_SpellBookBouleDeFeu {
				fmt.Printf("\nPotion de poison ajoutée à l'inventaire de %s !\n", characterName)
				functionshelper.AddInventory(characterName, "[Spell book] > Boule de feu")
				character.PiècesDOr -= prixActuel_SpellBookBouleDeFeu
				fmt.Printf("Votre nouveau solde : %d pièces d'or\n", character.PiècesDOr)
			} else {
				manque := prixActuel_SpellBookBouleDeFeu - character.PiècesDOr
				fmt.Printf("\nVous n'avez pas assez de pièces d'or (il vous manque %d pièces)\n", manque)
			}
			fmt.Print("\nVoulez-vous acheter autre chose ? (o/n) : ")
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "4", "4.":
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func ItemView(s string) {
	s = strings.TrimSpace(strings.ToLower(s))
	fmt.Print("Souhaitez-vous avoir le détail d'un objet ? (Oui/Non): ")
	var character *characters.Character
	if characters.C2_b {
		switch s {
		case strings.ToLower(characters.C1.Nom):
			character = characters.C1
		case strings.ToLower(characters.C2.Nom):
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
	fmt.Print("De quel objet souhaitez-vous avoir le détail ? \n(Veuillez mettre le NOM de l'objet) \n")
	item := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
	if item == "" {
		fmt.Print("Aucun item saisi.\n")
		return
	}
	itemFound := false
	for _, inventoryItem := range character.Inventaire {
		itemStr := strings.ToLower(fmt.Sprintf("%v", inventoryItem))
		if itemStr == item {
			functionshelper.DisplayItemDetails(inventoryItem)
			itemFound = true
			break
		}
	}
	if !itemFound {
		fmt.Printf("L'item '%s' n'est pas dans l'inventaire.\n", item)
	}
}
