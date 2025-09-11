package main

import (
	"fmt"
	"math/rand"
	"time"
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

var c_temp_name string
var c_temp_classe string
var c_temp_niveau int
var c_temp_points_de_vie_maximum int
var c_temp_points_de_vie_actuels int
var c_temp_inventaire []string

func main() {
	IsFirst := true
	var input string
	c1 = initCharacter("Zinrel", "Elfe", 1, 100, 40, []string{"potions de soin", "potions de soin", "potions de soin"})
	if IsFirst {
		rand.Seed(time.Now().UnixNano())
		min_pt_vie_rand, max_pt_vie_rand := 95, 105
		fmt.Print("Voulez vous utiliser le personnage par défaut ? Ou voulez vous créer un nouveau personnage ?\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'")
		fmt.Print("\n")
		displayInfo(c1.nom)
		fmt.Print("\nVotre réponse : ")
		fmt.Scan(&input)
		if input == "Oui" || input == "Oui." || input == "oui" || input == "oui." {
			fmt.Print("\nVeuillez entrer le nom du personnage : ")
			fmt.Scan(&c_temp_name)
			fmt.Print("\nVeuillez entrer la classe du personnage : ") //implémenter la liste des classes
			fmt.Scan(&c_temp_classe)
			fmt.Printf("\nVotre personnage s'appelle : %s\nLa classe de votre personnage est : %s", c_temp_name, c_temp_classe)
			fmt.Print("\nVotre personnage commence au niveau 0.")
			c_temp_niveau = 0
			c_temp_points_de_vie_maximum = randomBetween(min_pt_vie_rand, max_pt_vie_rand)
			c_temp_points_de_vie_actuels = c_temp_points_de_vie_maximum - randomBetween(20, 70)
			fmt.Printf("\nVotre personnage avec %d/%d points de vie.", c_temp_points_de_vie_actuels, c_temp_points_de_vie_maximum)
			c_temp_inventaire = []string{"potions de soin", "potions de soin", "potions de soin"}
			fmt.Print("\nvous commencez avec 'Potion de soin x3'")
			c2 = initCharacter(c_temp_name, c_temp_classe, c_temp_niveau, c_temp_points_de_vie_maximum, c_temp_points_de_vie_actuels, c_temp_inventaire)
		}
		IsFirst = false
	}
	menu()
}

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

func menu() {
	fmt.Print("\n\nPour sélectionner une option, veuillez entrer le numéro de celle-ci : ")
	fmt.Print("\n1. Afficher les informations du personnage")
	fmt.Print("\n2. Accéder au contenu de l’inventaire")
	fmt.Print("\n3. Quitter")
	var option string
	fmt.Scan(&option)
	var commande string
	if option == "1" || option == "1." {
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher ses caractéristiques :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'")
		fmt.Scan(&commande)
		displayInfo(commande)
	} else if option == "2" || option == "2." {
		fmt.Print("Veuillez sélectionner un personnage de la liste pour afficher son inventaire :\nPour afficher la liste des personnages veuillez entrer 'Liste des personnages'")
		fmt.Scan(&commande)
		accessInventory(commande)
	} else if option == "3" || option == "3." {
		return
	}
}

func commandList() {

}

func accessInventory(s string) {
	if s == c1.nom {
		fmt.Print("\nVoici l'inventaire du personnage %s :\n%v", c1.nom, c1.inventaire)
	} else if s == c2.nom {
		fmt.Print("\nVoici l'inventaire du personnage %s :\n%v", c2.nom, c2.inventaire)
	}
}

func displayInfo(s string) {
	if s == "Liste des personnages" || s == "liste des personnages" || s == "Liste" || s == "personnages" || s == "Personnages" {
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c1.nom, c1.classe, c1.niveau, c1.points_de_vie_maximum, c1.points_de_vie_actuels, c1.inventaire)
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c2.nom, c2.classe, c2.niveau, c2.points_de_vie_maximum, c2.points_de_vie_actuels, c2.inventaire)
		return
	}
	if s == c1.nom {
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c1.nom, c1.classe, c1.niveau, c1.points_de_vie_maximum, c1.points_de_vie_actuels, c1.inventaire)
		return
	}
	if s == c2.nom {
		fmt.Printf("\nCaractéristiques du personnage :\n   - Nom : %s\n   - Classe : %s\n   - Niveau : %d\n   - Points de vie maximum : %d\n   - Points de vie actuels : %d\n   - Inventaire : %v\n", c2.nom, c2.classe, c2.niveau, c2.points_de_vie_maximum, c2.points_de_vie_actuels, c2.inventaire)
		return
	}
}

func randomBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}
