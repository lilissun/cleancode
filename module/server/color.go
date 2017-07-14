package main

import (
	"encoding/json"

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

const (
	serviceNameRGBtoHSL = "rgb2hsl"
)

func init() {
	register(router, serviceNameRGBtoHSL, fromRGBtoHSL)
}

func fromRGBtoHSL(req string) string {
	var rgb RGBColor
	err := json.Unmarshal([]byte(req), &rgb)
	if err != nil {
		return err.Error()
	}
	color := colorful.Color{R: rgb.Red, G: rgb.Green, B: rgb.Blue}
	var hsl HSLColor
	hsl.Hue, hsl.Saturation, hsl.Luminance = color.Hsl()
	resp, err := json.Marshal(hsl)
	if err != nil {
		return err.Error()
	}
	return string(resp)
}
