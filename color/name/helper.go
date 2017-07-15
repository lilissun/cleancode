package name

import "math"

const (
	precision = 6
)

func round(float float64) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(float*shift+.5) / shift
}
