package combat

import (
	"fmt"
	"strings"
	"time"

	"ASCII_Aventure/characters"
	"ASCII_Aventure/couleurs"
	functionshelper "ASCII_Aventure/functions_helper"
	"ASCII_Aventure/inputcontrol"
	"ASCII_Aventure/items"
	"ASCII_Aventure/skills"
	"ASCII_Aventure/startscreen"
)

var currentEnemy *characters.Monster
var combatEnCours bool = false
var selectedCharacter *characters.Character

func Combat(Tour int) {
	startscreen.ClearScreen()
	if combatEnCours {
		return
	}
	combatEnCours = true
	var character *characters.Character
	if Tour == 0 {
		if characters.C2_b && characters.C2 != nil {
			selectedCharacter = selectionnerPersonnage()
		} else {
			selectedCharacter = characters.C1
		}
	}
	character = selectedCharacter
	if character == nil {
		fmt.Printf("%sErreur: Aucun personnage sélectionné%s\n", couleurs.Red, couleurs.Reset)
		return
	}
	if currentEnemy == nil {
		currentEnemy = getCurrentEnemy()
	}
	startscreen.ClearScreen()
	afficherInterfaceCombat(character, currentEnemy, Tour)
	if !characters.IsDead() && !isDeadMonster() {
		// tour du joueur
		joueurAJoué := characterTurn(character, currentEnemy)
		// vérifier si l'ennemi est mort après l'attaque du joueur
		if !isDeadMonster() && joueurAJoué {
			// tour du monstre seulement si le joueur a effectué une action
			goblinPattern(currentEnemy, character, Tour)
		}
		fmt.Printf("\n%sAppuyez sur Entrée pour continuer...%s", couleurs.Blue, couleurs.Reset)
		functionshelper.ReadInput()
		combatEnCours = false
		Combat(Tour + 1)
	} else {
		terminerCombat()
	}
	combatEnCours = false
}

func goblinPattern(attaquant *characters.Monster, cible *characters.Character, tour int) {
	if attaquant == nil || cible == nil {
		return
	}
	// on calcule le pourcentage de vie du gobelin
	pourcentageVieGobelin := float64(attaquant.PointsDeVieActuels) / float64(attaquant.PointsDeVieMaximum)
	// on détermine l'action à effectuer
	var actionChoisie string
	// si le gobelin a moins de 30% de PV, il privilégie le soin
	if pourcentageVieGobelin < 0.3 && functionshelper.RandomBetween(1, 100) <= 70 {
		actionChoisie = "Soin"
	} else if tour%4 == 0 && functionshelper.RandomBetween(1, 100) <= 80 {
		// tous les 4 tours, haute chance d'utiliser "Coup de Rage"
		actionChoisie = "Coup de Rage"
	} else if functionshelper.RandomBetween(1, 100) <= 40 {
		// 40% de chance d'utiliser "Griffure"
		actionChoisie = "Griffure"
	} else {
		// sinon attaque normale
		actionChoisie = "Attaque normale"
	}
	fmt.Printf("\n%s=== Tour du %s ====%s\n", couleurs.Red+couleurs.Bold, attaquant.Nom, couleurs.Reset)
	switch actionChoisie {
	case "Soin":
		utiliserSkillGobelin(attaquant, cible, "Soin")
	case "Coup de Rage":
		utiliserSkillGobelin(attaquant, cible, "Coup de Rage")
	case "Griffure":
		utiliserSkillGobelin(attaquant, cible, "Griffure")
	default:
		// une attaque normale avec possibilité critique
		attaqueNormaleGobelin(attaquant, cible, tour)
	}
	time.Sleep(1 * time.Second)
}

func attaqueNormaleGobelin(attaquant *characters.Monster, cible *characters.Character, tour int) {
	degatsBase := calculateMonsterDamage(attaquant)
	var degatsInfliges int
	var typeAttaque string
	// attaque critique tous les 3 tours
	if tour%3 == 0 {
		degatsInfliges = int(float64(degatsBase) * 2.0)
		typeAttaque = " avec une attaque critique"
		fmt.Printf("\n%sATTAQUE CRITIQUE !%s\n", couleurs.Red+couleurs.Bold, couleurs.Reset)
	} else {
		degatsInfliges = degatsBase
		typeAttaque = ""
	}
	if attaquant.Equipement.Arme != "" {
		applyMonsterWeaponWear(attaquant)
	}
	fmt.Printf("\n%s%s%s inflige à %s%s%s %s%d%s de dégâts%s\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.Green, cible.Nom, couleurs.Reset, couleurs.Yellow, degatsInfliges, couleurs.Reset, typeAttaque)
	cible.PointsDeVieActuels -= degatsInfliges
	if cible.PointsDeVieActuels < 0 {
		cible.PointsDeVieActuels = 0
	}
	applyArmorWear(cible)
	fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, cible.Nom, couleurs.Reset, couleurs.White, cible.PointsDeVieActuels, couleurs.Reset, couleurs.White, cible.PointsDeVieMaximum, couleurs.Reset)
}

