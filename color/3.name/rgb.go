package name

import colorful "github.com/lucasb-eyer/go-colorful"

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

// ToHSL converts color from RGB space to HSL space
func (rgb RGBColor) ToHSL() HSLColor {
	color := colorful.Color{R: rgb.Red, G: rgb.Green, B: rgb.Blue}
	hue, saturation, luminance := color.Hsl()
	return NewHSLColor(hue, saturation, luminance)
}

// Equal tests if rgb equals to x
func (rgb RGBColor) Equal(x RGBColor) bool {
	return rgb.Red == x.Red &&
		rgb.Green == x.Green &&
		rgb.Blue == x.Blue
}
