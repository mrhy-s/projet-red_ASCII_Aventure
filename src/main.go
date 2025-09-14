package main

import (
	"ASCII_Aventure/characters"
	"ASCII_Aventure/classes"
	functionsactions "ASCII_Aventure/functions_actions"
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/items"
	"ASCII_Aventure/menus"
	"ASCII_Aventure/skills"
	"ASCII_Aventure/startscreen"
	"fmt"
)

var IsFirst bool

func main() {
	startScreen := true
	functionsactions.PotionGratuite = true
	IsFirst = true
	// initialisation
	characters.C1 = characters.InitCharacter("Zinrel", "Elfe", 1, 80, 40, []string{"potion de soin", "potion de soin", "potion de soin"}, []string{"Coup de poing"}, 100)
	items.Potion_de_soin = items.InitPotion("potion de soin", "any", 0, "Une potion magique qui soigne les blessures et restaure 50 points de vie", "heal:50PV")
	items.Potion_de_poison = items.InitPotion("potion de poison", "any", 0, "Une potion toxique qui inflige des dégâts", "poison:20PV")
	classes.Humain = classes.InitClasse("Humain", "Polyvalent et adaptable")
	classes.Elfe = classes.InitClasse("Elfe", "Agile et magique")
	classes.Nain = classes.InitClasse("Nain", "Robuste et artisan")
	skills.CoupDePoing = skills.InitSkill("Coup de poing", "any", 1, "Frappe d'un coup de poing puissant")
	skills.CoupDeHache = skills.InitSkill("Coup de hache", "Nain", 1, "Frappe d'un coup de hache puissant")
	skills.TirÀLarc = skills.InitSkill("Tir à l'arc", "Elfe", 1, "Tir une flèche puissante")
	if startScreen { // affichage de l'écran de démarrage (une seule fois)
		startscreen.StartScreen()
		startScreen = false
	}
	characterSelection() // sélection du personnage
}

func characterSelection() {
	for {
		fmt.Print("Voulez vous utiliser le personnage par défaut ? Ou voulez vous créer un nouveau personnage ?\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'\n\n")
		if IsFirst {
			characters.DisplayCharacterTable(*characters.C1)
			IsFirst = false
		}
		fmt.Print("\nVotre réponse : ")
		input := functionshelper.ReadInput()

		switch input {
		case "Oui", "Oui.", "oui", "oui.":
			characters.C2_b = true
			characters.C2 = functionshelper.CharacterCreation()
			menus.Menu()
			return
		case "Non", "Non.", "non", "non.":
			startscreen.ClearScreen()
			menus.Menu()
			return
		default:
			fmt.Printf("On a dit 'oui' ou 'non' pas : %s (╥﹏╥)\n\n", input)
		}
	}
}