func utiliserSkillGobelin(attaquant *characters.Monster, cible *characters.Character, skill string) {
	fmt.Printf("\n%s%s%s utilise %s%s%s!\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.Purple+couleurs.Bold, skill, couleurs.Reset)
	switch skill {
	case "Soin":
		// le gobelin se soigne
		soinQuantite := functionshelper.RandomBetween(15, 25)
		attaquant.PointsDeVieActuels += soinQuantite
		if attaquant.PointsDeVieActuels > attaquant.PointsDeVieMaximum {
			attaquant.PointsDeVieActuels = attaquant.PointsDeVieMaximum
		}
		fmt.Printf("%s%s%s se soigne de %s%d%s PV!\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.Green, soinQuantite, couleurs.Reset)
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.White, attaquant.PointsDeVieActuels, couleurs.Reset, couleurs.White, attaquant.PointsDeVieMaximum, couleurs.Reset)
	case "Griffure":
		// attaque rapide avec saignement
		degats := calculateMonsterDamage(attaquant) + functionshelper.RandomBetween(3, 8)
		fmt.Printf("%s%s%s griffe sauvagement %s%s%s et inflige %s%d%s dégâts!\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.Green, cible.Nom, couleurs.Reset, couleurs.Yellow, degats, couleurs.Reset)
		cible.PointsDeVieActuels -= degats
		if cible.PointsDeVieActuels < 0 {
			cible.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, cible.Nom, couleurs.Reset, couleurs.White, cible.PointsDeVieActuels, couleurs.Reset, couleurs.White, cible.PointsDeVieMaximum, couleurs.Reset)
	case "Coup de Rage":
		// attaque puissante avec bonus de dégâts
		degatsBase := calculateMonsterDamage(attaquant)
		degats := int(float64(degatsBase) * 1.8) // 80% de bonus de dégats
		fmt.Printf("%s%s%s entre dans une rage folle et frappe violemment!\n", couleurs.Red+couleurs.Bold, attaquant.Nom, couleurs.Reset)
		fmt.Printf("%s%s%s inflige %s%d%s dégâts de rage à %s%s%s!\n", couleurs.Red, attaquant.Nom, couleurs.Reset, couleurs.Yellow+couleurs.Bold, degats, couleurs.Reset, couleurs.Green, cible.Nom, couleurs.Reset)
		cible.PointsDeVieActuels -= degats
		if cible.PointsDeVieActuels < 0 {
			cible.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, cible.Nom, couleurs.Reset, couleurs.White, cible.PointsDeVieActuels, couleurs.Reset, couleurs.White, cible.PointsDeVieMaximum, couleurs.Reset)
	}
	applyArmorWear(cible)
	if attaquant.Equipement.Arme != "" {
		applyMonsterWeaponWear(attaquant)
	}
}

func applyMonsterWeaponWear(monster *characters.Monster) {
	if monster.Equipement.Arme == "" {
		return
	}
	var armeItem *items.Item
	switch monster.Equipement.Arme {
	case "Dague rouillée":
		armeItem = items.Dague_rouillée
	case "Gourdin clouté":
		armeItem = items.Gourdin_clouté
	case "Arc tordu":
		armeItem = items.Arc_tordu
	case "Épée en fer":
		armeItem = items.Épée_en_fer
	}
	if armeItem != nil && armeItem.DurabilitéesActuelle > 0 {
		usure := functionshelper.RandomBetween(1, 2)
		armeItem.DurabilitéesActuelle -= usure

		if armeItem.DurabilitéesActuelle <= 0 {
			armeItem.DurabilitéesActuelle = 0
			fmt.Printf("%s%s%s a cassé son %s%s%s !%s\n",
				couleurs.Red, monster.Nom, couleurs.Reset,
				couleurs.Yellow, armeItem.Nom, couleurs.Reset, couleurs.Reset)
			monster.Equipement.Arme = ""
		}
	}
}

func calculateMonsterDamage(monster *characters.Monster) int {
	baseDamage := 5 // ici les dégâts de base du gobelin
	// on ajouter les dégâts de l'arme équipée
	if monster.Equipement.Arme != "" {
		var armeItem *items.Item
		switch monster.Equipement.Arme {
		case "Dague rouillée":
			armeItem = items.Dague_rouillée
		case "Gourdin clouté":
			armeItem = items.Gourdin_clouté
		case "Arc tordu":
			armeItem = items.Arc_tordu
		case "Épée en fer":
			armeItem = items.Épée_en_fer
		}
		if armeItem != nil {
			baseDamage += armeItem.Degats
		}
	}
	return baseDamage
}

func calculatePlayerDamage(character *characters.Character) int {
	baseDamage := 10
	if character.Equipment.Arme != "" && character.Equipment.Arme != "Poing" {
		var armeItem *items.Item
		switch character.Equipment.Arme {
		case "Dague rouillée":
			armeItem = items.Dague_rouillée
		case "Gourdin clouté":
			armeItem = items.Gourdin_clouté
		case "Arc tordu":
			armeItem = items.Arc_tordu
		case "Épée en fer", "épée en fer":
			armeItem = items.Épée_en_fer
		}
		if armeItem != nil {
			// on vérifie si l'arme est encore utilisable
			if armeItem.DurabilitéesActuelle > 0 {
				// ajouter les dégâts de l'arme
				baseDamage += armeItem.Degats
				// appliquer l'usure à l'arme
				applyWeaponWear(character, armeItem)
			} else {
				// l'arme est cassée
				fmt.Printf("%sVotre %s est cassée et ne peut plus être utilisée !%s\n",
					couleurs.Red, armeItem.Nom, couleurs.Reset)
				fmt.Printf("%sVous attaquez avec vos poings à la place.%s\n",
					couleurs.Yellow, couleurs.Reset)
				// on déséquipe automatiquement l'arme cassée
				character.Equipment.Arme = ""
			}
		}
	}
	return baseDamage
}

func selectionnerPersonnage() *characters.Character {
	fmt.Printf("\n%sAvec quel personnage souhaitez-vous attaquer ?%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s1. %s%s%s\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset)
	fmt.Printf("%s2. %s%s%s\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset)
	fmt.Printf("%sVotre choix : %s", couleurs.Blue, couleurs.Reset)
	choixPerso := functionshelper.ReadInput()
	switch strings.ToLower(strings.TrimSpace(choixPerso)) {
	case "1", "1.", strings.ToLower(characters.C1.Nom):
		return characters.C1
	case "2", "2.", strings.ToLower(characters.C2.Nom):
		return characters.C2
	default:
		fmt.Printf("%sChoix invalide%s\n", couleurs.Red, couleurs.Reset)
		return nil
	}
}

func afficherInterfaceCombat(character *characters.Character, enemy *characters.Monster, tour int) {
	if character == nil || enemy == nil {
		return
	}
	fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                    %sCOMBAT - Tour %-2d%s ── %s%s%s vs %s%s%s                            %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Red, tour, couleurs.Reset, couleurs.Green, character.Nom, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────┬─┬─────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s%-39s%s %s│%s %s│%s %s%-39s%s %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Green, character.Nom, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────┼─┼─────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                         %s│%s %s│%s                                         %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	pvJoueur := fmt.Sprintf("PV : %d/%d", character.PointsDeVieActuels, character.PointsDeVieMaximum)
	pvEnnemi := fmt.Sprintf("PV : %d/%d", enemy.PointsDeVieActuels, enemy.PointsDeVieMaximum)
	fmt.Printf("%s│%s %s%-39s%s %s│%s %s│%s %s%-39s%s %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, pvJoueur, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.White, pvEnnemi, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                         %s│%s %s│%s                                         %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %sDurabilité équipement :%s                 %s│%s %s│%s %sDurabilité équipement :%s                 %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Yellow, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Yellow, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                         %s│%s %s│%s                                         %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	equipementsJoueur := getCharacterEquipementsDurability(character)
	equipementsEnnemi := getMonsterEquipementsDurability(enemy)
	maxLines := 4
	if len(equipementsJoueur) > maxLines {
		maxLines = len(equipementsJoueur)
	}
	if len(equipementsEnnemi) > maxLines {
		maxLines = len(equipementsEnnemi)
	}
	for i := 0; i < maxLines; i++ {
		var equipJoueur string
		var equipEnnemi string
		if i < len(equipementsJoueur) {
			equipJoueur = equipementsJoueur[i]
		} else {
			equipJoueur = "-"
		}
		if i < len(equipementsEnnemi) {
			equipEnnemi = equipementsEnnemi[i]
		} else {
			equipEnnemi = "-"
		}
		ligneJoueur := fmt.Sprintf("- %s", equipJoueur)
		ligneEnnemi := fmt.Sprintf("- %s", equipEnnemi)
		fmt.Printf("%s│%s %s%-39s%s %s│%s %s│%s %s%-39s%s %s│%s\n",
			couleurs.Blue, couleurs.Reset,
			couleurs.White, ligneJoueur, couleurs.Reset,
			couleurs.Blue, couleurs.Reset,
			couleurs.Blue, couleurs.Reset,
			couleurs.White, ligneEnnemi, couleurs.Reset,
			couleurs.Blue, couleurs.Reset)
		fmt.Printf("%s│%s                                         %s│%s %s│%s                                         %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	fmt.Printf("%s└─────────────────────────────────────────┴─┴─────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
}

func getCurrentEnemy() *characters.Monster {
	if characters.Gobelin == nil { // si pas d'ennemi, créer un gobelin par défaut
		pv := functionshelper.RandomBetween(38, 40) // points de vie aléatoires entre 38 et 40
		equipements := []string{
			"Dague rouillée", "Gourdin clouté", "Arc tordu",
		}
		armures := []string{
			"Cuir bouilli rapiécé", "Casque bosselé",
		}
		armeIndex := functionshelper.RandomBetween(0, len(equipements)-1) // sélection aléatoire d'une arme
		armeChoisie := equipements[armeIndex]
		var armureChoisie string // sélection aléatoire d'une armure (50% de chance)
		if functionshelper.RandomBetween(0, 1) == 1 {
			armureIndex := functionshelper.RandomBetween(0, len(armures)-1)
			armureChoisie = armures[armureIndex]
		}
		inventaire := []string{"Bourse de cuir"} // inventaire avec la bourse de cuir
		equipement := characters.Equipement{}    // attribution de l'équipement aléatoire
		if armeChoisie == "Dague rouillée" || armeChoisie == "Gourdin clouté" || armeChoisie == "Arc tordu" {
			equipement.Arme = armeChoisie
		}
		switch armureChoisie {
		case "Cuir bouilli rapiécé":
			equipement.Tunique = armureChoisie
		case "Casque bosselé":
			equipement.Chapeau = armureChoisie
		}
		characters.Gobelin = &characters.Monster{
			Nom:                "Gobelin",
			Classe:             "Gobelin",
			PointsDeVieMaximum: pv,
			PointsDeVieActuels: pv,
			Inventaire:         inventaire,
			Skill:              []string{"Griffure", "Coup de Rage", "Soin"},
			Equipement:         equipement,
		}
	}
	return characters.Gobelin
}

func getCharacterEquipementsDurability(character *characters.Character) []string {
	var equipements []string
	// vérifier l'arme
	if character.Equipment.Arme != "" && character.Equipment.Arme != "Poing" { // implémenter les poings
		// récupérer l'item correspondant à l'arme équipée
		var armeItem *items.Item
		switch character.Equipment.Arme {
		case "Dague rouillée":
			armeItem = items.Dague_rouillée
		case "Gourdin clouté":
			armeItem = items.Gourdin_clouté
		case "Arc tordu":
			armeItem = items.Arc_tordu
		case "Épée en fer":
			armeItem = items.Épée_en_fer
		}
		if armeItem != nil {
			equipements = append(equipements, fmt.Sprintf("Arme: %d/%d", armeItem.DurabilitéesActuelle, armeItem.DurabilitéMaximum))
		}
	}
	if character.Equipment.Tunique != "" { // vérifier la tunique
		var tuniqueItem *items.Item
		switch character.Equipment.Tunique {
		case "Tunique de l'aventurier":
			tuniqueItem = items.Tunique_de_laventurier
		case "Cuir bouilli rapiécé":
			tuniqueItem = items.Cuir_bouilli_rapiécé
		}
		if tuniqueItem != nil {
			equipements = append(equipements, fmt.Sprintf("Tunique: %d/%d", tuniqueItem.DurabilitéesActuelle, tuniqueItem.DurabilitéMaximum))
		}
	}
	if character.Equipment.Chapeau != "" { // vérifier le chapeau
		var chapeauItem *items.Item
		switch character.Equipment.Chapeau {
		case "Chapeau de l'aventurier":
			chapeauItem = items.Chapeau_de_laventurier
		case "Casque bosselé":
			chapeauItem = items.Casque_bosselé
		}
		if chapeauItem != nil {
			equipements = append(equipements, fmt.Sprintf("Chapeau: %d/%d",
				chapeauItem.DurabilitéesActuelle, chapeauItem.DurabilitéMaximum))
		}
	}
	if character.Equipment.Bottes != "" { // vérifier les bottes
		var bottesItem *items.Item
		switch character.Equipment.Bottes {
		case "Bottes de l'aventurier":
			bottesItem = items.Bottes_de_laventurier
		}
		if bottesItem != nil {
			equipements = append(equipements, fmt.Sprintf("Bottes: %d/%d",
				bottesItem.DurabilitéesActuelle, bottesItem.DurabilitéMaximum))
		}
	}
	if len(equipements) == 0 {
		equipements = append(equipements, "Aucun équipement")
	}
	return equipements
}

// globalement c'est comme au dessus mais pour les monstres
func getMonsterEquipementsDurability(monster *characters.Monster) []string {
	var equipements []string
	if monster.Equipement.Arme != "" {
		var armeItem *items.Item
		switch monster.Equipement.Arme {
		case "Dague rouillée":
			armeItem = items.Dague_rouillée
		case "Gourdin clouté":
			armeItem = items.Gourdin_clouté
		case "Arc tordu":
			armeItem = items.Arc_tordu
		case "Épée en fer":
			armeItem = items.Épée_en_fer
		}
		if armeItem != nil {
			equipements = append(equipements, fmt.Sprintf("Arme: %d/%d", armeItem.DurabilitéesActuelle, armeItem.DurabilitéMaximum))
		}
	}
	if monster.Equipement.Tunique != "" {
		var tuniqueItem *items.Item
		switch monster.Equipement.Tunique {
		case "Tunique de l'aventurier":
			tuniqueItem = items.Tunique_de_laventurier
		case "Cuir bouilli rapiécé":
			tuniqueItem = items.Cuir_bouilli_rapiécé
		}
		if tuniqueItem != nil {
			equipements = append(equipements, fmt.Sprintf("Tunique: %d/%d", tuniqueItem.DurabilitéesActuelle, tuniqueItem.DurabilitéMaximum))
		}
	}
	if monster.Equipement.Chapeau != "" {
		var chapeauItem *items.Item
		switch monster.Equipement.Chapeau {
		case "Chapeau de l'aventurier":
			chapeauItem = items.Chapeau_de_laventurier
		case "Casque bosselé":
			chapeauItem = items.Casque_bosselé
		}
		if chapeauItem != nil {
			equipements = append(equipements, fmt.Sprintf("Chapeau: %d/%d",
				chapeauItem.DurabilitéesActuelle, chapeauItem.DurabilitéMaximum))
		}
	}
	if monster.Equipement.Bottes != "" {
		var bottesItem *items.Item
		switch monster.Equipement.Bottes {
		case "Bottes de l'aventurier":
			bottesItem = items.Bottes_de_laventurier
		}
		if bottesItem != nil {
			equipements = append(equipements, fmt.Sprintf("Bottes: %d/%d",
				bottesItem.DurabilitéesActuelle, bottesItem.DurabilitéMaximum))
		}
	}
	if len(equipements) == 0 {
		equipements = append(equipements, "Aucun équipement")
	}
	return equipements
}

func isDeadMonster() bool {
	if currentEnemy != nil && currentEnemy.PointsDeVieActuels <= 0 {
		fmt.Printf("\n\n%sVictoire ! %s%s%s est vaincu !%s\n\n", couleurs.Green+couleurs.Bold, couleurs.Red, currentEnemy.Nom, couleurs.Green+couleurs.Bold, couleurs.Reset)
		return true
	}
	return false
}

func RechercheEnemy() {
	startscreen.ClearScreen()
	const (
		texteChargement = "Recherche d'un ennemi en cours"
		largeurEcran    = 120
		etapesTotal     = 20
		delaiMs         = 150
	)
	espacesAGauche := max((largeurEcran-len(texteChargement)-10)/2, 0)
	fmt.Print("\n\n\033[?25l")   // cache le curseur
	defer fmt.Print("\033[?25h") // réaffiche le curseur à la fin
	for i := 0; i < etapesTotal; i++ {
		afficherEtapeChargement(i, espacesAGauche, texteChargement)
		time.Sleep(delaiMs * time.Millisecond)
	}
	fmt.Print("\n\n")
	fmt.Printf("%sEnnemi trouvé !%s\n", couleurs.Green+couleurs.Bold, couleurs.Reset)
	time.Sleep(2 * time.Second)
}

func afficherEtapeChargement(etape int, espaces int, texte string) {
	fmt.Print("\033[3;1H")
	fmt.Print("\033[2K")
	nombrePoints := etape % 4
	fmt.Printf("%*s%s%s%s%s\n", espaces, "", couleurs.Blue+couleurs.Bold, texte, strings.Repeat(".", nombrePoints), couleurs.Reset)
	pourcentage := (etape + 1) * 5
	barresRemplies := pourcentage / 5
	barresVides := 20 - barresRemplies
	fmt.Printf("%*s%s[%s%s%s%s%s] %d%%%s", espaces, "", couleurs.Blue, couleurs.Yellow, strings.Repeat("█", barresRemplies), couleurs.White, strings.Repeat("░", barresVides), couleurs.Blue, pourcentage, couleurs.Reset)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterTurn(character *characters.Character, enemy *characters.Monster) bool {
	if character == nil || enemy == nil {
		return false
	}
	fmt.Printf("\n%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                 TOUR DU JOUEUR                                      %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %sQue souhaitez-vous faire ?%s                                                          %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s1.%s Attaquer                                                                         %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s2.%s Inventaire                                                                       %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s3.%s Skills                                                                           %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("\n%sVotre choix : %s", couleurs.Blue, couleurs.Reset)
	inputcontrol.ClearInputBuffer()
	choix := strings.ToLower(strings.TrimSpace(functionshelper.ReadInput()))
	switch choix {
	case "1", "attaquer", "attaque":
		// attaque basique avec arme équipée
		degats := calculatePlayerDamage(character)
		fmt.Printf("\n%s%s%s utilise %sAttaque basique%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, couleurs.Reset)
		fmt.Printf("%s%s%s inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		// appliquer les dégâts
		enemy.PointsDeVieActuels -= degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		// afficher les PV restants de l'ennemi
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
		time.Sleep(1 * time.Second)
		return true
	case "2", "inventaire", "inv":
		return handleInventoryAction(character, enemy)
	case "3", "skills", "skill":
		return handleSkillsAction(character, enemy)
	default:
		fmt.Printf("%sChoix invalide. Veuillez choisir 1, 2 ou 3.%s\n", couleurs.Red, couleurs.Reset)
		return characterTurn(character, enemy) // redemander le choix
	}
}

func handleSkillsAction(character *characters.Character, enemy *characters.Monster) bool {
	if len(character.Skill) == 0 {
		fmt.Printf("\n%sVous ne connaissez aucun skill.%s\n", couleurs.Red, couleurs.Reset)
		fmt.Printf("%sAppuyez sur Entrée pour continuer...%s", couleurs.Blue, couleurs.Reset)
		functionshelper.ReadInput()
		return false // retourner au menu principal
	}
	fmt.Printf("\n%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                 SKILLS                                              %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	for i, skill := range character.Skill {
		fmt.Printf("%s│%s %s%d.%s %s%-79s%s %s │%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, i+1, couleurs.Reset, couleurs.White, skill, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s0.%s Retour                                                                           %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("\n%sChoisissez un skill à utiliser (numéro) : %s", couleurs.Blue, couleurs.Reset)
	choix := strings.TrimSpace(functionshelper.ReadInput())
	if choix == "0" {
		return false // retourner au menu
	}
	// convertir le choix en index
	var index int
	switch choix {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	case "5":
		index = 4
	case "6":
		index = 5
	case "7":
		index = 6
	case "8":
		index = 7
	case "9":
		index = 8
	case "10":
		index = 9
	default:
		fmt.Printf("%sChoix invalide.%s\n", couleurs.Red, couleurs.Reset)
		return handleSkillsAction(character, enemy) // redemander
	}
	if index < 0 || index >= len(character.Skill) {
		fmt.Printf("%sChoix invalide.%s\n", couleurs.Red, couleurs.Reset)
		return handleSkillsAction(character, enemy)
	}
	// utiliser le skill
	skill := character.Skill[index]
	fmt.Printf("\n%s%s%s utilise %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Purple, skill, couleurs.Reset)
	// appliquer l'effet selon le type de skill
	useSkill(character, enemy, skill)
	time.Sleep(1 * time.Second)
	return true
}

func useSkill(character *characters.Character, enemy *characters.Monster, skillName string) {
	var skill *skills.Skill
	// récupérer le skill correspondant
	switch skillName {
	case "Boule de Feu", "[Spell book] > Boule de feu":
		skill = skills.BouleDeFeu
	case "Coup de Rage":
		skill = skills.CoupDeRage
	case "Soin":
		skill = skills.Soin
	case "Coup de poing":
		skill = skills.CoupDePoing
	case "Coup de hache":
		skill = skills.CoupDeHache
	case "Tir à l'arc":
		skill = skills.TirÀLarc
	case "Griffure", "[Spell book] > Griffure":
		skill = skills.Griffure
	}
	if skill == nil {
		fmt.Printf("%sCe skill n'est pas encore implémenté.%s\n", couleurs.Red, couleurs.Reset)
		return
	}
	switch skillName {
	case "Boule de Feu", "[Spell book] > Boule de feu":
		fmt.Printf("%s%s%s lance une boule de feu et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	case "Coup de Rage":
		fmt.Printf("%s%s%s entre dans une rage folle et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	case "Soin":
		healAmount := 30
		character.PointsDeVieActuels += healAmount
		if character.PointsDeVieActuels > character.PointsDeVieMaximum {
			character.PointsDeVieActuels = character.PointsDeVieMaximum
		}
		fmt.Printf("%s%s%s utilise un sort de soin et récupère %s%d%s points de vie\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, healAmount, couleurs.Reset)
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.White, character.PointsDeVieActuels, couleurs.Reset, couleurs.White, character.PointsDeVieMaximum, couleurs.Reset)
	case "Coup de poing":
		fmt.Printf("%s%s%s donne un coup de poing et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	case "Coup de hache":
		fmt.Printf("%s%s%s frappe avec sa hache et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	case "Tir à l'arc":
		fmt.Printf("%s%s%s tire une flèche et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Red, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Red, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	case "Griffure":
		fmt.Printf("%s%s%s griffe sauvagement et inflige %s%d%s dégâts à %s%s%s\n", couleurs.Red, character.Nom, couleurs.Reset, couleurs.Yellow, skill.Degats, couleurs.Reset, couleurs.Green, enemy.Nom, couleurs.Reset)
		enemy.PointsDeVieActuels -= skill.Degats
		if enemy.PointsDeVieActuels < 0 {
			enemy.PointsDeVieActuels = 0
		}
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, enemy.Nom, couleurs.Reset, couleurs.White, enemy.PointsDeVieActuels, couleurs.Reset, couleurs.White, enemy.PointsDeVieMaximum, couleurs.Reset)
	default:
		fmt.Printf("%sCe skill n'est pas encore implémenté.%s\n", couleurs.Red, couleurs.Reset)
	}
}

func handleInventoryAction(character *characters.Character, monster *characters.Monster) bool {
	if len(character.Inventaire) == 0 {
		fmt.Printf("\n%sVotre inventaire est vide.%s\n", couleurs.Red, couleurs.Reset)
		fmt.Printf("%sAppuyez sur Entrée pour continuer...%s", couleurs.Blue, couleurs.Reset)
		functionshelper.ReadInput()
		return false // retourner au menu principal
	}
	fmt.Printf("\n%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                 INVENTAIRE                                          %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	for i, item := range character.Inventaire {
		fmt.Printf("%s│%s %s%d.%s %s%-79s%s %s │%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, i+1, couleurs.Reset, couleurs.White, item, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %s0.%s Retour                                                                           %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("\n%sChoisissez un objet à utiliser (numéro) : %s", couleurs.Blue, couleurs.Reset)
	choix := strings.TrimSpace(functionshelper.ReadInput())
	if choix == "0" {
		return false // retourner au menu
	}
	// convertir le choix en index
	var index int
	switch choix {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	case "5":
		index = 4
	case "6":
		index = 5
	case "7":
		index = 6
	case "8":
		index = 7
	case "9":
		index = 8
	case "10":
		index = 9
	default:
		fmt.Printf("%sChoix invalide.%s\n", couleurs.Red, couleurs.Reset)
		return handleInventoryAction(character, monster) // redemander
	}
	if index < 0 || index >= len(character.Inventaire) {
		fmt.Printf("%sChoix invalide.%s\n", couleurs.Red, couleurs.Reset)
		return handleInventoryAction(character, monster)
	}
	// utiliser l'objet
	item := character.Inventaire[index]
	fmt.Printf("\n%sVous utilisez %s%s%s\n", couleurs.Green, couleurs.Yellow, item, couleurs.Reset)
	// appliquer l'effet selon le type d'objet
	useItem(character, monster, item, index)
	time.Sleep(1 * time.Second)
	return true
}

func useItem(character *characters.Character, monster *characters.Monster, itemName string, index int) {
	switch itemName {
	case "Potion de soin", "potion de soin":
		healAmount := 50
		character.PointsDeVieActuels += healAmount
		if character.PointsDeVieActuels > character.PointsDeVieMaximum {
			character.PointsDeVieActuels = character.PointsDeVieMaximum
		}
		fmt.Printf("%s%s%s récupère %s%d%s points de vie\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.Yellow, healAmount, couleurs.Reset)
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, character.Nom, couleurs.Reset, couleurs.White, character.PointsDeVieActuels, couleurs.Reset, couleurs.White, character.PointsDeVieMaximum, couleurs.Reset)
		character.Inventaire = functionshelper.RemoveItem(character.Inventaire, index)
	case "Potion de poison", "potion de poison":
		healAmount := 20
		monster.PointsDeVieActuels -= healAmount / 3
		time.Sleep(1 * time.Second)
		monster.PointsDeVieActuels -= healAmount / 3
		time.Sleep(1 * time.Second)
		monster.PointsDeVieActuels -= healAmount / 3
		if monster.PointsDeVieActuels > monster.PointsDeVieMaximum {
			monster.PointsDeVieActuels = monster.PointsDeVieMaximum
		}
		fmt.Printf("%s%s%s perd %s%d%s points de vie\n", couleurs.Green, monster.Nom, couleurs.Reset, couleurs.Yellow, healAmount, couleurs.Reset)
		fmt.Printf("%s%s%s: %s%d%s/%s%d%s PV\n", couleurs.Green, monster.Nom, couleurs.Reset, couleurs.White, monster.PointsDeVieActuels, couleurs.Reset, couleurs.White, monster.PointsDeVieMaximum, couleurs.Reset)
		character.Inventaire = functionshelper.RemoveItem(character.Inventaire, index)
	default:
		fmt.Printf("%sCet objet ne peut pas être utilisé en combat.%s\n", couleurs.Red, couleurs.Reset)
	}
}

func terminerCombat() {
	if isDeadMonster() {
		fmt.Printf("\n%sCombat terminé !%s\n", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		var character *characters.Character
		if characters.C2_b && characters.C2 != nil {
			character = selectionnerPersonnageButin()
		} else {
			character = characters.C1
		}
		if character != nil && currentEnemy != nil {
			ajouterButin(character, currentEnemy)
		}
		currentEnemy = nil
		selectedCharacter = nil
		characters.Gobelin = nil
	} else if characters.IsDead() {
		fmt.Printf("\n%sDéfaite !%s\n", couleurs.Red+couleurs.Bold, couleurs.Reset)
		currentEnemy = nil
		characters.Gobelin = nil
	}
}

func applyWeaponWear(character *characters.Character, arme *items.Item) {
	if arme == nil || arme.DurabilitéesActuelle <= 0 {
		return
	}
	// on calcule l'usure (1-3 points de durabilité perdus)
	usure := functionshelper.RandomBetween(1, 3)
	arme.DurabilitéesActuelle -= usure
	// on s'assure que la durabilité ne descend pas en dessous de 0
	if arme.DurabilitéesActuelle < 0 {
		arme.DurabilitéesActuelle = 0
	}
	// afficher l'état de l'arme
	if arme.DurabilitéesActuelle == 0 {
		fmt.Printf("%sVotre %s%s%s s'est cassée ! (0/%d)%s\n", couleurs.Red, couleurs.Yellow, arme.Nom, couleurs.Red, arme.DurabilitéMaximum, couleurs.Reset)
		character.Equipment.Arme = ""
		fmt.Printf("%s%s%s a été automatiquement déséquipée.%s\n", couleurs.Yellow, arme.Nom, couleurs.Reset, couleurs.Reset)
	} else if arme.DurabilitéesActuelle <= 10 {
		fmt.Printf("%sVotre %s%s%s est en mauvais état ! (%d/%d)%s\n", couleurs.Yellow, couleurs.White, arme.Nom, couleurs.Yellow, arme.DurabilitéesActuelle, arme.DurabilitéMaximum, couleurs.Reset)
	} else {
		fmt.Printf("%sDurabilité de %s%s%s : %s%d%s/%s%d%s\n", couleurs.White, couleurs.Green, arme.Nom, couleurs.White, couleurs.Green, arme.DurabilitéesActuelle, couleurs.White, couleurs.Green, arme.DurabilitéMaximum, couleurs.Reset)
	}
}

func applyArmorWear(character *characters.Character) {
	// usure du chapeau
	if character.Equipment.Chapeau != "" {
		chapeau := functionshelper.GetItemByName(character.Equipment.Chapeau)
		if chapeau != nil && chapeau.DurabilitéesActuelle > 0 {
			usure := functionshelper.RandomBetween(1, 2)
			chapeau.DurabilitéesActuelle -= usure
			if chapeau.DurabilitéesActuelle < 0 {
				chapeau.DurabilitéesActuelle = 0
			}
			if chapeau.DurabilitéesActuelle == 0 {
				fmt.Printf("%sVotre %s s'est détérioré et ne protège plus !%s\n",
					couleurs.Red, chapeau.Nom, couleurs.Reset)
				character.Equipment.Chapeau = ""
			}
		}
	}
	// usure de la tunique
	if character.Equipment.Tunique != "" {
		tunique := functionshelper.GetItemByName(character.Equipment.Tunique)
		if tunique != nil && tunique.DurabilitéesActuelle > 0 {
			usure := functionshelper.RandomBetween(1, 2)
			tunique.DurabilitéesActuelle -= usure
			if tunique.DurabilitéesActuelle < 0 {
				tunique.DurabilitéesActuelle = 0
			}
			if tunique.DurabilitéesActuelle == 0 {
				fmt.Printf("%sVotre %s s'est détériorée !%s\n",
					couleurs.Red, tunique.Nom, couleurs.Reset)
				character.Equipment.Tunique = ""
			}
		}
	}
	// usure des bottes
	if character.Equipment.Bottes != "" {
		bottes := functionshelper.GetItemByName(character.Equipment.Bottes)
		if bottes != nil && bottes.DurabilitéesActuelle > 0 {
			usure := functionshelper.RandomBetween(1, 2)
			bottes.DurabilitéesActuelle -= usure
			if bottes.DurabilitéesActuelle < 0 {
				bottes.DurabilitéesActuelle = 0
			}
			if bottes.DurabilitéesActuelle == 0 {
				fmt.Printf("%sVos %s sont complètement usées !%s\n",
					couleurs.Red, bottes.Nom, couleurs.Reset)
				character.Equipment.Bottes = ""
			}
		}
	}
}

func ajouterButin(character *characters.Character, enemy *characters.Monster) {
	fmt.Printf("\n%sRécupération du butin...%s\n", couleurs.Yellow+couleurs.Bold, couleurs.Reset)
	var butinRecupere []string
	var butinNonRecupere []string
	// récupérer les objets de l'inventaire de l'ennemi
	for _, item := range enemy.Inventaire {
		if len(character.Inventaire) < 10 { // Limite d'inventaire
			character.Inventaire = append(character.Inventaire, item)
			butinRecupere = append(butinRecupere, item)
		} else {
			butinNonRecupere = append(butinNonRecupere, item)
		}
	}
	// récupérer l'équipement de l'ennemi s'il en a
	if enemy.Equipement.Arme != "" {
		if len(character.Inventaire) < 10 {
			character.Inventaire = append(character.Inventaire, enemy.Equipement.Arme)
			butinRecupere = append(butinRecupere, enemy.Equipement.Arme)
		} else {
			butinNonRecupere = append(butinNonRecupere, enemy.Equipement.Arme)
		}
	}
	if enemy.Equipement.Tunique != "" {
		if len(character.Inventaire) < 10 {
			character.Inventaire = append(character.Inventaire, enemy.Equipement.Tunique)
			butinRecupere = append(butinRecupere, enemy.Equipement.Tunique)
		} else {
			butinNonRecupere = append(butinNonRecupere, enemy.Equipement.Tunique)
		}
	}
	if enemy.Equipement.Chapeau != "" {
		if len(character.Inventaire) < 10 {
			character.Inventaire = append(character.Inventaire, enemy.Equipement.Chapeau)
			butinRecupere = append(butinRecupere, enemy.Equipement.Chapeau)
		} else {
			butinNonRecupere = append(butinNonRecupere, enemy.Equipement.Chapeau)
		}
	}
	if len(butinRecupere) > 0 {
		fmt.Printf("\n%sObjets récupérés par %s%s%s :%s\n", couleurs.Green, couleurs.Blue, character.Nom, couleurs.Green, couleurs.Reset)
		for _, item := range butinRecupere {
			fmt.Printf("%s  - %s%s%s\n", couleurs.White, couleurs.Yellow, item, couleurs.Reset)
		}
	}
	if len(butinNonRecupere) > 0 {
		fmt.Printf("\n%sObjets laissés sur place (inventaire plein) :%s\n", couleurs.Red, couleurs.Reset)
		for _, item := range butinNonRecupere {
			fmt.Printf("%s  - %s%s%s\n", couleurs.White, couleurs.Yellow, item, couleurs.Reset)
		}
	}
	if len(butinRecupere) == 0 && len(butinNonRecupere) == 0 {
		fmt.Printf("%sAucun butin trouvé sur l'ennemi.%s\n", couleurs.Yellow, couleurs.Reset)
	}
}

func selectionnerPersonnageButin() *characters.Character {
	fmt.Printf("\n%sQuel personnage doit récupérer le butin ?%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s1. %s%s%s (Inventaire: %d/10)\n", couleurs.White, couleurs.Green, characters.C1.Nom, couleurs.Reset, len(characters.C1.Inventaire))
	fmt.Printf("%s2. %s%s%s (Inventaire: %d/10)\n", couleurs.White, couleurs.Green, characters.C2.Nom, couleurs.Reset, len(characters.C2.Inventaire))
	fmt.Printf("%sVotre choix : %s", couleurs.Blue, couleurs.Reset)
	choixPerso := functionshelper.ReadInput()
	switch strings.ToLower(strings.TrimSpace(choixPerso)) {
	case "1", "1.", strings.ToLower(characters.C1.Nom):
		return characters.C1
	case "2", "2.", strings.ToLower(characters.C2.Nom):
		return characters.C2
	default:
		fmt.Printf("%sChoix invalide, %s%s%s récupère par défaut.%s\n", couleurs.Red, couleurs.Green, characters.C1.Nom, couleurs.Red, couleurs.Reset)
		return characters.C1
	}
}
