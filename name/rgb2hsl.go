package name

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// Convert converts color in RGB space to HSL space
func Convert(c []float64) ([]float64, error) {
	if len(c) != 3 {
		return nil, fmt.Errorf("color dimension=[%d] error", len(c))
	}
	c[0], c[1], c[2] = colorful.Color{R: c[0], G: c[1], B: c[2]}.Hsl()
	return c, nil
}
