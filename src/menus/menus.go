package menus

import (
	"ASCII_Aventure/characters"
	functionsactions "ASCII_Aventure/functions_actions"
	functionshelper "ASCII_Aventure/functions_helper"
	"fmt"
)

// ==============
// fonction menus
// ==============

func Menu() {
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
	option := functionshelper.ReadInput()
	switch option {
	case "1", "1.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :\n")

		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("Personnages disponibles: %s, %s\n", characters.C1.Nom, characters.C2.Nom)
		} else {
			fmt.Printf("Personnage disponible: %s\n", characters.C1.Nom)
		}
		fmt.Print("\nVotre réponse : ")
		commande := functionshelper.ReadInput()
		functionsactions.DisplayInfo(commande)
		Menu()
	case "2", "2.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher son inventaire :\n")
		if characters.C2_b && characters.C2 != nil {
			fmt.Printf("Personnages disponibles: %s, %s\n", characters.C1.Nom, characters.C2.Nom)
		} else {
			fmt.Printf("Personnage disponible: %s\n", characters.C1.Nom)
		}
		fmt.Print("\nVotre réponse : ")
		commande := functionshelper.ReadInput()
		functionsactions.AccessInventory(commande)
		functionsactions.ItemView(commande)
		Menu()
	case "3", "3.":
		fmt.Print("Nom du personnage : ")
		characterName := functionshelper.ReadInput()
		functionsactions.TakePot(characterName)
		Menu()
	case "4", "4.":
		functionsactions.Marchand(functionsactions.Tour)
		Menu()
	case "5", "5.":
		return
	default:
		fmt.Println("Option invalide")
		Menu()
	}
}
