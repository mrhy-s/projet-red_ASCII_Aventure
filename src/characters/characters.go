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
	Equipment          Equipement
	InventaireMaxSlots int
	UpgradesUtilisés   int
}

type Equipement struct {
	Chapeau  string
	Tunique  string
	Bottes   string
	Arme     string
	Bouclier string
}

type Monster struct {
	Nom                string
	Classe             string
	PointsDeVieMaximum int
	PointsDeVieActuels int
	Inventaire         []string
	Skill              []string
	Equipement         Equipement
}

var C1 *Character
var C2 *Character
var C2_b bool
var Gobelin *Monster

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
		Equipment:          Equipement{},
		InventaireMaxSlots: 10,
		UpgradesUtilisés:   0,
	}
}

func InitMonster(nom string, classe string, niveau int, pointsDeVieMaximum int, pointsDeVieActuels int, inventaire []string, skill []string) *Monster {
	return &Monster{
		Nom:                nom,
		Classe:             classe,
		PointsDeVieMaximum: pointsDeVieMaximum,
		PointsDeVieActuels: pointsDeVieActuels,
		Inventaire:         inventaire,
		Skill:              skill,
		Equipement:         Equipement{},
	}
}

func GobelinPattern() {

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
	fmt.Printf("%s┌─────────────────────────────────────────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s%s                         Caractéristiques du personnage :                            %s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Bold, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Nom :%s %s%-74s%s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Green, character.Nom, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Classe :%s %s%-71s%s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Purple, character.Classe, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Niveau :%s %s%-71d%s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Blue, character.Niveau, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	var pvColor string
	pvPercent := float64(character.PointsDeVieActuels) / float64(character.PointsDeVieMaximum)
	if pvPercent >= 0.7 {
		pvColor = couleurs.Green
	} else if pvPercent >= 0.4 {
		pvColor = couleurs.Yellow
	} else {
		pvColor = couleurs.Red
	}
	fmt.Printf("%s│%s   %s- Points de vie maximum :%s %s%-56d%s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.White, character.PointsDeVieMaximum, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s   %s- Points de vie actuels :%s %s%-56d%s%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, pvColor, character.PointsDeVieActuels, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	var invColor string
	if len(character.Inventaire) == 0 {
		invColor = couleurs.Red
		inventaireStr = "Vide"
	} else {
		invColor = couleurs.White
	}
	fmt.Printf("%s│%s   %s- Inventaire :%s [%s%-65s%s]%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, invColor, inventaireStr, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	var skillColor string
	if len(character.Skill) == 0 {
		skillColor = couleurs.Red
		skillStr = "Aucune"
	} else {
		skillColor = couleurs.White
	}
	fmt.Printf("%s│%s   %s- Skill :%s [%s%-70s%s]%s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, skillColor, skillStr, couleurs.Reset, couleurs.Blue, couleurs.Reset)

	fmt.Printf("%s│%s   %s- Pièces d'Or :%s %s%-65d%s %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset, couleurs.Yellow, character.PiècesDOr, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s                                                                                     %s│%s\n", couleurs.Blue, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s└─────────────────────────────────────────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
}

func IsDead() bool {
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
		return true
	}
	return false
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

func EquipItem(character *Character, itemName string) bool {
	itemLower := strings.ToLower(itemName)
	var bonusPV int
	var oldItem string
	switch {
	// Chapeaux/Casques
	case strings.Contains(itemLower, "chapeau"):
		bonusPV = 10
		oldItem = character.Equipment.Chapeau
		character.Equipment.Chapeau = itemName
	case strings.Contains(itemLower, "casque"):
		bonusPV = 15
		oldItem = character.Equipment.Chapeau
		character.Equipment.Chapeau = itemName
	// Tuniques/Armures
	case strings.Contains(itemLower, "tunique"):
		bonusPV = 25
		oldItem = character.Equipment.Tunique
		character.Equipment.Tunique = itemName
	case strings.Contains(itemLower, "cuir bouilli"):
		bonusPV = 35
		oldItem = character.Equipment.Tunique
		character.Equipment.Tunique = itemName
	// Bottes
	case strings.Contains(itemLower, "bottes"):
		bonusPV = 15
		oldItem = character.Equipment.Bottes
		character.Equipment.Bottes = itemName
	// Armes
	case strings.Contains(itemLower, "épée") || strings.Contains(itemLower, "epee"):
		bonusPV = 0 // Les armes ne donnent pas de PV
		oldItem = character.Equipment.Arme
		character.Equipment.Arme = itemName
	case strings.Contains(itemLower, "dague"):
		bonusPV = 0
		oldItem = character.Equipment.Arme
		character.Equipment.Arme = itemName
	case strings.Contains(itemLower, "gourdin"):
		bonusPV = 0
		oldItem = character.Equipment.Arme
		character.Equipment.Arme = itemName
	case strings.Contains(itemLower, "arc"):
		bonusPV = 0
		oldItem = character.Equipment.Arme
		character.Equipment.Arme = itemName
	default:
		fmt.Printf("%sCet objet ne peut pas être équipé : %s%s%s\n", couleurs.Red, couleurs.Yellow, itemName, couleurs.Reset)
		return false
	}
	// Remettre l'ancien équipement dans l'inventaire s'il y en avait un
	if oldItem != "" && oldItem != "Poing" {
		character.Inventaire = append(character.Inventaire, oldItem)
		// Retirer le bonus de l'ancien équipement
		oldBonus := GetEquipmentBonus(oldItem)
		character.PointsDeVieMaximum -= oldBonus
		fmt.Printf("%s%s déséquipe %s%s%s (-%d PV max)%s\n", couleurs.Yellow, character.Nom, couleurs.White, oldItem, couleurs.Yellow, oldBonus, couleurs.Reset)
	}
	// Ajouter le bonus du nouvel équipement
	character.PointsDeVieMaximum += bonusPV
	// Ajuster les PV actuels si ils dépassent le nouveau maximum
	if character.PointsDeVieActuels > character.PointsDeVieMaximum {
		character.PointsDeVieActuels = character.PointsDeVieMaximum
	}
	// Affichage selon le type d'équipement
	if bonusPV > 0 {
		fmt.Printf("\n%s%s équipe %s%s%s (+%d PV max)%s\n", couleurs.Green, character.Nom, couleurs.Yellow, itemName, couleurs.Green, bonusPV, couleurs.Reset)
	} else {
		fmt.Printf("\n%s%s équipe %s%s%s%s\n", couleurs.Green, character.Nom, couleurs.Yellow, itemName, couleurs.Green, couleurs.Reset)
	}
	// Retirer l'objet de l'inventaire
	removeInventory(character.Nom, itemName)
	return true
}

// import cycle not allowed
func removeInventory(characterName string, item string) bool {
	character, err := getCharacterByName(characterName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	for i, inventoryItem := range character.Inventaire {
		if inventoryItem == item {
			character.Inventaire = removeItem(character.Inventaire, i)
			return true
		}
	}
	return false
}

// import cycle not allowed
func removeItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// import cycle not allowed
func getCharacterByName(characterName string) (*Character, error) {
	name := strings.ToLower(strings.TrimSpace(characterName))
	if C2_b {
		switch name {
		case strings.ToLower(C1.Nom):
			return C1, nil
		case strings.ToLower(C2.Nom):
			return C2, nil
		default:
			return nil, fmt.Errorf("%spersonnage '%s' non trouvé%s", couleurs.Red, characterName, couleurs.Reset)
		}
	}
	return C1, nil
}

func GetEquipmentBonus(itemName string) int {
	itemLower := strings.ToLower(itemName)
	switch {
	// Chapeaux/Casques
	case strings.Contains(itemLower, "chapeau"):
		return 10
	case strings.Contains(itemLower, "casque"):
		return 5
	// Tuniques/Armures
	case strings.Contains(itemLower, "tunique"):
		return 25
	case strings.Contains(itemLower, "cuir bouilli"):
		return 15
	// Bottes
	case strings.Contains(itemLower, "bottes"):
		return 15
	// Les armes ne donnent pas de bonus PV
	case strings.Contains(itemLower, "épée") || strings.Contains(itemLower, "epee"):
		return 0
	case strings.Contains(itemLower, "dague"):
		return 0
	case strings.Contains(itemLower, "gourdin"):
		return 0
	case strings.Contains(itemLower, "arc"):
		return 0
	default:
		return 0
	}
}

/*
	func UnequipItem(character *Character, slotType string) bool {
		var item string
		switch strings.ToLower(slotType) {
		case "chapeau":
			if character.Equipment.Chapeau == "" {
				return false
			}
			item = character.Equipment.Chapeau
			character.Equipment.Chapeau = ""
		case "tunique":
			if character.Equipment.Tunique == "" {
				return false
			}
			item = character.Equipment.Tunique
			character.Equipment.Tunique = ""
		case "bottes":
			if character.Equipment.Bottes == "" {
				return false
			}
			item = character.Equipment.Bottes
			character.Equipment.Bottes = ""
		default:
			return false
		}
		bonusPV := GetEquipmentBonus(item)
		character.PointsDeVieMaximum -= bonusPV
		if character.PointsDeVieActuels > character.PointsDeVieMaximum {
			character.PointsDeVieActuels = character.PointsDeVieMaximum
		}
		character.Inventaire = append(character.Inventaire, item)
		fmt.Printf("\n%s%s déséquipe %s%s%s (-%d PV max)%s\n", couleurs.Yellow, character.Nom, couleurs.White, item, couleurs.Yellow, bonusPV, couleurs.Reset)
		return true
	}
*/
func UpgradeInventorySlot(character *Character) bool {
	if character.UpgradesUtilisés >= 3 {
		fmt.Printf("%s%s a déjà utilisé tous ses upgrades d'inventaire (3/3)%s\n", couleurs.Red, character.Nom, couleurs.Reset)
		return false
	}
	if character.PiècesDOr < 30 {
		manque := 30 - character.PiècesDOr
		fmt.Printf("%sVous n'avez pas assez de pièces d'or ! (il vous manque %d pièces)%s\n", couleurs.Red, manque, couleurs.Reset)
		return false
	}
	character.PiècesDOr -= 30
	character.InventaireMaxSlots += 10
	character.UpgradesUtilisés++
	fmt.Printf("\n%sInventaire de %s%s%s amélioré !%s\n",
		couleurs.Green, couleurs.White, character.Nom, couleurs.Green, couleurs.Reset)
	fmt.Printf("%sNouvelle capacité : %s%d slots%s (%d/3 upgrades utilisés)\n",
		couleurs.White, couleurs.Yellow, character.InventaireMaxSlots,
		couleurs.White, character.UpgradesUtilisés)
	fmt.Printf("%sNouveau solde : %s%d pièces d'or%s\n",
		couleurs.White, couleurs.Yellow, character.PiècesDOr, couleurs.Reset)
	return true
}

func DisplayEquipment(character *Character) {
	fmt.Printf("\n%s┌─────────────────────────────────────────────────┐%s\n", couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s│%s %sÉquipements de%s %s%-29s%s %s   │%s\n",
		couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
		couleurs.Green, character.Nom, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	fmt.Printf("%s├─────────────────────────────────────────────────┤%s\n", couleurs.Blue, couleurs.Reset)
	if character.Equipment.Chapeau != "" {
		fmt.Printf("%s│%s %sChapeau :%s %-36s %s │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			character.Equipment.Chapeau, couleurs.Blue, couleurs.Reset)
	} else {
		fmt.Printf("%s│%s %sChapeau :%s %sAucun%s                             %s    │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			couleurs.Red, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	if character.Equipment.Tunique != "" {
		fmt.Printf("%s│%s %sTunique :%s %-36s %s │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			character.Equipment.Tunique, couleurs.Blue, couleurs.Reset)
	} else {
		fmt.Printf("%s│%s %sTunique :%s %sAucune%s                            %s    │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			couleurs.Red, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	if character.Equipment.Bottes != "" {
		fmt.Printf("%s│%s %sBottes :%s %-37s %s │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			character.Equipment.Bottes, couleurs.Blue, couleurs.Reset)
	} else {
		fmt.Printf("%s│%s %sBottes :%s %sAucunes%s                            %s    │%s\n",
			couleurs.Blue, couleurs.Reset, couleurs.White, couleurs.Reset,
			couleurs.Red, couleurs.Reset, couleurs.Blue, couleurs.Reset)
	}
	fmt.Printf("%s└─────────────────────────────────────────────────┘%s\n", couleurs.Blue, couleurs.Reset)
}
