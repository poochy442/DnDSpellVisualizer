package shared

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SkillDuration represents the duration of a spell.
type SkillDuration int

// School represents the school of magic.
type School int

// Range represents the range of a spell.
type Range int

// AreaOfEffect represents the area of effect of a spell.
type AreaOfEffect int

// DamageType represents the type of damage a spell deals.
type DamageType int

// DrawStyle represents the drawing style for visualization.
type DrawStyle string

// ColorStyle represents the color style for visualization.
type ColorStyle string

// VisualizationConfig holds configuration for spell visualization.
type VisualizationConfig struct {
	Size       int        `json:"Size"`
	DrawStyle  DrawStyle  `json:"DrawStyle"`
	ColorStyle ColorStyle `json:"ColorStyle"`
}

// Spell represents a D&D spell.
type Spell struct {
	Name         string
	Description  string
	RulesText    string
	HigherLevels string
	Level        int
	School       School
	Duration     SkillDuration
	Range        Range
	DamageType   DamageType
	AreaOfEffect AreaOfEffect
}

const (
	Instantaneous SkillDuration = 1
	Round         SkillDuration = 2
	Minute        SkillDuration = 3
	TenMinutes    SkillDuration = 4
	Hour          SkillDuration = 5
	EightHours    SkillDuration = 6
	DayPlus       SkillDuration = 7
	Permanent     SkillDuration = 8
)

func ParseSkillDuration(s string) (SkillDuration, error) {
	switch strings.ToLower(s) {
	case "instantaneous":
		return Instantaneous, nil
	case "1 round", "up to 1 round", "6 seconds":
		return Round, nil
	case "1 minute", "up to 1 minute":
		return Minute, nil
	case "10 minutes", "up to 10 minutes":
		return TenMinutes, nil
	case "1 hour", "up to 1 hour", "up to 2 hours":
		return Hour, nil
	case "8 hours", "up to 8 hours":
		return EightHours, nil
	case "24 hours", "up to 24 hours", "7 days", "up to 7 days", "10 days", "up to 10 days", "30 days", "up to 30 days":
		return DayPlus, nil
	case "until dispelled", "special":
		return Permanent, nil
	default:
		return 0, fmt.Errorf("invalid SkillDuration: %s", s)
	}
}

const (
	NoDamageType DamageType = 0
	Acid         DamageType = 1
	Bludgeoning  DamageType = 2
	Cold         DamageType = 3
	Fire         DamageType = 4
	Force        DamageType = 5
	Lightning    DamageType = 6
	Necrotic     DamageType = 7
	Piercing     DamageType = 8
	Poison       DamageType = 9
	Psychic      DamageType = 10
	Radiant      DamageType = 11
	Slashing     DamageType = 12
	Thunder      DamageType = 13
)

func ParseDamageType(s string) (DamageType, error) {
	switch strings.ToLower(s) {
	case "":
		return NoDamageType, nil
	case "acid":
		return Acid, nil
	case "bludgeoning":
		return Bludgeoning, nil
	case "cold":
		return Cold, nil
	case "fire":
		return Fire, nil
	case "force":
		return Force, nil
	case "lightning":
		return Lightning, nil
	case "necrotic":
		return Necrotic, nil
	case "piercing":
		return Piercing, nil
	case "poison":
		return Poison, nil
	case "psychic":
		return Psychic, nil
	case "radiant":
		return Radiant, nil
	case "slashing":
		return Slashing, nil
	case "thunder":
		return Thunder, nil
	default:
		return 0, fmt.Errorf("invalid DamageType: %s", s)
	}
}

const (
	Abjuration    School = 1
	Conjuration   School = 2
	Divination    School = 3
	Enchantment   School = 4
	Evocation     School = 5
	Illusion      School = 6
	Necromancy    School = 7
	Transmutation School = 8
)

