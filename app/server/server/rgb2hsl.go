package main

import (
	"encoding/json"

	"bitbucket.org/lilissun/cleancode/color/color"
)

// Processor Name
const (
	RGB2HSLName = "rgb2hsl"
)

func init() {
	register(RGB2HSLName, rgb2hsl)
}

func rgb2hsl(req []byte) ([]byte, error) {
	var rgb color.RGBColor
	err := json.Unmarshal([]byte(req), &rgb)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(rgb.ToHSL())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
