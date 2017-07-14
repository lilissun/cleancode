package namefix

import "math"

const (
	epsilon = 0.00000001
)

func floatEqual(f1 float64, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}
