package spell_visualizer

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetAllSpells loads all spells from spells.json.
func GetAllSpells() ([]Spell, error) {
	file, err := os.Open("spells.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open spells.json: %w", err)
	}
	defer file.Close()

	var spells []Spell
	if err := json.NewDecoder(file).Decode(&spells); err != nil {
		return nil, fmt.Errorf("failed to decode spells.json: %w", err)
	}
	return spells, nil
}
