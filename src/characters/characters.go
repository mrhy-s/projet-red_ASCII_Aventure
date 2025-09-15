package characters

import (
	"ASCII_Aventure/couleurs"
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

type Character struct {
	Nom                string
	Classe             string
	Niveau             int
	PointsDeVieMaximum int
	PointsDeVieActuels int
	Inventaire         []string
	Skill              []string
	PiècesDOr          int
}

var C1 *Character
var C2 *Character
var C2_b bool

var C_temp_name string
var C_temp_classe string
var C_temp_niveau int
var C_temp_points_de_vie_maximum int
var C_temp_points_de_vie_actuels int
var C_temp_inventaire []string
var C_temp_skill []string

func InitCharacter(nom string, classe string, niveau int, pointsDeVieMaximum int, pointsDeVieActuels int, inventaire []string, skill []string, piècesDOr int) *Character {
	return &Character{
		Nom:                nom,
		Classe:             classe,
		Niveau:             niveau,
		PointsDeVieMaximum: pointsDeVieMaximum,
		PointsDeVieActuels: pointsDeVieActuels,
		Inventaire:         inventaire,
		Skill:              skill,
		PiècesDOr:          piècesDOr,
	}
}

func DisplayCharacterTable(character Character) {
	inventaireStr := strings.Join(character.Inventaire, ", ")
	skillStr := strings.Join(character.Skill, ", ")
	inventaireMaxWidth := 62
	skillMaxWidth := 67
	if utf8.RuneCountInString(inventaireStr) > inventaireMaxWidth {
		targetLength := inventaireMaxWidth - 3
		runes := []rune(inventaireStr)
		if len(runes) > targetLength {
			inventaireStr = string(runes[:targetLength]) + "..."
		}
	}
	if utf8.RuneCountInString(skillStr) > skillMaxWidth {
		targetLength := skillMaxWidth - 1
		runes2 := []rune(skillStr)
		if len(runes2) > targetLength {
			skillStr = string(runes2[:targetLength]) + "..."
		}
	}
	fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s%s                         Caractéristiques du personnage :                            %s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Nom :%s %s%-74s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Green, character.Nom, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Classe :%s %s%-71s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Purple, character.Classe, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Niveau :%s %s%-71d%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, character.Niveau, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	var pvColor string
	pvPercent := float64(character.PointsDeVieActuels) / float64(character.PointsDeVieMaximum)
	if pvPercent >= 0.7 {
		pvColor = couleurs.Green
	} else if pvPercent >= 0.4 {
		pvColor = couleurs.Yellow
	} else {
		pvColor = couleurs.Red
	}
	fmt.Printf("%s│%s   %s- Points de vie maximum :%s %s%-56d%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.White, character.PointsDeVieMaximum, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Points de vie actuels :%s %s%-56d%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, pvColor, character.PointsDeVieActuels, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	var invColor string
	if len(character.Inventaire) == 0 {
		invColor = couleurs.Red
		inventaireStr = "Vide"
	} else {
		invColor = couleurs.White
	}
	fmt.Printf("%s│%s   %s- Inventaire :%s [%s%-65s%s]%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, invColor, inventaireStr, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	var skillColor string
	if len(character.Skill) == 0 {
		skillColor = couleurs.Red
		skillStr = "Aucune"
	} else {
		skillColor = couleurs.White
	}
	fmt.Printf("%s│%s   %s- Skill :%s [%s%-70s%s]%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, skillColor, skillStr, couleurs.Reset, couleurs.Cyan, couleurs.Reset)

	fmt.Printf("%s│%s   %s- Pièces d'Or :%s %s%-65d%s %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Yellow, character.PiècesDOr, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Cyan, couleurs.Reset)
}

func IsDead() {
	var deadCharacter *Character
	if C2_b {
		if C1.PointsDeVieActuels <= 0 {
			deadCharacter = C1
		} else if C2 != nil && C2.PointsDeVieActuels <= 0 {
			deadCharacter = C2
		}
	} else {
		if C1.PointsDeVieActuels <= 0 {
			deadCharacter = C1
		}
	}
	if deadCharacter != nil {
		fmt.Printf("\n\n%sOh snap ! Votre personnage %s%s%s est mort%s\n\n", couleurs.Red+couleurs.Bold, couleurs.Green, deadCharacter.Nom, couleurs.Red+couleurs.Bold, couleurs.Reset)
		displayWastedMessage()
		deadCharacter.PointsDeVieMaximum = deadCharacter.PointsDeVieMaximum / 2
		if deadCharacter.PointsDeVieActuels > deadCharacter.PointsDeVieMaximum {
			deadCharacter.PointsDeVieActuels = deadCharacter.PointsDeVieMaximum
		}
		deadCharacter.PointsDeVieActuels = rand.Intn(deadCharacter.PointsDeVieMaximum-11) + 10
		fmt.Printf("\n\n%sVotre personnage est ressuscité avec %s%v%s sur %s%v%s points de vie...%s ", couleurs.Green, couleurs.White, deadCharacter.PointsDeVieActuels, couleurs.Green, couleurs.White, deadCharacter.PointsDeVieMaximum, couleurs.Green, couleurs.Reset)
	}
}

func displayWastedMessage() {
	const wastedArt = `░██       ░██                          ░██                      ░██ 
░██       ░██                          ░██                      ░██ 
░██  ░██  ░██  ░██████    ░███████  ░████████  ░███████   ░████████ 
░██ ░████ ░██       ░██  ░██           ░██    ░██    ░██ ░██    ░██ 
░██░██ ░██░██  ░███████   ░███████     ░██    ░█████████ ░██    ░██ 
░████   ░████ ░██   ░██         ░██    ░██    ░██        ░██   ░███ 
░███     ░███  ░█████░██  ░███████      ░████  ░███████   ░█████░██ `

	fmt.Printf("%s%s%s%s\n", couleurs.Red, couleurs.Bold, wastedArt, couleurs.Reset)
}
