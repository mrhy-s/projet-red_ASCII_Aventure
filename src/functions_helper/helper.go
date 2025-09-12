package functionshelper

import (
	"ASCII_Aventure/characters"
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

func RemoveInventory(characterName string, item string) bool {
	var character *characters.Character
	if characters.C2_b {
		switch characterName {
		case characters.C1.Nom:
			character = characters.C1
		case characters.C2.Nom:
			character = characters.C2
		default:
			fmt.Println("Personnage non trouvé")
			return false
		}
	} else {
		character = characters.C1
	}

	for i, inventoryItem := range character.Inventaire {
		if inventoryItem == item {
			character.Inventaire = RemoveItem(character.Inventaire, i)
			return true
		}
	}
	return false
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

func DisplayItemDetails(itemName string) {
	fmt.Printf("\n=== Détails de l'item ===\n")
	fmt.Printf("Nom: %s\n", itemName)
	switch itemName {
	case "potions de soin":
		fmt.Printf("Type: Consommable\n")
		fmt.Printf("Effet: Restaure 50 points de vie\n")
		fmt.Printf("Description: Une potion magique qui soigne les blessures\n")
	case "épée en fer":
		fmt.Printf("Type: Arme\n")
		fmt.Printf("Dégâts: +10\n")
		fmt.Printf("Durabilité: 100/100\n")
		fmt.Printf("Description: Une épée solide en fer forgé\n")
	default:
		fmt.Printf("Type: Objet\n")
		fmt.Printf("Description: Un objet mystérieux\n")
	}
	fmt.Printf("========================\n")
}
