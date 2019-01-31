package matply

import (
	"fmt"
)

// Multiply takes two matrices and checks if their dimensions are compatible, and if
// so returns the matrix which is resulted from multiplication of those two
func Multiply(m1 *Matrix, m2 *Matrix) (*Matrix, error) {
	if !areMatricesMultipliable(m1, m2) {
		return nil, fmt.Errorf("Dimensions not valid")
	}
	l := m1.Col
	m := &Matrix{
		Row:  m1.Row,
		Col:  m2.Col,
		Data: nil,
	}
	m.Data = make([][]float64, m.Row)
	for i := 0; i < m.Row; i++ {
		m.Data[i] = make([]float64, m.Col)
		for j := 0; j < m.Col; j++ {
			m.Data[i][j] = 0
			for k := 0; k < l; k++ {
				m.Data[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}
	return m, nil
}

func areMatricesMultipliable(m1, m2 *Matrix) bool {
	return m1.Col == m2.Row
}
