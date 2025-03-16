package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTupleIsPoint(t *testing.T) {
	sut := Tuple{4.3, -4.2, 3.1, PointW}
	require.Equal(t, sut.X, 4.3)
	require.Equal(t, sut.Y, -4.2)
	require.Equal(t, sut.Z, 3.1)
	require.Equal(t, sut.W, 1.0)
	require.False(t, sut.IsVector())
	require.True(t, sut.IsPoint())
}

func TestTuple_IsVector(t *testing.T) {
	sut := Tuple{4.3, -4.2, 3.1, VectorW}
	require.Equal(t, sut.X, 4.3)
	require.Equal(t, sut.Y, -4.2)
	require.Equal(t, sut.Z, 3.1)
	require.Equal(t, sut.W, 0.0)
	require.True(t, sut.IsVector())
	require.False(t, sut.IsPoint())
}

func TestTuple_CreatePoint(t *testing.T) {
	sut := NewPoint(4, -4, 3)
	require.Equal(t, Tuple{4, -4, 3, 1.0}, sut)
	require.True(t, sut.IsPoint())
}

func TestTuple_CreateVector(t *testing.T) {
	sut := NewVector(4, -4, 3)
	require.Equal(t, Tuple{4, -4, 3, 0.0}, sut)
	require.True(t, sut.IsVector())
}

func TestTuple_Equality(t *testing.T) {
	point1 := NewPoint(4, -4, 3)
	point2 := NewPoint(4, -4, 3)
	require.True(t, point1.Equals(point2))
}

func TestTuple_Addition(t *testing.T) {
	point := NewPoint(3, -2, 5)
	vector := NewVector(-2, 3, 1)
	sut := point.Add(vector)
	require.Equal(t, Tuple{1, 1, 6, PointW}, sut)
}

func TestTuple_SubtractionPoint(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)
	sut := point1.Sub(point2)
	require.Equal(t, Tuple{-2, -4, -6, VectorW}, sut)
}

func TestTuple_Subtraction_Vector(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)
	sut := point.Sub(vector)
	require.Equal(t, Tuple{-2, -4, -6, PointW}, sut)
}

func TestTuple_Subtraction_TwoVectors(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)
	sut := vector1.Sub(vector2)
	require.Equal(t, Tuple{-2, -4, -6, VectorW}, sut)
}

func TestTuple_Negation(t *testing.T) {
	vector := NewVector(1, -2, 3)
	sut := vector.Negate()
	require.Equal(t, Tuple{-1, 2, -3, VectorW}, sut)
}

func TestTuple_NegationWithW(t *testing.T) {
	vector := Tuple{1, -2, 3, -4}
	sut := vector.Negate()
	require.Equal(t, Tuple{-1, 2, -3, 4}, sut)
}

func TestTuple_MultiplyByScalar(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	sut := tuple.Scale(3.5)
	require.Equal(t, Tuple{3.5, -7, 10.5, -14}, sut)
}

func TestTuple_MultiplyByFraction(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	sut := tuple.Scale(0.5)
	require.Equal(t, Tuple{0.5, -1, 1.5, -2}, sut)
}

func TestTuple_Division(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}
	sut := tuple.Divide(2)
	require.Equal(t, Tuple{0.5, -1, 1.5, -2}, sut)
}

func TestTuple_Magnitude(t *testing.T) {
	tests := []struct {
		input    Tuple
		expected float64
	}{
		{
			NewVector(1, 0, 0),
			1,
		},
		{
			NewVector(0, 1, 0),
			1,
		},
		{
			NewVector(0, 0, 1),
			1,
		},
		{
			NewVector(1, 2, 3),
			math.Sqrt(14),
		},
		{
			NewVector(-1, -2, -3),
			math.Sqrt(14),
		},
	}

	for _, test := range tests {
		sut := test.input.Magnitude()
		require.InDelta(t, test.expected, sut, 1e-5)
	}
}

func TestTuple_NormalizeVector400(t *testing.T) {
	vector := NewVector(4, 0, 0)
	sut := vector.Normalize()
	require.Equal(t, Tuple{1, 0, 0, VectorW}, sut)
}

func TestTuple_NormalizeVector123(t *testing.T) {
	vector := NewVector(1, 2, 3)
	sut := vector.Normalize()
	require.Equal(t, Tuple{
		1 / math.Sqrt(14),
		2 / math.Sqrt(14),
		3 / math.Sqrt(14),
		VectorW},
		sut)
}

func TestTuple_MagnitudeOfNormalizedVector(t *testing.T) {
	vector := NewVector(1, 2, 3)
	norm := vector.Normalize()
	sut := norm.Magnitude()
	require.Equal(t, 1.0, sut)
}

func TestTuple_DotProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	sut := vector1.Dot(vector2)

	require.Equal(t, 20.0, sut)
}

func TestTuple_CrossProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	sut := vector1.Cross(vector2)

	require.Equal(t, Tuple{-1, 2, -1, VectorW}, sut)

	sut = vector2.Cross(vector1)

	require.Equal(t, Tuple{1, -2, 1, VectorW}, sut)
}

func TestTuple_Color(t *testing.T) {
	color := NewColor(-0.5, 0.4, 1.7)
	require.Equal(t, Tuple{-0.5, 0.4, 1.7, ColorW}, color)
}

func TestTuple_MultiplyColor(t *testing.T) {
	color1 := NewColor(1, 0.2, 0.4)
	color2 := NewColor(0.9, 1, 0.1)

	sut := color1.Hadamard(color2)

	require.InDelta(t, 0.9, sut.X, 1e-5)
	require.InDelta(t, 0.2, sut.Y, 1e-5)
	require.InDelta(t, 0.04, sut.Z, 1e-5)
	require.InDelta(t, ColorW, sut.W, 1e-5)
}
