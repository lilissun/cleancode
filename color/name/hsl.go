package name

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
	return NewRGBColor(round(color.R), round(color.G), round(color.B))
}

// Equal tests if rgb equals to x
func (hsl HSLColor) Equal(x HSLColor) bool {
	return hsl.Hue == x.Hue &&
		hsl.Saturation == x.Saturation &&
		hsl.Luminance == x.Luminance
}
