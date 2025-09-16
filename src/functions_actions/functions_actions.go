package functionsactions

import (
	characters "ASCII_Aventure/characters"
	"ASCII_Aventure/couleurs"
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
		fmt.Printf("\n%s=== LISTE DES PERSONNAGES ===%s\n", couleurs.Bold+couleurs.Cyan, couleurs.Reset)
		fmt.Printf("\n%sPersonnage 1 :%s\n", couleurs.Green, couleurs.Reset)
		characters.DisplayCharacterTable(*characters.C1)
		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("\n%sPersonnage 2 :%s\n", couleurs.Green, couleurs.Reset)
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
	fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
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
			fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
			return
		}
	} else {
		character = characters.C1
	}
	inventoryCount := len(character.Inventaire)
	fmt.Printf("\n%s┌─────────────────────────────────────────────────┐%s\n", couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s %sInventaire de%s %s%-33s%s %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Green, character.Nom, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s %s(%d/%d items)%s                                    %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Blue, inventoryCount, character.InventaireMaxSlots, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────┤%s\n", couleurs.Cyan, couleurs.Reset)
	if inventoryCount == 0 {
		fmt.Printf("%s│%s %sInventaire vide%s                                 %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Red, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	} else {
		for i, item := range character.Inventaire {
			itemName := fmt.Sprintf("%v", item)
			fmt.Printf("%s│%s %s%d.%s %-44s %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Yellow, i+1, couleurs.Reset, itemName, couleurs.Cyan, couleurs.Reset)
		}
	}
	fmt.Printf("%s└─────────────────────────────────────────────────┘%s\n", couleurs.Cyan, couleurs.Reset)
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
			fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
			return false
		}
	} else {
		character = characters.C1
	}
	nombreObjets := len(character.Inventaire)
	if nombreObjets < character.InventaireMaxSlots {
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
			fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
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
		fmt.Printf("\n%s%s n'a pas de potion dans son inventaire.%s\n", couleurs.Red, character.Nom, couleurs.Reset)
		return
	}
	character.PointsDeVieActuels += 50
	if character.PointsDeVieActuels > character.PointsDeVieMaximum {
		character.PointsDeVieActuels = character.PointsDeVieMaximum
	}
	character.Inventaire = functionshelper.RemoveItem(character.Inventaire, potionIndex)
	fmt.Printf("\n%s%s a utilisé une %s ! (+50 PV)%s\n", couleurs.Green, character.Nom, potionType, couleurs.Reset)
	fmt.Printf("%sPoints de vie: %s%d/%d%s\n", couleurs.White, couleurs.Green, character.PointsDeVieActuels, character.PointsDeVieMaximum, couleurs.Reset)
}

