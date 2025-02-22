package geometry

import (
	"fmt"
	"math"
)

const (
	epsilon = 1e-5
	VectorW = 0.0
	PointW  = 1.0
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, PointW}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, VectorW}
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func (t Tuple) IsPoint() bool {
	return equal(t.W, PointW)
}

func (t Tuple) IsVector() bool {
	return equal(t.W, VectorW)
}

func (t Tuple) Add(u Tuple) Tuple {
	return Tuple{t.X + u.X, t.Y + u.Y, t.Z + u.Z, t.W + u.W}
}

func (t Tuple) Sub(u Tuple) Tuple {
	return Tuple{t.X - u.X, t.Y - u.Y, t.Z - u.Z, t.W - u.W}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) Scale(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t Tuple) Divide(s float64) Tuple {
	if s == 0 {
		panic("divide by zero")
	}
	return Tuple{t.X / s, t.Y / s, t.Z / s, t.W / s}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(
		math.Pow(t.X, 2) +
			math.Pow(t.Y, 2) +
			math.Pow(t.Z, 2) +
			math.Pow(t.W, 2))
}

func (t Tuple) Normalize() Tuple {
	return Tuple{
		t.X / t.Magnitude(),
		t.Y / t.Magnitude(),
		t.Z / t.Magnitude(),
		t.W / t.Magnitude(),
	}
}

func (t Tuple) Dot(u Tuple) float64 {
	return t.X*u.X + t.Y*u.Y + t.Z*u.Z + t.W*u.W
}

func (t Tuple) Cross(u Tuple) Tuple {
	return NewVector(
		t.Y*u.Z-t.Z*u.Y,
		t.Z*u.X-t.X*u.Z,
		t.X*u.Y-t.Y*u.X,
	)
}

func (t Tuple) Equals(u Tuple) bool {
	return equal(t.X, u.X) &&
		equal(t.Y, u.Y) &&
		equal(t.Z, u.Z) &&
		equal(t.W, u.W)
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func (t Tuple) String() string {
	return fmt.Sprintf("X: %.2f, Y: %.2f, Z: %.2f, W: %.2f", t.X, t.Y, t.Z, t.W)
}
