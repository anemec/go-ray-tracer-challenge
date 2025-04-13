package geometry

type Matrix struct {
	m [][]float64
}

var Identity = NewMatrix([][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}})

func NewIdentity() *Matrix {
	return NewMatrix([][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}})
}

func NewMatrix(matrixValues [][]float64) *Matrix {
	return &Matrix{
		m: create(matrixValues),
	}
}

func create(matrixValues [][]float64) [][]float64 {
	matrix := make([][]float64, len(matrixValues))
	for r, row := range matrixValues {
		matrix[r] = append([]float64(nil), row...)
	}
	return matrix
}

func (m *Matrix) At(row, col int) float64 {
	return m.m[row][col]
}

func (m *Matrix) Multiply(matrix *Matrix) *Matrix {
	matrixValues := make([][]float64, len(matrix.m))
	for rIdx, row := range m.m {
		newRow := make([]float64, len(row))
		for cIdx := range row {
			newRow[cIdx] =
				m.At(rIdx, 0)*matrix.At(0, cIdx) +
					m.At(rIdx, 1)*matrix.At(1, cIdx) +
					m.At(rIdx, 2)*matrix.At(2, cIdx) +
					m.At(rIdx, 3)*matrix.At(3, cIdx)
		}
		matrixValues[rIdx] = newRow
	}
	return NewMatrix(matrixValues)
}

func (m *Matrix) MultiplyTuple(tuple Tuple) Tuple {
	newTupleValues := [4]float64{}
	for rIdx := range m.m {
		newTupleValues[rIdx] =
			m.At(rIdx, 0)*tuple.X +
				m.At(rIdx, 1)*tuple.Y +
				m.At(rIdx, 2)*tuple.Z +
				m.At(rIdx, 3)*tuple.W
	}
	return Tuple{
		newTupleValues[0],
		newTupleValues[1],
		newTupleValues[2],
		newTupleValues[3],
	}
}

func (m *Matrix) Equal(m2 *Matrix) bool {
	if len(m.m) != len(m2.m) {
		return false
	}
	for i, row := range m.m {
		if len(row) != len(m2.m[i]) {
			return false
		}

		for j, val := range row {
			if !equal(m2.m[i][j], val) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Transpose() *Matrix {
	newMatrixValues := make([][]float64, len(m.m))
	for i := 0; i < len(m.m); i++ {
		newRow := make([]float64, len(m.m))
		newMatrixValues[i] = newRow
	}

	for i, v := range m.m {
		for j, val := range v {
			newMatrixValues[j][i] = val
		}
	}
	return NewMatrix(newMatrixValues)
}
