package color

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
	hsl := rgb.ToHSL()
	if hsl.Equal(NewHSLColor(270, 0.8, 0.5)) == false {
		t.Errorf("convert rgb to hsl error")
	}
	newrgb := hsl.ToRGB()
	if rgb.Equal(newrgb) == false {
		t.Errorf("convert hsl to rgb error")
	}
}