package shared

import (
	"fmt"
	"math"
	"math/rand"
)

// Features for spell visualization
var Features = []string{
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

var DamageTypeColors = map[DamageType]struct {
	Main      string
	Secondary string
	Tertiary  string
}{
	NoDamageType: {"#000000", "#808080", "#333333"},
	Acid:         {"#075800", "#064d00", "#054200"},
	Bludgeoning:  {"#808080", "#707070", "#606060"},
	Cold:         {"#3fd0d4", "#37b6ba", "#2f9c9f"},
	Fire:         {"#ff5a00", "#df4f00", "#bf4400"},
	Force:        {"#ba38b1", "#a3319b", "#8b2a85"},
	Lightning:    {"#021CA4", "#021890", "#02157b"},
	Necrotic:     {"#738022", "#65701e", "#56601a"},
	Piercing:     {"#808080", "#707070", "#606060"},
	Poison:       {"#4aa31d", "#418f19", "#387a16"},
	Psychic:      {"#7a07b0", "#6b069a", "#5c0584"},
	Radiant:      {"#f5ef42", "#d6d13a", "#b8b332"},
	Slashing:     {"#808080", "#707070", "#606060"},
	Thunder:      {"#563bcc", "#4b34b3", "#412c99"},
}

func GenerateChaoticPattern(colors []string) (string, error) {
	mainColor := colors[0]
	secondaryColor := colors[1]
	tertiaryColor := colors[2]

	minFirstLayer := 15
	maxFirstLayer := 20
	minSecondLayer := 7
	maxSecondLayer := 13

	start := `<pattern id="chaoticPattern" patternUnits="userSpaceOnUse" width="100" height="100">`
	content := fmt.Sprintf(`<rect width="100" height="100" fill="%s" />`, mainColor)

	firstLayerShapes := []struct {
		cx, cy, radius float64
	}{}

	numFirstLayerShapes := rand.Intn(maxFirstLayer-minFirstLayer+1) + minFirstLayer
	firstLayerBaseRadius := 30.0 * (float64(minFirstLayer) / float64(numFirstLayerShapes))
	for i := range numFirstLayerShapes {
		angle := 2 * math.Pi * float64(i) / float64(numFirstLayerShapes)
		dist := rand.Float64()*48 + 15
		cx := 50 + dist*math.Cos(angle)
		cy := 50 + dist*math.Sin(angle)
		radius := rand.Float64()*firstLayerBaseRadius*0.5 + firstLayerBaseRadius*0.5
		bumpiness := rand.Intn(7) + 13
		// secColor, err := NewColor(secondaryColor)
		// if err != nil {
		// 	return "", err
		// }
		// rotation := rand.Float64()*2*maxFirstLayerRotation - maxFirstLayerRotation
		// secColorRot := secColor.RotateHSV(rotation)
		firstLayerShapes = append(firstLayerShapes, struct {
			cx, cy, radius float64
		}{cx, cy, radius})
		content += generateSmoothShape(cx, cy, radius, bumpiness, secondaryColor)
	}

	for _, parent := range firstLayerShapes {
		numSecondLayerShapes := rand.Intn(maxSecondLayer-minSecondLayer+1) + minSecondLayer
		secondLayerBaseRadius := parent.radius * (float64(minSecondLayer) / float64(numSecondLayerShapes)) * 0.42
		for j := range numSecondLayerShapes {
			angle := 2 * math.Pi * float64(j) / float64(numSecondLayerShapes)
			dist := rand.Float64() * parent.radius * 0.9
			cx := parent.cx + dist*math.Cos(angle)
			cy := parent.cy + dist*math.Sin(angle)
			radius := rand.Float64()*secondLayerBaseRadius*0.5 + secondLayerBaseRadius*0.5
			bumpiness := rand.Intn(7) + 13
			// terColor, err := NewColor(tertiaryColor)
			// if err != nil {
			// 	return "", err
			// }
			// rotation := rand.Float64()*2*maxSecondLayerRotation - maxSecondLayerRotation
			// terColorRot := terColor.RotateHSV(rotation)
			content += generateSmoothShape(cx, cy, radius, bumpiness, tertiaryColor)
		}
	}

	end := `</pattern>`
	return start + content + end, nil
}

func generateSmoothShape(cx, cy, radius float64, bumpiness int, color string) string {
	points := make([][2]float64, bumpiness)
	for i := range bumpiness {
		angle := 2 * math.Pi * float64(i) / float64(bumpiness)
		bumpRadius := radius + rand.Float64()*radius*0.18 - radius*0.09
		x := cx + bumpRadius*math.Cos(angle)
		y := cy + bumpRadius*math.Sin(angle)
		points[i] = [2]float64{x, y}
	}

	path := fmt.Sprintf("M %f,%f ", points[0][0], points[0][1])
	for i := range bumpiness {
		p0 := points[(i-1+bumpiness)%bumpiness]
		p1 := points[i]
		p2 := points[(i+1)%bumpiness]
		p3 := points[(i+2)%bumpiness]
		c1x := p1[0] + (p2[0]-p0[0])/6
		c1y := p1[1] + (p2[1]-p0[1])/6
		c2x := p2[0] - (p3[0]-p1[0])/6
		c2y := p2[1] - (p3[1]-p1[1])/6
		path += fmt.Sprintf("C %f,%f %f,%f %f,%f ", c1x, c1y, c2x, c2y, p2[0], p2[1])
	}
	path += "Z"
	return fmt.Sprintf(`<path d="%s" fill="%s" />`, path, color)
}
