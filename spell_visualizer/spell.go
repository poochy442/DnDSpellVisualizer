package spell_visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
)

// Spell is an alias for shared.Spell.
type Spell = shared.Spell

// LoadSpells loads all spells and returns a map by name.
func LoadSpells() (map[string]Spell, error) {
	spellList, err := GetAllSpells()
	if err != nil {
		return nil, err
	}
	spells := make(map[string]Spell)
	for _, spell := range spellList {
		spells[spell.Name] = spell
	}
	return spells, nil
}
