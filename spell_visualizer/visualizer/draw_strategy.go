package visualizer

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
	"fmt"
	"math"
)

type DrawStrategy interface {
	Draw(spell *shared.Spell, color string, size int) (content string, err error)
}

type ClassicDrawStrategy struct{}

func (d ClassicDrawStrategy) Draw(spell *shared.Spell, color string, size int) (string, error) {
	const numPoints = 13
	const centerX = 100
	const centerY = 100

	points := ""
	lines := ""

	pointLocations := make([][2]float64, numPoints)
	for i := range numPoints {
		angle := 2 * math.Pi * float64(i) / float64(numPoints)
		x := centerX + float64(size)*math.Sin(angle)
		y := centerY - float64(size)*math.Cos(angle)
		pointLocations[i] = [2]float64{x, y}
	}

	features := shared.Features

	// Level
	feature := features[spell.Level]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+1)%numPoints][0], pointLocations[(i+1)%numPoints][1], color)
		}
	}
	// School
	feature = features[spell.School]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+2)%numPoints][0], pointLocations[(i+2)%numPoints][1], color)
		}
	}
	// Damage Type
	feature = features[spell.DamageType]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+3)%numPoints][0], pointLocations[(i+3)%numPoints][1], color)
		}
	}
	// Area of Effect
	feature = features[spell.AreaOfEffect]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+4)%numPoints][0], pointLocations[(i+4)%numPoints][1], color)
		}
	}
	// Range
	feature = features[spell.Range]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+5)%numPoints][0], pointLocations[(i+5)%numPoints][1], color)
		}
	}
	// Duration
	feature = features[spell.Duration]
	for i := range pointLocations {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+6)%numPoints][0], pointLocations[(i+6)%numPoints][1], color)
		}
	}

	for i := range numPoints {
		x, y := pointLocations[i][0], pointLocations[i][1]
		strokeColor := color
		fillColor := color
		if i == 0 {
			strokeColor = color
			fillColor = "white"
		}
		points += fmt.Sprintf(`<circle cx="%f" cy="%f" r="5" stroke="%s" stroke-width="2" fill="%s" />`, x, y, strokeColor, fillColor)
	}

	return lines + points, nil
}
