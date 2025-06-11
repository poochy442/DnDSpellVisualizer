package visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
	"fmt"
)

func VisualizeSpell(spell *shared.Spell, config shared.VisualizationConfig) (string, error) {
	size := 80
	if config.Size > 0 {
		size = config.Size
	}

	var colorStrategy ColorStrategy
	switch config.ColorStyle {
	case shared.ColorComplex:
		colorStrategy = ComplexColorStrategy{}
	default:
		colorStrategy = ClassicColorStrategy{}
	}

	var drawStrategy DrawStrategy
	switch config.DrawStyle {
	case "":
	case shared.StyleClassic:
		drawStrategy = ClassicDrawStrategy{}
	default:
		return "", fmt.Errorf("unsupported draw style: %s", config.DrawStyle)
	}

	defs, color, err := colorStrategy.GetDefs(spell)
	if err != nil {
		return "", err
	}

	content, err := drawStrategy.Draw(spell, color, size)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 220" preserveAspectRatio="xMidYMid meet">
            %s
            %s
            <text x="100" y="210" font-size="20" text-anchor="middle" fill="%s">%s</text>
        </svg>`, defs, content, color, spell.Name), nil
}
