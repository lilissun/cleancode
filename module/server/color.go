package main

import (
	"encoding/json"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// RGBColor defines colors in RGB
type RGBColor struct {
	Red   float64 `json:"r"`
	Green float64 `json:"g"`
	Blue  float64 `json:"b"`
}

// HSLColor defines colors in HSL
type HSLColor struct {
	Hue        float64 `json:"h"`
	Saturation float64 `json:"s"`
	Luminance  float64 `json:"l"`
}

const (
	serviceNameRGBtoHSL = "rgb2hsl"
)

func init() {
	register(router, serviceNameRGBtoHSL, fromRGBtoHSL)
}

func fromRGBtoHSL(req []byte) []byte {
	var rgb RGBColor
	err := json.Unmarshal(req, &rgb)
	if err != nil {
		return []byte(err.Error())
	}
	color := colorful.Color{R: rgb.Red, G: rgb.Green, B: rgb.Blue}
	var hsl HSLColor
	hsl.Hue, hsl.Saturation, hsl.Luminance = color.Hsl()
	resp, err := json.Marshal(hsl)
	if err != nil {
		return []byte(err.Error())
	}
	return resp
}
