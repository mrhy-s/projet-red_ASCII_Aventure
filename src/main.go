package main

import "fmt"

type Character struct {
	nom                   string
	classe                string
	niveau                int
	points_de_vie_maximum int
	points_de_vie_actuels int
	inventaire            []string
}

var c1 *Character

func main() {
	c1 = initCharacter("Zinrel", "Elfe", 1, 100, 40, []string{"potions de soin", "potions de soin", "potions de soin"})
	fmt.Println(c1)
}

func initCharacter(nom string, classe string, niveau int, points_de_vie_maximum int, points_de_vie_actuels int, inventaire []string) *Character {
	character_template := &Character{
		nom:                   nom,
		classe:                classe,
		niveau:                niveau,
		points_de_vie_maximum: points_de_vie_maximum,
		points_de_vie_actuels: points_de_vie_actuels,
		inventaire:            []string{},
	}
	return character_template
}
