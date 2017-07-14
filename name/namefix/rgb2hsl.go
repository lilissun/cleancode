package namefix

import (
	colorful "github.com/lucasb-eyer/go-colorful"
)

// RGBColor defines colors in RGB
type RGBColor struct {
	Red   float64
	Green float64
	Blue  float64
}

// HSLColor defines colors in HSL
type HSLColor struct {
	Hue        float64
	Saturation float64
	Luminance  float64
}

// FromRGBtoHSL converts color in RGB space to HSL space
func FromRGBtoHSL(rgb RGBColor) HSLColor {
	color := colorful.Color{R: rgb.Red, G: rgb.Green, B: rgb.Blue}
	var hsl HSLColor
	hsl.Hue, hsl.Saturation, hsl.Luminance = color.Hsl()
	return hsl
}
