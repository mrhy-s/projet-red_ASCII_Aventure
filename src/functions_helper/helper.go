package functionshelper

import (
	characters "ASCII_Aventure/characters"
	"ASCII_Aventure/classes"
	"ASCII_Aventure/couleurs"
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
			fmt.Printf("%sPersonnage non trouvé%s\n", couleurs.Red, couleurs.Reset)
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
			return nil, fmt.Errorf("%spersonnage '%s' non trouvé%s", couleurs.Red, characterName, couleurs.Reset)
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
	fmt.Printf("%s┌──────────────────────────────────────────────────────────────┐%s\n", couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s│%s%s                      Détails de l'item                       %s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	fmt.Printf("%s├──────────────────────────────────────────────────────────────┤%s\n", couleurs.Cyan, couleurs.Reset)
	if item == nil {
		fmt.Printf("%s│%s%sNom:%s %s%-57s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Yellow, itemName, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s%sType:%s %sObjet%s                                                   %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Purple, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		fmt.Printf("%s│%s%sDescription:%s %sUn objet mystérieux%s                              %s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
	} else {
		fmt.Printf("%s│%s%sNom:%s %s%-57s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Yellow, item.Nom, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		if item.Type != "" {
			fmt.Printf("%s│%s%sType:%s %s%-56s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Purple, item.Type, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		}
		if item.Effet != "" {
			fmt.Printf("%s│%s%sEffet:%s %s%-55s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Green, item.Effet, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		}
		if item.Degats > 0 {
			fmt.Printf("%s│%s%sDégâts:%s %s+%-53d%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Red, item.Degats, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
		}
		if item.DurabilitéMaximum > 0 {
			var durabilityColor string
			durabilityPercent := float64(item.DurabilitéesActuelle) / float64(item.DurabilitéMaximum)
			if durabilityPercent >= 0.7 {
				durabilityColor = couleurs.Green
			} else if durabilityPercent >= 0.4 {
				durabilityColor = couleurs.Yellow
			} else {
				durabilityColor = couleurs.Red
			}
			fmt.Printf("%s│%s%sDurabilité:%s %s%d%s/%s%d%s%-48s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, durabilityColor, item.DurabilitéesActuelle, couleurs.Reset, couleurs.White, item.DurabilitéMaximum, couleurs.Reset, "", couleurs.Cyan, couleurs.Reset)
		}
		if item.Description != "" {
			desc := item.Description
			maxWidth := 49
			lines := WrapText(desc, maxWidth)
			for i, line := range lines {
				if i == 0 {
					fmt.Printf("%s│%s%sDescription:%s %s%-49s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.White, line, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
				} else {
					fmt.Printf("%s│%s             %s%-49s%s%s│%s\n", couleurs.Cyan, couleurs.Reset, couleurs.White, line, couleurs.Reset, couleurs.Cyan, couleurs.Reset)
				}
			}
		}
	}
	fmt.Printf("%s└──────────────────────────────────────────────────────────────┘%s\n", couleurs.Cyan, couleurs.Reset)
	if strings.Contains(itemName, "[Spell book]") {
		fmt.Printf("\n%sVoulez-vous apprendre ce livre de sort ?%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		input := strings.ToLower(strings.TrimSpace(ReadInput()))
		if input == "oui" {
			SpellBook("boule de feu", itemName)
		}
	}
	if strings.Contains(itemName, "Chapeau de l'aventurier") {
		fmt.Printf("\n%sVoulez-vous équiper cet objet ?%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		input := strings.ToLower(strings.TrimSpace(ReadInput()))
		if input == "oui" || input == "Oui" {
			var targetCharacter *characters.Character
			var characterName string
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sSur quel personnage souhaitez-vous équiper l'item ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
				characterName = strings.TrimSpace(ReadInput())
				var err error
				targetCharacter, err = GetCharacterByName(characterName)
				if err != nil {
					fmt.Printf("%sErreur: %s%s\n", couleurs.Red, err.Error(), couleurs.Reset)
					return
				}
			} else {
				targetCharacter = characters.C1
			}
			characters.EquipItem(targetCharacter, itemName)
		}
	}
	if strings.Contains(itemName, "Tunique de l'aventurier") {
		fmt.Printf("\n%sVoulez-vous équiper cet objet ?%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		input := strings.ToLower(strings.TrimSpace(ReadInput()))
		if input == "oui" || input == "Oui" {
			var targetCharacter *characters.Character
			var characterName string
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sSur quel personnage souhaitez-vous équiper l'item ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
				characterName = strings.TrimSpace(ReadInput())
				var err error
				targetCharacter, err = GetCharacterByName(characterName)
				if err != nil {
					fmt.Printf("%sErreur: %s%s\n", couleurs.Red, err.Error(), couleurs.Reset)
					return
				}
			} else {
				targetCharacter = characters.C1
			}
			characters.EquipItem(targetCharacter, itemName)
		}
	}
	if strings.Contains(itemName, "Bottes de l'aventurier") {
		fmt.Printf("\n%sVoulez-vous équiper cet objet ?%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		input := strings.ToLower(strings.TrimSpace(ReadInput()))
		if input == "oui" || input == "Oui" {
			var targetCharacter *characters.Character
			var characterName string
			if characters.C2_b && characters.C2 != nil {
				fmt.Printf("\n%sSur quel personnage souhaitez-vous équiper l'item ?%s\n", couleurs.Purple, couleurs.Reset)
				fmt.Printf("%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
				characterName = strings.TrimSpace(ReadInput())
				var err error
				targetCharacter, err = GetCharacterByName(characterName)
				if err != nil {
					fmt.Printf("%sErreur: %s%s\n", couleurs.Red, err.Error(), couleurs.Reset)
					return
				}
			} else {
				targetCharacter = characters.C1
			}
			characters.EquipItem(targetCharacter, itemName)
		}
	}
}

func SpellBook(spell string, itemName string) {
	var targetCharacter *characters.Character
	var characterName string
	if characters.C2_b && characters.C2 != nil {
		fmt.Printf("\n%sSur quel personnage souhaitez-vous utiliser le livre de sort ?%s\n", couleurs.Purple, couleurs.Reset)
		fmt.Printf("%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
		characterName = strings.TrimSpace(ReadInput())
		var err error
		targetCharacter, err = GetCharacterByName(characterName)
		if err != nil {
			fmt.Printf("%sErreur: %s%s\n", couleurs.Red, err.Error(), couleurs.Reset)
			return
		}
	} else {
		targetCharacter = characters.C1
		characterName = characters.C1.Nom
	}
	if contains(targetCharacter.Skill, spell) {
		fmt.Printf("\n%s%s%s connaît déjà le sort : %s%s%s\n", couleurs.Yellow, couleurs.Green, characterName, couleurs.Yellow, spell, couleurs.Reset)
		return
	}
	targetCharacter.Skill = append(targetCharacter.Skill, spell)
	fmt.Printf("\n%s%s%s a appris le sort : %s%s%s\n", couleurs.Green, couleurs.White, characterName, couleurs.Green, spell, couleurs.Reset)
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
	case "Tunique de l'aventurier":
		return items.Tunique_de_laventurier
	case "Chapeau de l'aventurier":
		return items.Chapeau_de_laventurier
	case "Bottes de l'aventurier":
		return items.Bottes_de_laventurier
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
		fmt.Printf("\n%s- Pour créer un nouveau personnage veuillez écrire 'Oui'%s\n", couleurs.Green, couleurs.Reset)
		fmt.Printf("%s- Pour utiliser le personnage par défaut veuillez écrire 'Non'%s\n\n", couleurs.Blue, couleurs.Reset)
		characters.DisplayCharacterTable(*characters.C1)
		fmt.Printf("\n%sVotre réponse :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
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
			fmt.Printf("%sOn a dit 'oui' ou 'non' pas : %s%s%s\n\n", couleurs.Red, couleurs.Yellow, input, couleurs.Reset)
		}
	}
}

func CharacterCreation() *characters.Character {
	var nom string
	if !characters.C2_b {
		for {
			fmt.Printf("%sVeuillez entrer le nom du personnage (uniquement des lettres) :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
			input := ReadInput()
			if IsValidName(input) {
				nom = FormatName(input)
				break
			} else {
				fmt.Printf("%sErreur : Le nom ne peut contenir que des lettres. Veuillez réessayer.%s\n", couleurs.Red, couleurs.Reset)
			}
		}
		var classe string
		for {
			fmt.Printf("\n%sVeuillez entrer la classe du personnage parmi :%s\n", couleurs.Purple, couleurs.Reset)
			fmt.Printf("   %s- Classe:%s %s%s%s %s- %s%s\n", couleurs.White, couleurs.Reset, couleurs.Green, classes.Humain.Nom, couleurs.Reset, couleurs.White, classes.Humain.Description, couleurs.Reset)
			fmt.Printf("   %s- Classe:%s %s%s%s %s- %s%s\n", couleurs.White, couleurs.Reset, couleurs.Green, classes.Elfe.Nom, couleurs.Reset, couleurs.White, classes.Elfe.Description, couleurs.Reset)
			fmt.Printf("   %s- Classe:%s %s%s%s %s- %s%s\n", couleurs.White, couleurs.Reset, couleurs.Green, classes.Nain.Nom, couleurs.Reset, couleurs.White, classes.Nain.Description, couleurs.Reset)
			fmt.Printf("\n%sVotre choix :%s ", couleurs.Blue+couleurs.Bold, couleurs.Reset)
			classInput := ReadInput()
			if IsValidClass(classInput) {
				classe = NormalizeClassName(classInput)
				break
			} else {
				fmt.Printf("\n%sClasse invalide : '%s'%s\n", couleurs.Red, classInput, couleurs.Reset)
				fmt.Printf("%sVeuillez choisir parmi : Humain, Elfe, ou Nain%s\n", couleurs.Red, couleurs.Reset)
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
		fmt.Printf("\n%sVoici votre nouveau personnage :%s\n", couleurs.Green+couleurs.Bold, couleurs.Reset)
		characters.DisplayCharacterTable(*nouveauPersonnage)
		characters.C2_b = true
		return nouveauPersonnage
	} else {
		fmt.Printf("%sVous avez déjà un second personnage%s\n", couleurs.Yellow, couleurs.Reset)
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
