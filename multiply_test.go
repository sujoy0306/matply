package matply_test

import (
	"testing"

	mat "github.com/smkch95/matply"
	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	testcases := []struct {
		m1       *mat.Matrix
		m2       *mat.Matrix
		expected *mat.Matrix
	}{
		{
			m1: &mat.Matrix{
				Row:  2,
				Col:  1,
				Data: [][]float64{{1}, {1}},
			},
			m2: &mat.Matrix{
				Row:  2,
				Col:  2,
				Data: [][]float64{{1, 1}, {1, 1}},
			},
			expected: nil,
		},
		{
			m1: &mat.Matrix{
				Row:  2,
				Col:  2,
				Data: [][]float64{{1, 1}, {1, 1}},
			},
			m2: &mat.Matrix{
				Row:  2,
				Col:  2,
				Data: [][]float64{{1, 1}, {1, 1}},
			},
			expected: &mat.Matrix{
				Row:  2,
				Col:  2,
				Data: [][]float64{{2, 2}, {2, 2}},
			},
		},
	}

	assert := assert.New(t)

	for _, tc := range testcases {
		m, e := mat.Multiply(tc.m1, tc.m2)
		if e == nil {
			assert.True(areMatricesEqual(tc.expected, m))
		} else {
			assert.Nil(m)
		}
	}
}
