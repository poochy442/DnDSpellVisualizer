package visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
)

type ColorStrategy interface {
	GetDefs(spell *shared.Spell) (defs string, color string, err error)
}

type ClassicColorStrategy struct{}

func (c ClassicColorStrategy) GetDefs(spell *shared.Spell) (string, string, error) {
	colors := shared.DamageTypeColors[spell.DamageType]
	return "", colors.Main, nil
}

type ComplexColorStrategy struct{}

func (c ComplexColorStrategy) GetDefs(spell *shared.Spell) (string, string, error) {
	colors := shared.DamageTypeColors[spell.DamageType]
	pattern, err := shared.GenerateChaoticPattern([]string{colors.Main, colors.Secondary, colors.Tertiary})
	if err != nil {
		return "", "", err
	}
	defs := "<defs>" + pattern + "</defs>"
	return defs, "url(#chaoticPattern)", nil
}
