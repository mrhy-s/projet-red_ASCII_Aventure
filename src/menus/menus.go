package menus

import (
	combat "ASCII_Aventure/Combat"
	characters "ASCII_Aventure/characters"
	"ASCII_Aventure/couleurs"
	functionsactions "ASCII_Aventure/functions_actions"
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/savegame"
	"ASCII_Aventure/startscreen"
	"fmt"
	"strings"
)

// ==============
// fonction menus
// ==============

func Menu() {
	if characters.C2_b && characters.C2 != nil {
		savegame.SaveCharacter(characters.C1)
		savegame.SaveCharacter(characters.C2)
	} else {
		savegame.SaveCharacter(characters.C1)
	}
	characters.IsDead()
	fmt.Printf("\n")
	fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s%s                                  MENU PRINCIPAL                                     %s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s1.%s Afficher les informations du personnage                                          %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s2.%s Accéder au contenu de l'inventaire                                               %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s3.%s Utiliser une potion de soin                                                      %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s4.%s Boutique du Marchand                                                             %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s5.%s Boutique du Forgeron                                                             %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s6.%s Créer un nouveau personnage                                                      %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s7.%s Rechercher un ennemi à attaquer                                                  %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s8.%s Passer un tour                                                                   %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s9.%s Quitter                                                                          %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Red, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("\n%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
	option := functionshelper.ReadInput()
	switch option {
	case "1", "1.":
		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("%sVeuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :%s\n", couleurs.Purple, couleurs.Reset)
			fmt.Printf("%sPersonnages disponibles:%s %s%s%s, %s%s%s\n", couleurs.White, couleurs.Reset, couleurs.Green, characters.C1.Nom, couleurs.Reset, couleurs.Green, characters.C2.Nom, couleurs.Reset)
			fmt.Printf("\n%sVotre réponse :%s ", couleurs.Blue, couleurs.Reset)
			commande := functionshelper.ReadInput()
			functionsactions.DisplayInfo(commande)
			if strings.Contains(commande, characters.C1.Nom) {
				characters.DisplayEquipment(characters.C1)
			} else if strings.Contains(commande, characters.C2.Nom) {
				characters.DisplayEquipment(characters.C2)
			}
		} else {
			functionsactions.DisplayInfo(characters.C1.Nom)
			characters.DisplayEquipment(characters.C1)
		}
		Menu()
	case "2", "2.":
		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("%sVeuillez sélectionner un personnage de la liste pour afficher son inventaire :%s\n", couleurs.Purple, couleurs.Reset)
			fmt.Printf("%sPersonnages disponibles:%s %s%s%s, %s%s%s\n", couleurs.White, couleurs.Reset, couleurs.Green, characters.C1.Nom, couleurs.Reset, couleurs.Green, characters.C2.Nom, couleurs.Reset)
			fmt.Printf("\n%sVotre réponse :%s ", couleurs.Blue, couleurs.Reset)
			commande := functionshelper.ReadInput()
			functionsactions.AccessInventory(commande)
			functionsactions.ItemView(commande)
		} else {
			functionsactions.AccessInventory(characters.C1.Nom)
			functionsactions.ItemView(characters.C1.Nom)
		}
		Menu()
	case "3", "3.":
		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("%sNom du personnage :%s ", couleurs.Purple, couleurs.Reset)
			characterName := functionshelper.ReadInput()
			functionsactions.TakePot(characterName)
		} else {
			functionsactions.TakePot(characters.C1.Nom)
		}
		Menu()
	case "4", "4.":
		startscreen.ClearScreen()
		functionsactions.Marchand()
		Menu()
	case "5", "5.":
		startscreen.ClearScreen()
		functionsactions.Forgeron()
		Menu()
	case "6", "6.":
		characters.C2 = functionshelper.CharacterCreation()
		Menu()
	case "7", "7.":
		startscreen.ClearScreen()
		combat.RechercheEnemy()
		combat.Combat(functionsactions.Tour)
		Menu()
	case "8", "8.":
		startscreen.ClearScreen()
		functionsactions.Tour++
		Menu()
	case "9", "9.":
		return
	default:
		startscreen.ClearScreen()
		Menu()
		fmt.Printf("%sOption invalide%s\n", couleurs.Red, couleurs.Reset)
	}
}
