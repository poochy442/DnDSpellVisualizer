package spell_visualizer

import (
	"fmt"
	"math"
)

var features = []string{
	"0000000000000",
	"0000000000001",
	"0000000000011",
	"0000000000101",
	"0000000000111",
	"0000000010001",
	"0000000010011",
	"0000000010101",
	"0000000010111",
	"0000000100001",
	"0000000101001",
	"0000000101111",
	"0000000110001",
	"0000000110111",
	"0000000111011",
	"0000000111111",
	"0000001000001",
}

var damageTypeColors = map[DamageType]string{
	NoDamageType: "black",
	Acid:         "#075800",
	Bludgeoning:  "gray",
	Cold:         "#3fd0d4",
	Fire:         "#ff5a00",
	Force:        "#ba38b1",
	Lightning:    "#021CA4",
	Necrotic:     "#738022",
	Piercing:     "gray",
	Poison:       "#4aa31d",
	Psychic:      "#7a07b0",
	Radiant:      "#f5ef42",
	Slashing:     "gray",
	Thunder:      "#563bcc",
}

func VisualizeSpell(spell *Spell) string {
	const numPoints = 13
	const radius = 80
	const centerX = 100
	const centerY = 100

	damageColor := damageTypeColors[spell.DamageType]

	points := ""
	lines := ""

	pointLocations := make([][2]float64, numPoints)
	for i := range numPoints {
		angle := 2 * math.Pi * float64(i) / float64(numPoints)
		x := centerX + radius*math.Sin(angle)
		y := centerY - radius*math.Cos(angle)
		pointLocations[i] = [2]float64{x, y}
	}

	// Level
	feature := features[spell.Level]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+1)%numPoints][0], pointLocations[(i+1)%numPoints][1], damageColor)
		}
	}

	// School
	feature = features[spell.School]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+2)%numPoints][0], pointLocations[(i+2)%numPoints][1], damageColor)
		}
	}

	// Damage Type
	feature = features[spell.DamageType]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+3)%numPoints][0], pointLocations[(i+3)%numPoints][1], damageColor)
		}
	}

	// Area of Effect
	feature = features[spell.AreaOfEffect]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+4)%numPoints][0], pointLocations[(i+4)%numPoints][1], damageColor)
		}
	}

	// Range
	feature = features[spell.Range]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+5)%numPoints][0], pointLocations[(i+5)%numPoints][1], damageColor)
		}
	}

	// Duration
	feature = features[spell.Duration]
	for i := range len(pointLocations) {
		if feature[i] == '1' {
			lines += fmt.Sprintf(`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="%s" stroke-width="1" />`, pointLocations[i][0], pointLocations[i][1], pointLocations[(i+6)%numPoints][0], pointLocations[(i+6)%numPoints][1], damageColor)
		}
	}

	for i := range numPoints {
		x, y := pointLocations[i][0], pointLocations[i][1]
		strokeColor := damageColor
		fillColor := damageColor
		if i == 0 {
			strokeColor = damageColor
			fillColor = "white"
		}
		points += fmt.Sprintf(`<circle cx="%f" cy="%f" r="5" stroke="%s" stroke-width="2" fill="%s" />`, x, y, strokeColor, fillColor)
	}

	return fmt.Sprintf(`
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 220" preserveAspectRatio="xMidYMid meet">
            %s
            %s
            <text x="100" y="210" font-size="20" text-anchor="middle" fill="black">%s</text>
        </svg>`, lines, points, spell.Name)
}
