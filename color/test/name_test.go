package test

import (
	"testing"
)

// NewRGBColor(0.5, 0.1, 0.9)
// should have
// Hue == 270
// Saturation == 0.8
// Luminance == 0.5
func TestConvertRGBtoHSL(t *testing.T) {
	rgb := NewRGBColor(0.5, 0.1, 0.9)
	hsl := NewHSLColor(270, 0.8, 0.5)
	newhsl := rgb.ToHSL()
	if hsl.Equal(newhsl) == false {
		t.Errorf("convert rgb=[%v] to hsl=[%v] but expect=[%v]", rgb, newhsl, hsl)
	}
	newrgb := hsl.ToRGB()
	if rgb.Equal(newrgb) == false {
		t.Errorf("convert hsl=[%v] to rgb=[%v] but expect=[%v]", hsl, newrgb, rgb)
	}
}
