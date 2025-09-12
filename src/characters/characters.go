package characters

import (
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

func InitCharacter(nom string, classe string, niveau int, pointsDeVieMaximum int, pointsDeVieActuels int, inventaire []string) *Character {
	characterTemplate := &Character{
		Nom:                nom,
		Classe:             classe,
		Niveau:             niveau,
		PointsDeVieMaximum: pointsDeVieMaximum,
		PointsDeVieActuels: pointsDeVieActuels,
		Inventaire:         inventaire,
	}
	return characterTemplate
}

func DisplayCharacterTable(character Character) {
	inventaireStr := strings.Join(character.Inventaire, ", ")
	// Largeur totale du cadre : 85 caractères
	// Largeur utile pour le contenu : 85 - 4 (bordures et espaces) = 81 caractères
	// Largeur pour les valeurs après "   - XXX : " : environ 65-67 caractères selon le label
	// const totalWidth = 85
	// const contentWidth = 81
	inventaireMaxWidth := 62 // 81 - 18 - 1 = 62
	// Tronquer l'inventaire si nécessaire
	if utf8.RuneCountInString(inventaireStr) > inventaireMaxWidth {
		// Garder de la place pour "..."
		targetLength := inventaireMaxWidth - 3
		runes := []rune(inventaireStr)
		if len(runes) > targetLength {
			inventaireStr = string(runes[:targetLength]) + "..."
		}
	}

	fmt.Println("┌─────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Printf("│%-85s│\n", "Caractéristiques du personnage :")
	fmt.Printf("│   - Nom : %-74s│\n", character.Nom)
	fmt.Printf("│   - Classe : %-71s│\n", character.Classe)
	fmt.Printf("│   - Niveau : %-71d│\n", character.Niveau)
	fmt.Printf("│   - Points de vie maximum : %-56d│\n", character.PointsDeVieMaximum)
	fmt.Printf("│   - Points de vie actuels : %-56d│\n", character.PointsDeVieActuels)
	fmt.Printf("│   - Inventaire : [%-65s]│\n", inventaireStr)
	fmt.Printf("│%-85s│\n", "")
	fmt.Println("└─────────────────────────────────────────────────────────────────────────────────────┘")
}

func IsDead() {
	var deadCharacter *Character
	if C2_b {
		if C1.PointsDeVieActuels <= 0 {
			deadCharacter = C1
		} else if C2.PointsDeVieActuels <= 0 {
			deadCharacter = C2
		}
	} else {
		if C1.PointsDeVieActuels <= 0 {
			deadCharacter = C1
		}
	}
	if deadCharacter != nil {
		fmt.Printf("\n\nOh snap ! Votre personnage %s est mort (╥﹏╥)\n\n", deadCharacter.Nom)
		displayWastedMessage()
		deadCharacter.PointsDeVieMaximum = deadCharacter.PointsDeVieMaximum / 2
		if deadCharacter.PointsDeVieActuels > deadCharacter.PointsDeVieMaximum {
			deadCharacter.PointsDeVieActuels = deadCharacter.PointsDeVieMaximum
		}
		deadCharacter.PointsDeVieActuels = rand.Intn(deadCharacter.PointsDeVieMaximum-11) + 10
		fmt.Printf("\n\nVotre personnage est ressuscité avec %v sur %v points de vie... ", deadCharacter.PointsDeVieActuels, deadCharacter.PointsDeVieMaximum)
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

	fmt.Print(wastedArt)
}