func ParseSchool(s string) (School, error) {
	switch strings.ToLower(s) {
	case "abjuration":
		return Abjuration, nil
	case "conjuration":
		return Conjuration, nil
	case "divination":
		return Divination, nil
	case "enchantment":
		return Enchantment, nil
	case "evocation":
		return Evocation, nil
	case "illusion":
		return Illusion, nil
	case "necromancy":
		return Necromancy, nil
	case "transmutation":
		return Transmutation, nil
	default:
		return 0, fmt.Errorf("invalid School: %s", s)
	}
}

const (
	NoAreaOfEffect AreaOfEffect = 0
	Cone           AreaOfEffect = 1
	Cube           AreaOfEffect = 2
	Cylinder       AreaOfEffect = 3
	Line           AreaOfEffect = 4
	Sphere         AreaOfEffect = 5
)

func ParseAreaOfEffect(s string) (AreaOfEffect, error) {
	switch strings.ToLower(s) {
	case "":
		return NoAreaOfEffect, nil
	case "cone":
		return Cone, nil
	case "cube":
		return Cube, nil
	case "cylinder":
		return Cylinder, nil
	case "line":
		return Line, nil
	case "sphere":
		return Sphere, nil
	default:
		return 0, fmt.Errorf("invalid AreaOfEffect: %s", s)
	}
}

const (
	Self      Range = 1
	Touch     Range = 2
	Short     Range = 3
	Medium    Range = 4
	Long      Range = 5
	UltraLong Range = 6
	Unlimited Range = 7
	Special   Range = 8
)

func ParseRange(s string) (Range, error) {
	switch strings.ToLower(s) {
	case "self":
		return Self, nil
	case "touch":
		return Touch, nil
	case "5 feet", "10 feet", "30 feet":
		return Short, nil
	case "60 feet", "90 feet":
		return Medium, nil
	case "100 feet", "120 feet", "150 feet":
		return Long, nil
	case "300 feet", "500 feet", "1 mile", "500 miles":
		return UltraLong, nil
	case "unlimited":
		return Unlimited, nil
	case "special", "sight":
		return Special, nil
	default:
		return 0, fmt.Errorf("invalid Range: %s", s)
	}
}

const (
	StyleClassic DrawStyle = "Classic"
	StyleCurved  DrawStyle = "Curved"
	StyleLinear  DrawStyle = "Linear"

	ColorClassic ColorStyle = "Classic"
	ColorComplex ColorStyle = "Complex"
)

func (s *Spell) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		RulesText    string `json:"rulesText"`
		HigherLevels string `json:"higherLevels"`
		Level        int    `json:"level"`
		School       struct {
			Name string `json:"name"`
		}
		Duration string `json:"duration"`
		Range    string `json:"range"`
		Damage   struct {
			DamageType struct {
				Name string `json:"name"`
			} `json:"damage_type"`
		} `json:"damage"`
		AreaOfEffect struct {
			Type string `json:"type"`
		} `json:"area_of_effect"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	s.Name = aux.Name
	s.Description = aux.Description
	s.RulesText = aux.RulesText
	s.HigherLevels = aux.HigherLevels
	s.Level = aux.Level

	var err error
	s.School, err = ParseSchool(aux.School.Name)
	if err != nil {
		return err
	}

	s.Duration, err = ParseSkillDuration(aux.Duration)
	if err != nil {
		return err
	}

	s.Range, err = ParseRange(aux.Range)
	if err != nil {
		return err
	}

	s.DamageType, err = ParseDamageType(aux.Damage.DamageType.Name)
	if err != nil {
		return err
	}

	s.AreaOfEffect, err = ParseAreaOfEffect(aux.AreaOfEffect.Type)
	if err != nil {
		return err
	}

	return nil
}

func (vc *VisualizationConfig) UnmarshalJSON(data []byte) error {
	*vc = VisualizationConfig{
		Size:       0,
		DrawStyle:  StyleClassic,
		ColorStyle: ColorClassic,
	}

	type rawConfig VisualizationConfig
	return json.Unmarshal(data, (*rawConfig)(vc))
}