func Marchand() {
	prixActuel_PotionDeVie := 10
	prixActuel_PotionDePoison := 15
	prixActuel_SpellBookBouleDeFeu := 25
	prixActuel_Fourruredeloup := 4
	prixActuel_Peaudetroll := 7
	prixActuel_Cuirdesanglier := 3
	prixActuel_Plumedecorbeau := 1
	for {
		fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s%s                                BOUTIQUE DU MARCHAND                                 %s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %sBienvenue dans ma boutique%s                                                          %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %sVoici les items disponibles :%s                                                       %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		if PotionGratuite {
			fmt.Printf("%s│%s %s1.%s Potion de soin - %sGRATUIT%s                                                         %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Bold+couleurs.Yellow, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		} else {
			fmt.Printf("%s│%s %s1.%s Potion de soin - %s%-3d pièces d'or%s                                                 %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_PotionDeVie, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		}
		fmt.Printf("%s│%s %s2.%s Potion de poison - %s%-3d pièces d'or%s                                               %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_PotionDePoison, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s3.%s [Spell book] > Boule de feu - %s%-3d pièces d'or%s                                    %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_SpellBookBouleDeFeu, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s4.%s Fourrure de loup - %s%-3d pièces d'or%s                                               %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_Fourruredeloup, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s5.%s Peau de troll - %s%-3d pièces d'or%s                                                  %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_Peaudetroll, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s6.%s Cuir de sanglier - %s%-3d pièces d'or%s                                               %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_Cuirdesanglier, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s7.%s Plume de corbeau - %s%-3d pièces d'or%s                                               %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Yellow, prixActuel_Plumedecorbeau, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s8.%s Revendre des objets                                                              %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s8.%s Retourner au menu principal                                                      %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Red, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Cyan, couleurs.Reset)
		fmt.Printf("\n%sQue souhaitez vous faire ?%s\n", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		choix := functionshelper.ReadInput()
		switch choix {
		case "1", "1.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la potion ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if PotionGratuite {
				fmt.Printf("\n%sPotion de vie GRATUITE ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "potion de soin")
				PotionGratuite = false
			} else {
				if character.PiècesDOr >= prixActuel_PotionDeVie {
					fmt.Printf("\n%sPotion de vie ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
					functionshelper.AddInventory(characterName, "potion de soin")
					character.PiècesDOr -= prixActuel_PotionDeVie
					fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
				} else {
					manque := prixActuel_PotionDeVie - character.PiècesDOr
					fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
				}
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "2", "2.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la potion ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_PotionDePoison {
				fmt.Printf("\n%sPotion de poison ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "potion de poison")
				character.PiècesDOr -= prixActuel_PotionDePoison
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_PotionDePoison - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "3", "3.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter le livre de sort ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_SpellBookBouleDeFeu {
				fmt.Printf("\n%s[Spell book] > Boule de feu ajouté à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "[Spell book] > Boule de feu")
				character.PiècesDOr -= prixActuel_SpellBookBouleDeFeu
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_SpellBookBouleDeFeu - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "4", "4.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la fourrure de loup ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_Fourruredeloup {
				fmt.Printf("\n%sFourrure de loup ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "Fourrure de loup")
				character.PiècesDOr -= prixActuel_Fourruredeloup
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_Fourruredeloup - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "5", "5.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la peau de troll ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_Peaudetroll {
				fmt.Printf("\n%sPeau de troll ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "Peau de troll")
				character.PiècesDOr -= prixActuel_Peaudetroll
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_Peaudetroll - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "6", "6.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter le cuir de sanglier ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_Cuirdesanglier {
				fmt.Printf("\n%sCuir de sanglier ajouté à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "Cuir de sanglier")
				character.PiècesDOr -= prixActuel_Cuirdesanglier
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_Cuirdesanglier - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "7", "7.":
			var characterName string
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la plume de corbeau ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					character = characters.C1
				case "2", "2.":
					characterName = characters.C2.Nom
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				characterName = characters.C1.Nom
				character = characters.C1
			}
			if !CheckInventorySpace(characterName) {
				fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
				fmt.Printf("%sVous devez libérer de l'espace avant d'acheter un nouvel item.%s\n", couleurs.Red, couleurs.Reset)
				continue
			}
			if character.PiècesDOr >= prixActuel_Plumedecorbeau {
				fmt.Printf("\n%sPlume de corbeau ajoutée à l'inventaire de %s !%s\n", couleurs.Green, characterName, couleurs.Reset)
				functionshelper.AddInventory(characterName, "Plume de corbeau")
				character.PiècesDOr -= prixActuel_Plumedecorbeau
				fmt.Printf("%sVotre nouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
			} else {
				manque := prixActuel_Plumedecorbeau - character.PiècesDOr
				fmt.Printf("\n%sVous n'avez pas assez de pièces d'or (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
			}
			fmt.Printf("\n%sVoulez-vous acheter autre chose ? (o/n) :%s ", couleurs.Blue, couleurs.Reset)
			continuer := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
			if continuer != "o" && continuer != "oui" {
				return
			}
		case "8", "8.":
			var characterName string
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite revendre ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					characterName = characters.C1.Nom
					AccessInventory(characterName)
					ItemView(characterName)
				case "2", "2.":
					characterName = characters.C2.Nom
					AccessInventory(characterName)
					ItemView(characterName)
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				AccessInventory(characters.C1.Nom)
				ItemView(characters.C1.Nom)
				fmt.Printf("\n%sQuel(s) objet(s) souhaitez vous revendre ?%s\n", couleurs.Blue, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				itemName := functionshelper.ReadInput()
				for len(itemName) <= 6 {
					if len(itemName) <= 6 {
						fmt.Printf("%sErreur : vous devez entrer le nom d'un objet %s\n", couleurs.Red, couleurs.Reset)
						fmt.Printf("\n%sQuel(s) objet(s) souhaitez vous revendre ?%s\n", couleurs.Blue, couleurs.Reset)
					}
					itemName = functionshelper.ReadInput()
					functionshelper.RemoveInventory(characters.C1.Nom, itemName)
					if len(itemName) >= 6 {
						fmt.Printf("%sVous avez bien vendu %s%s%s ! %s\n", couleurs.Green, couleurs.Yellow, itemName, couleurs.Green, couleurs.Reset)
					}
				}
			}
		case "9", "9.":
			return
		default:
			fmt.Printf("%sChoix invalide, veuillez réessayer.%s\n", couleurs.Red, couleurs.Reset)
		}
	}
}

func Forgeron() {
	for {
		fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s%s                                FORGE DE L'AVENTURIER                                %s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %sBienvenue dans ma forge !%s                                                           %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %sVoici les équipements que je peux fabriquer :%s                                       %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s1.%s Chapeau de l'aventurier - Nécessite : 1 Plume de Corbeau + 1 Cuir de Sanglier    %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s2.%s Tunique de l'aventurier - Nécessite : 2 Fourrure de loup + 1 Peau de Troll       %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s3.%s Bottes de l'aventurier - Nécessite : 1 Fourrure de loup + 1 Cuir de Sanglier     %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s4.%s Amélioration d'inventaire                                                        %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s %s5.%s Retour au menu                                                                   %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Red, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Cyan, couleurs.Reset)

		fmt.Printf("\n%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		choice := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))

		switch choice {
		case "1", "1.":
			forgerItem("Chapeau de l'aventurier", []string{"Plume de corbeau", "Cuir de sanglier"}, []int{1, 1})
		case "2", "2.":
			forgerItem("Tunique de l'aventurier", []string{"Fourrure de loup", "Peau de troll"}, []int{2, 1})
		case "3", "3.":
			forgerItem("Bottes de l'aventurier", []string{"Fourrure de loup", "Cuir de sanglier"}, []int{1, 1})
		case "4", "4.":
			var character *characters.Character
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sQuel personnage souhaite acheter la peau de troll ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
				fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)
				choixPerso := functionshelper.ReadInput()
				switch choixPerso {
				case "1", "1.":
					character = characters.C1
				case "2", "2.":
					character = characters.C2
				default:
					fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
					continue
				}
			} else {
				character = characters.C1
			}
			characters.UpgradeInventorySlot(character)
		case "5", "5.":
			return
		default:
			fmt.Printf("%sChoix invalide, veuillez réessayer%s\n", couleurs.Red, couleurs.Reset)
		}
	}
}

func forgerItem(equipmentName string, materials []string, quantities []int) {
	var characterName string
	var character *characters.Character
	if characters.C2_b && characters.C2 != nil {
		fmt.Printf("\n%sQuel personnage souhaite forger %s ?%s\n", couleurs.Purple, equipmentName, couleurs.Reset)
		fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
		fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
		fmt.Printf("%sVotre choix :%s ", couleurs.Blue, couleurs.Reset)

		choixPerso := functionshelper.ReadInput()
		switch choixPerso {
		case "1", "1.":
			characterName = characters.C1.Nom
			character = characters.C1
		case "2", "2.":
			characterName = characters.C2.Nom
			character = characters.C2
		default:
			fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
			return
		}
	} else {
		characterName = characters.C1.Nom
		character = characters.C1
	}
	if character.PiècesDOr < 5 {
		manque := 5 - character.PiècesDOr
		fmt.Printf("\n%sVous n'avez pas assez de pièces d'or pour forger ! (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
		return
	}
	if !CheckInventorySpace(characterName) {
		fmt.Printf("\n%sL'inventaire de %s est plein ! (10/10 items)%s\n", couleurs.Red, characterName, couleurs.Reset)
		fmt.Printf("%sVous devez libérer de l'espace avant de forger un équipement.%s\n", couleurs.Red, couleurs.Reset)
		return
	}
	for _, item := range character.Inventaire {
		if strings.ToLower(item) == strings.ToLower(equipmentName) {
			fmt.Printf("\n%s%s possède déjà %s !%s\n", couleurs.Yellow, characterName, equipmentName, couleurs.Reset)
			return
		}
	}
	missingMaterials := []string{}
	materialCount := make(map[string]int)
	for _, item := range character.Inventaire {
		itemLower := strings.ToLower(item)
		materialCount[itemLower]++
	}
	for i, material := range materials {
		materialLower := strings.ToLower(material)
		needed := quantities[i]
		possessed := materialCount[materialLower]

		if possessed < needed {
			missing := needed - possessed
			missingMaterials = append(missingMaterials, fmt.Sprintf("%d %s", missing, material))
		}
	}
	if len(missingMaterials) > 0 {
		fmt.Printf("\n%sMatériaux insuffisants pour forger %s !%s\n", couleurs.Red, equipmentName, couleurs.Reset)
		fmt.Printf("%sIl vous manque : %s%s%s\n", couleurs.Red, couleurs.Yellow, strings.Join(missingMaterials, ", "), couleurs.Reset)
		return
	}
	for i, material := range materials {
		materialLower := strings.ToLower(material)
		needed := quantities[i]
		removed := 0
		for j := len(character.Inventaire) - 1; j >= 0 && removed < needed; j-- {
			if strings.ToLower(character.Inventaire[j]) == materialLower {
				character.Inventaire = functionshelper.RemoveItem(character.Inventaire, j)
				removed++
			}
		}
	}
	character.PiècesDOr -= 5
	functionshelper.AddInventory(characterName, equipmentName)
	fmt.Printf("\n%s%s a forgé avec succès : %s%s%s%s\n", couleurs.Green, characterName, couleurs.Yellow, equipmentName, couleurs.Green, couleurs.Reset)
	fmt.Printf("%sMatériaux utilisés :%s\n", couleurs.White, couleurs.Reset)
	for i, material := range materials {
		fmt.Printf("%s- %d %s%s\n", couleurs.Yellow, quantities[i], material, couleurs.Reset)
	}
	fmt.Printf("%sPièces d'or dépensées : %s5%s\n", couleurs.White, couleurs.Yellow, couleurs.Reset)
	fmt.Printf("%sNouveau solde : %s%d pièces d'or%s\n", couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
}

func ItemView(s string) {
	s = strings.TrimSpace(strings.ToLower(s))
	fmt.Printf("%sSouhaitez-vous avoir le détail d'un objet ? (Oui/Non):%s ", couleurs.Blue, couleurs.Reset)
	var character *characters.Character
	if characters.C2_b {
		switch s {
		case strings.ToLower(characters.C1.Nom):
			character = characters.C1
		case strings.ToLower(characters.C2.Nom):
			character = characters.C2
		default:
			fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
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
	fmt.Printf("%sDe quel objet souhaitez-vous avoir le détail ?%s\n", couleurs.Purple, couleurs.Reset)
	fmt.Printf("%s(Veuillez mettre le NOM de l'objet)%s\n", couleurs.White, couleurs.Reset)
	item := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
	if item == "" {
		fmt.Printf("%sAucun item saisi.%s\n", couleurs.Red, couleurs.Reset)
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
		fmt.Printf("%sL'item '%s' n'est pas dans l'inventaire.%s\n", couleurs.Red, item, couleurs.Reset)
	}
}
