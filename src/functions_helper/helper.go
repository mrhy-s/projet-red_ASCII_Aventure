package functionshelper

import (
	characters "ASCII_Aventure/characters"
	"ASCII_Aventure/classes"
	"ASCII_Aventure/items"
	"ASCII_Aventure/startscreen"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

// ================
// fonction helpers
// ================

func RandomBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RemoveItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func AddInventory(characterName string, item string) {
	var character *characters.Character
	if characters.C2_b {
		switch characterName {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = characters.C1
	}
	character.Inventaire = append(character.Inventaire, item)
}

func HasItemInInventory(character *characters.Character, item string) int {
	inventaireStr := strings.Join(character.Inventaire, ", ")
	index := Index(inventaireStr, item)
	return index
}

func Index(s, toFind string) int {
	if toFind == "" {
		return 0
	}

	sRunes := []rune(s)
	toFindRunes := []rune(toFind)

	if len(toFindRunes) > len(sRunes) {
		return -1
	}

	for i := 0; i <= len(sRunes)-len(toFindRunes); i++ {
		match := true
		for j := 0; j < len(toFindRunes); j++ {
			if sRunes[i+j] != toFindRunes[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func GetCharacterByName(characterName string) (*characters.Character, error) {
	name := strings.ToLower(strings.TrimSpace(characterName))
	if characters.C2_b {
		switch name {
		case strings.ToLower(characters.C1.Nom):
			return characters.C1, nil
		case strings.ToLower(characters.C2.Nom):
			return characters.C2, nil
		default:
			return nil, fmt.Errorf("personnage '%s' non trouvé", characterName)
		}
	}
	return characters.C1, nil
}

func RemoveInventory(characterName string, item string) bool {
	character, err := GetCharacterByName(characterName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	for i, inventoryItem := range character.Inventaire {
		if inventoryItem == item {
			character.Inventaire = RemoveItem(character.Inventaire, i)
			return true
		}
	}
	return false
}

func DisplayItemDetails(itemName string) {
	item := GetItemByName(strings.ToLower(itemName))
	fmt.Println("┌──────────────────────────────────────────────────────────────┐")
	fmt.Println("│                      Détails de l'item                       │")
	fmt.Println("├──────────────────────────────────────────────────────────────┤")
	if item == nil {
		fmt.Printf("│Nom: %-57s│\n", itemName)
		fmt.Println("│Type: Objet                                                   │")
		fmt.Println("│Description: Un objet mystérieux                              │")
	} else {
		fmt.Printf("│Nom: %-57s│\n", item.Nom)
		if item.Type != "" {
			fmt.Printf("│Type: %-56s│\n", item.Type)
		}
		if item.Effet != "" {
			fmt.Printf("│Effet: %-55s│\n", item.Effet)
		}
		if item.Degats > 0 {
			fmt.Printf("│Dégâts: +%-53d│\n", item.Degats)
		}
		if item.DurabilitéMaximum > 0 {
			fmt.Printf("│Durabilité: %d/%-48d│\n", item.DurabilitéesActuelle, item.DurabilitéMaximum)
		}
		if item.Description != "" {
			desc := item.Description
			maxWidth := 49
			lines := WrapText(desc, maxWidth)
			for i, line := range lines {
				if i == 0 {
					fmt.Printf("│Description: %-49s│\n", line)
				} else {
					fmt.Printf("│             %-49s│\n", line)
				}
			}
		}
	}
	fmt.Println("└──────────────────────────────────────────────────────────────┘")
	if strings.Contains(itemName, "[Spell book]") {
		fmt.Print("\nVoulez-vous apprendre ce livre de sort ? ")
		input := strings.ToLower(strings.TrimSpace(ReadInput()))
		if input == "oui" {
			SpellBook("boule de feu", itemName)
		}
	}
}

func SpellBook(spell string, itemName string) {
	var targetCharacter *characters.Character
	var characterName string
	if characters.C2_b && characters.C2 != nil {
		fmt.Print("\nSur quel personnage souhaitez-vous utiliser le livre de sort ?\nVotre choix : ")
		characterName = strings.TrimSpace(ReadInput())
		var err error
		targetCharacter, err = GetCharacterByName(characterName)
		if err != nil {
			fmt.Printf("Erreur: %s\n", err.Error())
			return
		}
	} else {
		targetCharacter = characters.C1
		characterName = characters.C1.Nom
	}
	if contains(targetCharacter.Skill, spell) {
		fmt.Printf("\n%s connaît déjà le sort : %s\n", characterName, spell)
		return
	}
	targetCharacter.Skill = append(targetCharacter.Skill, spell)
	fmt.Printf("\n%s a appris le sort : %s\n", characterName, spell)
	RemoveInventory(characterName, itemName)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetItemByName(itemName string) *items.Item {
	switch itemName {
	case "potion de soin":
		return items.Potion_de_soin
	case "potion de poison":
		return items.Potion_de_poison
	case "épée en fer":
		return items.Epee_en_fer
	case "boule de feu", "Boule de feu", "[Spell book] > Boule de feu", "[spell book] > boule de feu":
		return items.Spell_book_bdf
	default:
		return nil
	}
}

func WrapText(text string, width int) []string {
	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{text}
	}
	currentLine := ""
	for _, word := range words {
		if currentLine == "" {
			currentLine = word
		} else {
			testLine := currentLine + " " + word
			if len(testLine) <= width {
				currentLine = testLine
			} else {
				lines = append(lines, currentLine)
				currentLine = word
			}
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func IsValidClass(input string) bool {
	validClasses := []string{
		strings.ToLower(classes.Humain.Nom),
		strings.ToLower(classes.Elfe.Nom),
		strings.ToLower(classes.Nain.Nom),
	}

	inputLower := strings.ToLower(strings.TrimSpace(input))

	for _, class := range validClasses {
		if inputLower == class {
			return true
		}
	}
	return false
}

func NormalizeClassName(input string) string {
	inputLower := strings.ToLower(strings.TrimSpace(input))
	switch inputLower {
	case "humain":
		return "Humain"
	case "elfe":
		return "Elfe"
	case "nain":
		return "Nain"
	default:
		return input
	}
}

func IsValidName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}
	matched, _ := regexp.MatchString(`^[a-zA-ZÀ-ÿ\s]+$`, strings.TrimSpace(name))
	return matched
}

func FormatName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}
	words := strings.Fields(name)
	var formattedWords []string

	for _, word := range words {
		if len(word) > 0 {
			formatted := strings.ToUpper(string([]rune(word)[0])) + strings.ToLower(string([]rune(word)[1:]))
			formattedWords = append(formattedWords, formatted)
		}
	}
	return strings.Join(formattedWords, " ")
}

func CharacterSelection() {
	for {
		fmt.Print("\n- Pour créer un nouveau personnage veuillez écrire 'Oui'\n- Pour utiliser le personnage par défaut veuillez écrire 'Non'\n\n")
		characters.DisplayCharacterTable(*characters.C1)
		fmt.Print("\nVotre réponse : ")
		input := ReadInput()

		switch input {
		case "Oui", "Oui.", "oui", "oui.":
			characters.C2_b = true
			characters.C2 = CharacterCreation()
			return
		case "Non", "Non.", "non", "non.":
			startscreen.ClearScreen()
			return
		default:
			fmt.Printf("On a dit 'oui' ou 'non' pas : %s (╥﹏╥)\n\n", input)
		}
	}
}

func CharacterCreation() *characters.Character {
	var nom string
	if !characters.C2_b {
		for {
			fmt.Print("Veuillez entrer le nom du personnage (uniquement des lettres) : ")
			input := ReadInput()
			if IsValidName(input) {
				nom = FormatName(input)
				break
			} else {
				fmt.Println("Erreur : Le nom ne peut contenir que des lettres. Veuillez réessayer.")
			}
		}
		var classe string
		for {
			fmt.Print("\nVeuillez entrer la classe du personnage parmi : \n")
			fmt.Printf("   - Classe: %s - %s\n", classes.Humain.Nom, classes.Humain.Description)
			fmt.Printf("   - Classe: %s - %s\n", classes.Elfe.Nom, classes.Elfe.Description)
			fmt.Printf("   - Classe: %s - %s\n", classes.Nain.Nom, classes.Nain.Description)
			fmt.Print("\nVotre choix : ")
			classInput := ReadInput()
			if IsValidClass(classInput) {
				classe = NormalizeClassName(classInput)
				break
			} else {
				fmt.Printf("\nClasse invalide : '%s'\n", classInput)
				fmt.Println("Veuillez choisir parmi : Humain, Elfe, ou Nain")
			}
		}
		niveau := 1
		var pointsDeVieMaximum, pointsDeVieActuels int
		var inventaire []string
		var skill []string
		piècesdor := 100
		switch classe {
		case "Humain":
			pointsDeVieMaximum = RandomBetween(95, 105)
			skill = []string{"Coup de poing"}
			inventaire = []string{"potion de soin", "potion de soin", "potion de soin"}
		case "Elfe":
			pointsDeVieMaximum = RandomBetween(75, 85)
			skill = []string{"Tir à l'arc"}
			inventaire = []string{"potion de soin", "potion de soin"}
		case "Nain":
			pointsDeVieMaximum = RandomBetween(115, 125)
			skill = []string{"Coup de hache"}
			inventaire = []string{"potion de soin", "potion de soin"}
		default: // par défaut (unused)
			pointsDeVieMaximum = RandomBetween(95, 105)
			skill = []string{"Coup de poing"}
			inventaire = []string{"potion de soin", "potion de soin", "potion de soin"}
		}
		pointsDeVieActuels = pointsDeVieMaximum - RandomBetween(20, 70)
		if pointsDeVieActuels < 1 {
			pointsDeVieActuels = 1
		}
		nouveauPersonnage := characters.InitCharacter(nom, classe, niveau, pointsDeVieMaximum, pointsDeVieActuels, inventaire, skill, piècesdor)
		fmt.Println("\nVoici votre nouveau personnage :")
		characters.DisplayCharacterTable(*nouveauPersonnage)
		characters.C2_b = true
		return nouveauPersonnage
	} else {
		fmt.Println("Vous avrez déjà un second personnage")
		return nil
	}
}

/// =====
/// Debug
/// =====
/*
func InstaKill(characterName string) {
	var character *characters.Character
	if characters.C2_b {
		switch characterName {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return
		}
	} else {
		character = characters.C1
	}
	character.PointsDeVieActuels = 0
}
*/
