package spell_visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
)

type Spell = shared.Spell

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
