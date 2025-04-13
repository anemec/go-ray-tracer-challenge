package geometry

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_New4x4Matrix(t *testing.T) {
	matrixValues := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	sut := NewMatrix(matrixValues)

	for r, row := range matrixValues {
		for c, value := range row {
			require.Equal(t, sut.At(r, c), value)
		}
	}
}

func Test_New3x3Matrix(t *testing.T) {
	matrixValues := [][]float64{
		{1, 2, 3},
		{5.5, 6.5, 7.5},
		{9, 10, 11},
		{13.5, 14.5, 15.5},
	}
	sut := NewMatrix(matrixValues)

	for r, row := range matrixValues {
		for c, value := range row {
			require.Equal(t, sut.At(r, c), value)
		}
	}
}

func Test_New2x2Matrix(t *testing.T) {
	matrixValues := [][]float64{
		{1, 2},
		{5.5, 6.5},
		{9, 10},
		{13.5, 14.5},
	}
	sut := NewMatrix(matrixValues)

	for r, row := range matrixValues {
		for c, value := range row {
			require.Equal(t, sut.At(r, c), value)
		}
	}
}

func Test_MatrixEqual(t *testing.T) {
	matrixValues1 := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	matrixValues2 := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	matrixOne := NewMatrix(matrixValues1)
	matrixTwo := NewMatrix(matrixValues2)

	require.True(t, matrixOne.Equal(matrixTwo))
}

func Test_MatrixNotEqual(t *testing.T) {
	matrixValues1 := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	matrixValues2 := [][]float64{
		{1, 2, 3, 1},
		{5.5, 6.5, 7.4, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	matrixOne := NewMatrix(matrixValues1)
	matrixTwo := NewMatrix(matrixValues2)

	require.False(t, matrixOne.Equal(matrixTwo))
}

func Test_MatrixMultiply(t *testing.T) {
	matrixValues1 := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	matrixValues2 := [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}

	matrixOne := NewMatrix(matrixValues1)
	matrixTwo := NewMatrix(matrixValues2)

	expectedValues := [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}

	expectedMatrix := NewMatrix(expectedValues)

	sut := matrixOne.Multiply(matrixTwo)

	require.True(t, expectedMatrix.Equal(sut))
}

func Test_MatrixMultiplyTuple(t *testing.T) {
	matrixValues := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	matrix := NewMatrix(matrixValues)
	tuple := Tuple{
		1, 2, 3, 1,
	}
	expected := Tuple{
		18, 24, 33, 1,
	}

	sut := matrix.MultiplyTuple(tuple)

	require.Equal(t, expected, sut)
}

func Test_IdentityMatrix(t *testing.T) {
	matrixValues := [][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}
	matrix := NewMatrix(matrixValues)
	sut := matrix.Multiply(Identity)

	require.Equal(t, matrix, sut)
}

func Test_TransposeMatrix(t *testing.T) {
	matrixValues := [][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}
	matrix := NewMatrix(matrixValues)

	expectedValues := [][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}
	expectedMatrix := NewMatrix(expectedValues)

	sut := matrix.Transpose()

	require.Equal(t, expectedMatrix, sut)
}

func Test_TransposeIdentityMatrix(t *testing.T) {
	identity := NewIdentity()
	sut := identity.Transpose()
	require.Equal(t, Identity, sut)
}
