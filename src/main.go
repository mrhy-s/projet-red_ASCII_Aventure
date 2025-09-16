package main

import (
	"ASCII_Aventure/characters"
	"ASCII_Aventure/classes"
	functionsactions "ASCII_Aventure/functions_actions"
	"ASCII_Aventure/items"
	"ASCII_Aventure/menus"
	"ASCII_Aventure/savegame"
	"ASCII_Aventure/skills"
	"ASCII_Aventure/startscreen"
)

var IsFirst bool

func main() {
	savegame.LoadCharacter()
	startScreen := true
	functionsactions.PotionGratuite = true
	IsFirst = true
	// initialisation
	items.Cuir_de_sanglier = items.InitRessources("Cuir de sanglier", "du cuir récupéré sur un sanglier, pourrait être utile dans les mains d'une personne compétente")
	items.Peau_de_troll = items.InitRessources("Peau de troll", "la peau putride d'un troll, mieux vaut s'en debarasser au plus vite")
	items.Fourrure_de_loup = items.InitRessources("Fourrure de loup", "une fourrure de loup qui offre une certaine protection au froid ... et un style particulier !")
	items.Plume_de_corbeau = items.InitRessources("Plume de corbeau", "une plume de corbeau étrange dont l'utilité reste un mystère")
	classes.Goblin = classes.InitClasse("Goblin", "Les gobelins sont d'affreux petits humanoïdes mesurant à peine plus de 90 centimètres. \nLeur corps décharné est surmonté d'une tête démesurée généralement chauve, avec de grandes oreilles et des yeux rouges perçants, parfois jaunes.")
	characters.Goblin = characters.InitMonster("Goblin", "Goblin", 1, 100, 100, []string{}, []string{"Griffure"})
	skills.Griffure = skills.InitSkill("[Spell book] > Griffure", "Goblin", 0, "Inflige une griffure douloureuse")
	characters.C1 = characters.InitCharacter("Zinrel", "Elfe", 1, 80, 40, []string{"potion de soin", "potion de soin", "potion de soin"}, []string{"Coup de poing"}, 100)
	items.Potion_de_soin = items.InitPotion("potion de soin", "any", 0, "Une potion magique qui soigne les blessures et restaure 50 points de vie", "heal:50PV")
	items.Potion_de_poison = items.InitPotion("potion de poison", "any", 0, "Une potion toxique qui inflige des dégâts", "poison:20PV")
	classes.Humain = classes.InitClasse("Humain", "Polyvalent et adaptable")
	classes.Elfe = classes.InitClasse("Elfe", "Agile et magique")
	classes.Nain = classes.InitClasse("Nain", "Robuste et artisan")
	skills.CoupDePoing = skills.InitSkill("Coup de poing", "any", 1, "Frappe d'un coup de poing puissant")
	skills.CoupDeHache = skills.InitSkill("Coup de hache", "Nain", 1, "Frappe d'un coup de hache puissant")
	skills.TirÀLarc = skills.InitSkill("Tir à l'arc", "Elfe", 1, "Tir une flèche puissante")
	skills.BouleDeFeu = skills.InitSkill("[Spell book] > Boule de feu", "any", 1, "Tir une boule de feu")
	items.Bottes_de_laventurier = items.InitArmure("Bottes de l'aventurier", "any", 0, 80, "Bottes en cuir simples", 8)
	items.Tunique_de_laventurier = items.InitArmure("Tunique de l'aventurier", "any", 0, 100, "Tunique en cuir simples", 10)
	items.Chapeau_de_laventurier = items.InitArmure("Chapeau de l'aventurier", "any", 0, 70, "Chapeau en cuir simples", 5)
	if startScreen { // affichage de l'écran de démarrage (une seule fois)
		startscreen.StartScreen()
		menus.Menu()
		startScreen = false
	}
}
