package spell_visualizer

import (
	"log"
)

type SkillDuration string
type DamageType string
type School string

const (
	Instantaneous SkillDuration = "Instantaneous"
	Fire          DamageType    = "Fire"
	Evocation     School        = "Evocation"
)

type Spell struct {
	Name         string
	Description  string
	RulesText    string
	HigherLevels string
	Level        int
	School       School
	Duration     SkillDuration
	Range        int
	DamageType   DamageType
	AreaOfEffect bool
}

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
