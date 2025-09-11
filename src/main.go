package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Character struct {
	nom                   string
	classe                string
	niveau                int
	points_de_vie_maximum int
	points_de_vie_actuels int
	inventaire            []string
}

var c1 *Character
var c2 *Character
var c2_b bool

var c_temp_name string
var c_temp_classe string
var c_temp_niveau int
var c_temp_points_de_vie_maximum int
var c_temp_points_de_vie_actuels int
var c_temp_inventaire []string

func main() {
	IsFirst := true
	c1 = initCharacter("Zinrel", "Elfe", 1, 100, 40, []string{"potions de soin", "potions de soin", "potions de soin"})
	if IsFirst {
		min_pt_vie_rand, max_pt_vie_rand := 95, 105
		fmt.Print("Voulez vous utiliser le personnage par défaut ? Ou voulez vous créer un nouveau personnage ?\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'")
		fmt.Print("\n")
		displayInfo(c1.nom)
		fmt.Print("\nVotre réponse : ")
		input := readInput()
		if input == "Oui" || input == "Oui." || input == "oui" || input == "oui." {
			fmt.Print("\nVeuillez entrer le nom du personnage : ")
			c_temp_name = readInput()
			fmt.Print("\nVeuillez entrer la classe du personnage : ")
			c_temp_classe = readInput()
			fmt.Printf("\nVotre personnage s'appelle : %s\nLa classe de votre personnage est : %s", c_temp_name, c_temp_classe)
			fmt.Print("\nVotre personnage commence au niveau 0.")
			c_temp_niveau = 0
			c_temp_points_de_vie_maximum = randomBetween(min_pt_vie_rand, max_pt_vie_rand)
			c_temp_points_de_vie_actuels = c_temp_points_de_vie_maximum - randomBetween(20, 70)
			fmt.Printf("\nVotre personnage avec %d/%d points de vie.", c_temp_points_de_vie_actuels, c_temp_points_de_vie_maximum)
			c_temp_inventaire = []string{"potions de soin", "potions de soin", "potions de soin"}
			fmt.Print("\nvous commencez avec 'Potion de soin x3'")
			c2_b = true
			c2 = initCharacter(c_temp_name, c_temp_classe, c_temp_niveau, c_temp_points_de_vie_maximum, c_temp_points_de_vie_actuels, c_temp_inventaire)
		}
		IsFirst = false
	}
	menu()
}

// ================
// fonction actions
// ================

func initCharacter(nom string, classe string, niveau int, points_de_vie_maximum int, points_de_vie_actuels int, inventaire []string) *Character {
	character_template := &Character{
		nom:                   nom,
		classe:                classe,
		niveau:                niveau,
		points_de_vie_maximum: points_de_vie_maximum,
		points_de_vie_actuels: points_de_vie_actuels,
		inventaire:            inventaire,
	}
	return character_template
}

func displayInfo(s string) {
	s = strings.TrimSpace(s)
	if strings.Contains(strings.ToLower(s), "liste") {
		fmt.Printf("\nPersonnage 1 :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie : %d/%d\n   - Inventaire : %v\n",
			c1.nom, c1.classe, c1.niveau, c1.points_de_vie_actuels, c1.points_de_vie_maximum, c1.inventaire)

		if c2_b && c2 != nil {
			fmt.Printf("\nPersonnage 2 :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie : %d/%d\n   - Inventaire : %v\n",
				c2.nom, c2.classe, c2.niveau, c2.points_de_vie_actuels, c2.points_de_vie_maximum, c2.inventaire)
		}
		return
	}
	if s == c1.nom {
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c1.nom, c1.classe, c1.niveau, c1.points_de_vie_maximum, c1.points_de_vie_actuels, c1.inventaire)
		return
	}
	if c2_b && c2 != nil && s == c2.nom {
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c2.nom, c2.classe, c2.niveau, c2.points_de_vie_maximum, c2.points_de_vie_actuels, c2.inventaire)
		return
	}
	fmt.Println("Personnage non trouvé")
}

func accessInventory(s string) {
	if c2_b {
		switch s {
		case c1.nom:
			fmt.Printf("\nVoici l'inventaire du personnage %s :\n%v", c1.nom, c1.inventaire)
		case c2.nom:
			fmt.Printf("\nVoici l'inventaire du personnage %s :\n%v", c2.nom, c2.inventaire)
		}
	} else {
		fmt.Printf("\nVoici l'inventaire du personnage %s :\n%v", c1.nom, c1.inventaire)
	}
}

func takePot(characterName string) {
	var character *Character
	if c2_b {
		switch characterName {
		case c1.nom:
			character = c1
		case c2.nom:
			character = c2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = c1
	}
	potionIndex := -1
	for i, item := range character.inventaire {
		if item == "potions de soin" {
			potionIndex = i
			break
		}
	}
	if potionIndex == -1 {
		fmt.Printf("\n%s n'a pas de potion de soin dans son inventaire.\n", character.nom)
		return
	}
	character.points_de_vie_actuels += 50
	if character.points_de_vie_actuels > character.points_de_vie_maximum {
		character.points_de_vie_actuels = character.points_de_vie_maximum
	}
	character.inventaire = removeItem(character.inventaire, potionIndex)
	fmt.Printf("\n%s a utilisé une potion de soin ! (+50 PV)\n", character.nom)
	fmt.Printf("Points de vie: %d/%d\n", character.points_de_vie_actuels, character.points_de_vie_maximum)
}

// ==============
// fonction menus
// ==============

func menu() {
	fmt.Print("\n\nPour sélectionner une option, veuillez entrer le numéro de celle-ci : ")
	fmt.Print("\n1. Afficher les informations du personnage")
	fmt.Print("\n2. Accéder au contenu de l'inventaire")
	fmt.Print("\n3. Utiliser une potion de soin")
	fmt.Print("\n4. Quitter")
	fmt.Print("\n\nVotre réponse : ")
	option := readInput()
	switch option {
	case "1", "1.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'\n")
		fmt.Print("\nVotre réponse : ")
		commande := readInput()
		displayInfo(commande)
		menu()
	case "2", "2.":
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher son inventaire :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'\n")
		fmt.Print("\nVotre réponse : ")
		commande := readInput()
		accessInventory(commande)
		menu()
	case "3", "3.":
		fmt.Print("Nom du personnage : ")
		characterName := readInput()
		takePot(characterName)
		menu()
	case "4", "4.":
		// logique marchand
	case "5", "5.":
		return
	default:
		fmt.Println("Option invalide")
		menu()
	}
}

// ================
// fonction helpers
// ================

func randomBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func removeItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
