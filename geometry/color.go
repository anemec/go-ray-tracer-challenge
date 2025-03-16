package geometry

import (
	"fmt"
	"math"
)

func NewColor(x, y, z float64) Tuple {
	return Tuple{x, y, z, ColorW}
}

func (t Tuple) Hadamard(u Tuple) Tuple {
	return Tuple{
		t.X * u.X,
		t.Y * u.Y,
		t.Z * u.Z,
		t.W * u.W,
	}
}

func (t Tuple) ToColorString() string {
	r := convertToPixel(t.X)
	g := convertToPixel(t.Y)
	b := convertToPixel(t.Z)
	return fmt.Sprintf("%d %d %d", r, g, b)
}

func convertToPixel(value float64) int {
	if value < 0 {
		return 0
	}
	if value >= 1 {
		return 255
	}

	return int(math.Round(value * 255))
}
