package matply_test

import (
	"testing"

	mat "github.com/smkch95/matply"
	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	testcases := []struct {
		data     [][]float64
		expected *mat.Matrix
	}{
		{nil, nil},
		{[][]float64{}, nil},
		{[][]float64{{1, 2}, {3}}, nil},
		{[][]float64{{1, 2}, {3, 4}}, &mat.Matrix{
			Row:  2,
			Col:  2,
			Data: [][]float64{{1, 2}, {3, 4}},
		}},
	}

	assert := assert.New(t)

	for _, tc := range testcases {
		m, e := mat.NewMatrix(tc.data)
		if e == nil {
			assert.True(areMatricesEqual(tc.expected, m))
		} else {
			assert.Nil(m)
		}
	}
}

func areMatricesEqual(m1, m2 *mat.Matrix) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if (m1 == nil && m2 != nil) || (m1 != nil && m2 == nil) {
		return false
	}
	if m1.Col != m2.Col || m1.Row != m2.Row {
		return false
	}
	for i := 0; i < m1.Row; i++ {
		for j := 0; j < m1.Col; j++ {
			if m1.Data[i][j] != m2.Data[i][j] {
				return false
			}
		}
	}
	return true
}
