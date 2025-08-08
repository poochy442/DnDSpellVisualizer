package spell_visualizer

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Color represents a color in HEX, RGB, and HSV formats.
type Color struct {
	HEX string
	RGB RGB
	HSV HSV
}

type RGB struct {
	R, G, B int
}

type HSV struct {
	H, S, V float64
}

// NewColor creates a Color from a HEX string.
func NewColor(hex string) (*Color, error) {
	r, g, b, err := hexToRGB(hex)
	if err != nil {
		return nil, err
	}
	h, s, v := rgbToHSV(r, g, b)
	return &Color{
		HEX: hex,
		RGB: RGB{R: r, G: g, B: b},
		HSV: HSV{H: h, S: s, V: v},
	}, nil
}

// NewColorFromRGB creates a Color from RGB values.
func NewColorFromRGB(r, g, b int) *Color {
	h, s, v := rgbToHSV(r, g, b)
	return &Color{
		RGB: RGB{R: r, G: g, B: b},
		HSV: HSV{H: h, S: s, V: v},
		HEX: rgbToHex(r, g, b),
	}
}

// NewColorFromHSV creates a Color from HSV values.
func NewColorFromHSV(h, s, v float64) *Color {
	r, g, b := hsvToRGB(h, s, v)
	return &Color{
		HSV: HSV{H: h, S: s, V: v},
		RGB: RGB{R: r, G: g, B: b},
		HEX: rgbToHex(r, g, b),
	}
}

func hexToRGB(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid HEX color: %s", hex)
	}
	r, err1 := strconv.ParseInt(hex[0:2], 16, 64)
	g, err2 := strconv.ParseInt(hex[2:4], 16, 64)
	b, err3 := strconv.ParseInt(hex[4:6], 16, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid HEX color: %s", hex)
	}
	return int(r), int(g), int(b), nil
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func rgbToHSV(r, g, b int) (float64, float64, float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0
	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	delta := max - min

	var h float64
	if delta == 0 {
		h = 0
	} else if max == rf {
		h = math.Mod((gf-bf)/delta, 6)
	} else if max == gf {
		h = ((bf-rf)/delta + 2)
	} else {
		h = ((rf-gf)/delta + 4)
	}
	h = h * 60
	if h < 0 {
		h += 360
	}

	s := 0.0
	if max != 0 {
		s = delta / max
	}
	v := max
	return h, s, v
}

func hsvToRGB(h, s, v float64) (int, int, int) {
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := v - c
	var rf, gf, bf float64

	switch {
	case h >= 0 && h < 60:
		rf, gf, bf = c, x, 0
	case h >= 60 && h < 120:
		rf, gf, bf = x, c, 0
	case h >= 120 && h < 180:
		rf, gf, bf = 0, c, x
	case h >= 180 && h < 240:
		rf, gf, bf = 0, x, c
	case h >= 240 && h < 300:
		rf, gf, bf = x, 0, c
	case h >= 300 && h < 360:
		rf, gf, bf = c, 0, x
	default:
		rf, gf, bf = 0, 0, 0
	}
	r := int(math.Round((rf + m) * 255))
	g := int(math.Round((gf + m) * 255))
	b := int(math.Round((bf + m) * 255))
	return r, g, b
}

// RotateHSV rotates the color's hue by the given angle.
func (c *Color) RotateHSV(angle float64) *Color {
	h := math.Mod(c.HSV.H+angle, 360)
	if h < 0 {
		h += 360
	}
	return NewColorFromHSV(h, c.HSV.S, c.HSV.V)
}

func (c *Color) String() string {
	if c.HEX != "" {
		return c.HEX
	}
	return rgbToHex(c.RGB.R, c.RGB.G, c.RGB.B)
}
