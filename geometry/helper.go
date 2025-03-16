package geometry

import "math"

const (
	epsilon = 1e-5
)

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
