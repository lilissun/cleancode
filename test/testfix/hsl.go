package testfix

import colorful "github.com/lucasb-eyer/go-colorful"

// HSLColor defines colors in HSL
type HSLColor struct {
	Hue        float64
	Saturation float64
	Luminance  float64
}

// NewHSLColor constructs HSLColor
func NewHSLColor(hue float64, saturation float64, luminance float64) HSLColor {
	return HSLColor{Hue: hue, Saturation: saturation, Luminance: luminance}
}

// ToRGB converts color from HSL space tp RGB space
func (hsl HSLColor) ToRGB() RGBColor {
	color := colorful.Hsl(hsl.Hue, hsl.Saturation, hsl.Luminance)
	var rgb RGBColor
	rgb.Red, rgb.Green, rgb.Blue = color.R, color.G, color.B
	return rgb
}

// Equal tests if rgb equals to x
func (hsl HSLColor) Equal(x HSLColor) bool {
	return floatEqual(hsl.Hue, x.Hue) &&
		floatEqual(hsl.Saturation, x.Saturation) &&
		floatEqual(hsl.Luminance, x.Luminance)
}
