package savegame

import (
	"ASCII_Aventure/characters"
	"encoding/json"
	"os"
)

func SaveCharacter(character *characters.Character) error {
	jsonData, err := json.MarshalIndent(character, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile("jsonData.json", jsonData, 0644)
}

func LoadCharacter() (characters.Character, error) {
	var character characters.Character
	jsonData, err := os.ReadFile("jsonData.json")
	if err != nil {
		return character, err
	}
	err = json.Unmarshal(jsonData, &character)
	return character, err
}
