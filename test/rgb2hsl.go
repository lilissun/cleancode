package test

import (
	colorful "github.com/lucasb-eyer/go-colorful"
)

// RGBColor defines colors in RGB
type RGBColor struct {
	Red   float64
	Green float64
	Blue  float64
}

// NewRGBColor constructs RGBColor
func NewRGBColor(red float64, green float64, blue float64) RGBColor {
	return RGBColor{Red: red, Green: green, Blue: blue}
}

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

// ConvertRGBtoHSL converts color in RGB space to HSL space
func ConvertRGBtoHSL(rgb RGBColor) HSLColor {
	color := colorful.Color{R: rgb.Red, G: rgb.Green, B: rgb.Blue}
	var hsl HSLColor
	hsl.Hue, hsl.Saturation, hsl.Luminance = color.Hsl()
	return hsl
}
