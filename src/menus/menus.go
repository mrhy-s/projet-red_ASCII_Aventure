package menus

import (
	characters "ASCII_Aventure/characters"
	functionsactions "ASCII_Aventure/functions_actions"
	functionshelper "ASCII_Aventure/functions_helper"
	"fmt"
)

// ==============
// fonction menus
// ==============

func Menu() {
	characters.IsDead()
	fmt.Printf("\n")
	fmt.Printf("┌─────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│                                  MENU PRINCIPAL                                     │\n")
	fmt.Printf("├─────────────────────────────────────────────────────────────────────────────────────┤\n")
	fmt.Printf("│ 1. Afficher les informations du personnage                                          │\n")
	fmt.Printf("│ 2. Accéder au contenu de l'inventaire                                               │\n")
	fmt.Printf("│ 3. Utiliser une potion de soin                                                      │\n")
	fmt.Printf("│ 4. Boutique du Marchand                                                             │\n")
	fmt.Printf("│ 5. Passer un tour                                                                   │\n")
	fmt.Printf("│ 6. Quitter                                                                          │\n")
	fmt.Printf("└─────────────────────────────────────────────────────────────────────────────────────┘\n")
	fmt.Print("\nVotre choix : ")
	option := functionshelper.ReadInput()
	switch option {
	case "1", "1.":
		if characters.C2_b && characters.C2 != nil {
			fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :\n")
			fmt.Printf("Personnages disponibles: %s, %s\n", characters.C1.Nom, characters.C2.Nom)
			fmt.Print("\nVotre réponse : ")
			commande := functionshelper.ReadInput()
			functionsactions.DisplayInfo(commande)
		} else {
			functionsactions.DisplayInfo(characters.C1.Nom)
		}

		Menu()
	case "2", "2.":
		if characters.C2_b && characters.C2 != nil {
			fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher son inventaire :\n")
			fmt.Printf("Personnages disponibles: %s, %s\n", characters.C1.Nom, characters.C2.Nom)
			fmt.Print("\nVotre réponse : ")
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
			fmt.Print("Nom du personnage : ")
			characterName := functionshelper.ReadInput()
			functionsactions.TakePot(characterName)
		} else {
			functionsactions.TakePot(characters.C1.Nom)
		}
		Menu()
	case "4", "4.":
		functionsactions.Marchand(functionsactions.Tour)
		Menu()
	case "5", "5.":
		functionsactions.Tour++
		Menu()
	case "6", "6.":
		return
	default:
		fmt.Println("Option invalide")
		Menu()
	}
}
