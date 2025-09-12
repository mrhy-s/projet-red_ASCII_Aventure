package main

import (
	"ASCII_Aventure/characters"
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/menus"
	"fmt"
)

func main() {
	IsFirst := true
	characters.C1 = characters.InitCharacter("Zinrel", "Elfe", 1, 100, 40, []string{"potion de soin", "potion de soin", "potion de soin"})

	if IsFirst {
		min_pt_vie_rand, max_pt_vie_rand := 95, 105
		fmt.Print("Voulez vous utiliser le personnage par défaut ? Ou voulez vous créer un nouveau personnage ?\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'\n\n")
		characters.DisplayCharacterTable(*characters.C1)
		fmt.Print("\nVotre réponse : ")
		input := functionshelper.ReadInput()
		switch input {
		case "Oui", "Oui.", "oui", "oui.":
			fmt.Print("\nVeuillez entrer le nom du personnage : ")
			characters.C_temp_name = functionshelper.ReadInput()
			fmt.Print("\nVeuillez entrer la classe du personnage : ")
			characters.C_temp_classe = functionshelper.ReadInput()
			fmt.Printf("\nVotre personnage s'appelle : %s\nLa classe de votre personnage est : %s", characters.C_temp_name, characters.C_temp_classe)
			fmt.Print("\nVotre personnage commence au niveau 0.")
			characters.C_temp_niveau = 0
			characters.C_temp_points_de_vie_maximum = functionshelper.RandomBetween(min_pt_vie_rand, max_pt_vie_rand)
			characters.C_temp_points_de_vie_actuels = characters.C_temp_points_de_vie_maximum - functionshelper.RandomBetween(20, 70)
			fmt.Printf("\nVotre personnage avec %d/%d points de vie.", characters.C_temp_points_de_vie_actuels, characters.C_temp_points_de_vie_maximum)
			characters.C_temp_inventaire = []string{"potion de soin", "potion de soin", "potion de soin"}
			fmt.Print("\nvous commencez avec 'Potion de soin x3'\n")
			characters.C2_b = true
			characters.C2 = characters.InitCharacter(characters.C_temp_name, characters.C_temp_classe, characters.C_temp_niveau, characters.C_temp_points_de_vie_maximum, characters.C_temp_points_de_vie_actuels, characters.C_temp_inventaire)
			fmt.Println("\nVoici votre nouveau personnage :")
			characters.DisplayCharacterTable(*characters.C2)
		case "Non", "Non.", "non", "non.":
			menus.Menu()
			/*
				case "DEBUG:kill": // DEBUG
					fmt.Print("\nEntrer un personnage a tuer \nVotre réponse : ") // DEBUG
					input := functionshelper.ReadInput()                          // DEBUG
					functionshelper.InstaKill(input)
			*/ // DEBUG
		default:
			fmt.Printf("On a dit 'oui' ou 'non' pas : %s (╥﹏╥)\n\n", input)
			main()
		}
		IsFirst = false
	}
	menus.Menu()
}

// faire de sorte a ce que le marchand ne propose la potion gratuite qu'une fois
// faire de sorte a ce que quand on accède a l'inventaire on puisse avoir accès aux informations de l'item
