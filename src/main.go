package main

import (
	"ASCII_Aventure/characters"
	"ASCII_Aventure/classes"
	functionsactions "ASCII_Aventure/functions_actions"
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/items"
	"ASCII_Aventure/menus"
	"ASCII_Aventure/savegame"
	"ASCII_Aventure/skills"
	"ASCII_Aventure/startscreen"
)

func main() {
	savegame.LoadCharacter()
	startScreen := true
	functionsactions.PotionGratuite = true
	// initialisation
	items.Fer_brut = items.InitRessources("Fer brut", "Minerai de fer brut, utilisé pour faire des équipements", 2)
	items.Cuir_de_sanglier = items.InitRessources("Cuir de sanglier", "du cuir récupéré sur un sanglier, pourrait être utile dans les mains d'une personne compétente", 2)
	items.Peau_de_troll = items.InitRessources("Peau de troll", "la peau putride d'un troll, mieux vaut s'en debarasser au plus vite", 5)
	items.Fourrure_de_loup = items.InitRessources("Fourrure de loup", "une fourrure de loup qui offre une certaine protection au froid ... et un style particulier !", 2)
	items.Plume_de_corbeau = items.InitRessources("Plume de corbeau", "une plume de corbeau étrange dont l'utilité reste un mystère", 0)
	classes.Gobelin = classes.InitClasse("Gobelin", "Les gobelins sont d'affreux petits humanoïdes mesurant à peine plus de 90 centimètres. \nLeur corps décharné est surmonté d'une tête démesurée généralement chauve, avec de grandes oreilles et des yeux rouges perçants, parfois jaunes.")
	characters.C1 = characters.InitCharacter("Zinrel", "Elfe", 1, 80, 40, []string{"potion de soin", "potion de soin", "potion de soin"}, []string{"Tir à l'arc"}, 100)
	items.Potion_de_soin = items.InitPotion("potion de soin", "any", 0, "Une potion magique qui soigne les blessures et restaure 50 points de vie", "heal:50PV", 8)
	items.Potion_de_poison = items.InitPotion("potion de poison", "any", 0, "Une potion toxique qui inflige des dégâts", "poison:20PV", 10)

	// classes humain
	classes.Humain = classes.InitClasse("Humain", "Polyvalent et adaptable")
	classes.Elfe = classes.InitClasse("Elfe", "Agile et magique")
	classes.Nain = classes.InitClasse("Nain", "Robuste et artisan")

	// skills monstres
	skills.Griffure = skills.InitSkill("[Spell book] > Griffure", "Gobelin", 0, "Inflige une griffure douloureuse", 14)
	skills.CoupDeRage = skills.InitSkill("Coup de Rage", "any", 1, "Entre dans une rage folle et frappe violemment", 15)
	skills.Soin = skills.InitSkill("Soin", "any", 1, "Utilise un sort de soin pour récupérer des PV", 0)

	// skills avanturier
	skills.CoupDePoing = skills.InitSkill("Coup de poing", "any", 1, "Frappe d'un coup de poing puissant", 12)
	skills.CoupDeHache = skills.InitSkill("Coup de hache", "Nain", 1, "Frappe d'un coup de hache puissant", 15)
	skills.TirÀLarc = skills.InitSkill("Tir à l'arc", "Elfe", 1, "Tir une flèche puissante", 12)
	skills.BouleDeFeu = skills.InitSkill("[Spell book] > Boule de feu", "any", 1, "Tir une boule de feu", 25)

	// armure de l'aventurier
	items.Bottes_de_laventurier = items.InitArmure("Bottes de l'aventurier", "any", 0, 80, 80, "Bottes en cuir simples", 8)
	items.Tunique_de_laventurier = items.InitArmure("Tunique de l'aventurier", "any", 0, 100, 100, "Tunique en cuir simples", 10)
	items.Chapeau_de_laventurier = items.InitArmure("Chapeau de l'aventurier", "any", 0, 70, 70, "Chapeau en cuir simples", 5)
	items.Épée_en_fer = items.InitArme("Épée en fer", "any", 1, 100, "Une épée solide forgée en fer, plus efficace que les armes de base", 15, 8)

	// dague rouillée - Lame ébrèchée mais encore tranchante
	dague_rouillée_durabilité_Actuelle := functionshelper.RandomBetween(8, 25)
	dague_rouillée_durabilité_Max := functionshelper.RandomBetween(25, 35)
	dague_rouillée_dégâts := functionshelper.RandomBetween(3, 8)
	items.Dague_rouillée = items.InitArmeMonster("Dague rouillée", "any", 1, dague_rouillée_durabilité_Actuelle, dague_rouillée_durabilité_Max, "Lame ébrèchée mais encore tranchante", dague_rouillée_dégâts, 3)

	// gourdin clouté - Bout de bois avec des clous tordus
	gourdin_clouté_durabilité_Actuelle := functionshelper.RandomBetween(12, 28)
	gourdin_clouté_durabilité_Max := functionshelper.RandomBetween(30, 45)
	gourdin_clouté_dégâts := functionshelper.RandomBetween(4, 10)
	items.Gourdin_clouté = items.InitArmeMonster("Gourdin clouté", "any", 1, gourdin_clouté_durabilité_Actuelle, gourdin_clouté_durabilité_Max, "Bout de bois avec des clous tordus", gourdin_clouté_dégâts, 4)

	// arc tordu - Arc mal façonné, flèches de fortune
	arc_tordu_durabilité_Actuelle := functionshelper.RandomBetween(6, 22)
	arc_tordu_durabilité_Max := functionshelper.RandomBetween(20, 40)
	arc_tordu_dégâts := functionshelper.RandomBetween(2, 7)
	items.Arc_tordu = items.InitArmeMonster("Arc tordu", "any", 1, arc_tordu_durabilité_Actuelle, arc_tordu_durabilité_Max, "Arc mal façonné, flèches de fortune", arc_tordu_dégâts, 2)

	// cuir bouilli rapiécé - Armure de cuir mal entretenue
	cuir_bouilli_durabilité_Actuelle := functionshelper.RandomBetween(15, 35)
	cuir_bouilli_durabilité_Max := functionshelper.RandomBetween(40, 60)
	items.Cuir_bouilli_rapiécé = items.InitArmure("Cuir bouilli rapiécé", "any", 1, cuir_bouilli_durabilité_Max, cuir_bouilli_durabilité_Actuelle, "Armure de cuir mal entretenue", 6)

	// casque bosselé - Heaume déformé et rouillé
	casque_bosselé_durabilité_Actuelle := functionshelper.RandomBetween(10, 25)
	casque_bosselé_durabilité_Max := functionshelper.RandomBetween(30, 45)
	items.Casque_bosselé = items.InitArmure("Casque bosselé", "any", 1, casque_bosselé_durabilité_Max, casque_bosselé_durabilité_Actuelle, "Heaume déformé et rouillé", 4)

	// bourse gold
	bourse_gold := functionshelper.RandomBetween(1, 8)
	items.Bourse_de_cuir = items.InitBourse(bourse_gold)
	if startScreen { // affichage de l'écran de démarrage (une seule fois)
		startscreen.StartScreen()
		menus.NewMap()
		menus.MenuMap()
		startScreen = false
	}
}
