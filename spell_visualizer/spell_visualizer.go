package spell_visualizer

import "fmt"

func VisualizeSpell(spell *Spell) string {
	return fmt.Sprintf(`
        <svg xmlns="http://www.w3.org/2000/svg" width="200" height="200">
            <circle cx="100" cy="100" r="80" stroke="black" stroke-width="3" fill="red" />
            <text x="100" y="105" font-size="20" text-anchor="middle" fill="white">%s</text>
        </svg>`, spell.Name)
}
