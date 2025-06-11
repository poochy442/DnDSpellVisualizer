package spell_visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
	"log"
)

type Spell = shared.Spell

var spells map[string]Spell

func init() {
	spellList, err := GetAllSpells()
	if err != nil {
		log.Fatalf("Failed to load spells: %v", err)
	}
	spells = make(map[string]Spell)
	for _, spell := range spellList {
		spells[spell.Name] = spell
	}
}

func GetSpellByName(name string) *Spell {
	if spell, exists := spells[name]; exists {
		return &spell
	}
	return nil
}
