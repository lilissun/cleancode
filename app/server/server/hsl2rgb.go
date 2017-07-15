package main

import (
	"encoding/json"

	"bitbucket.org/lilissun/cleancode/test/color"
)

// Processor Name
const (
	HSL2RGBName = "hsl2rgb"
)

func init() {
	register(HSL2RGBName, hsl2rgb)
}

func hsl2rgb(req []byte) ([]byte, error) {
	var hsl color.HSLColor
	err := json.Unmarshal([]byte(req), &hsl)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(hsl.ToRGB())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
