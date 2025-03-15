package spell_visualizer

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetAllSpells() ([]Spell, error) {
	file, err := os.Open("spells.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open spells.json: %v", err)
	}
	defer file.Close()

	var spells []Spell
	if err := json.NewDecoder(file).Decode(&spells); err != nil {
		return nil, fmt.Errorf("failed to decode spells.json: %v", err)
	}

	return spells, nil
}
