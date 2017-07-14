package test

import "testing"

func TestConvertRGBtoHSL(t *testing.T) {
	rgb := NewRGBColor(0.5, 0.1, 0.9)
	hsl := ConvertRGBtoHSL(rgb)
	if hsl.Hue != 270 || hsl.Saturation != 0.8 || hsl.Luminance != 0.5 {
		t.Errorf("convert rgb to hsl error")
	}
}
