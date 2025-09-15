package savegame

import (
	"ASCII_Aventure/characters"
	"encoding/json"
	"os"
)

func SaveCharacter(c characters.Character) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("character.json", data, 0644)
}

func LoadCharacter() (characters.Character, error) {
	var c characters.Character
	data, err := os.ReadFile("character.json")
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	return c, err
}
