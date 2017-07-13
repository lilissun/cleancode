package main

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func main() {
	c := []float64{0.1, 0.5, 0.3}
	fmt.Printf("RGB Color: %v\n", c)
	c[0], c[1], c[2] = colorful.Color{c[0], c[1], c[2]}.Hsl()
	fmt.Printf("HSL Color: %v\n", c)
}
