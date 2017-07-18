package main

import (
	"encoding/json"

	"github.com/lilissun/cleancode/color/color"
)

// Processor Name
const (
	HSL2RGBName = "hsl2rgb"
)

func init() {
	router[HSL2RGBName] = hsl2rgb
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
