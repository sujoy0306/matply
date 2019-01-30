package matply_test

import (
	"strings"
	"testing"

	mat "github.com/smkch95/matply"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	testcases := []struct {
		data     string
		format   mat.Format
		expected *mat.Matrix
	}{
		{
			data:     "",
			format:   "wrong_format",
			expected: nil,
		},
		{
			data: `1.2,3.4
5.6, "7.8"`,
			format:   mat.Csv,
			expected: nil,
		},
		{
			data: `1.2,3.4
5.6,abcd`,
			format:   mat.Csv,
			expected: nil,
		},
		{
			data: `1.2,3.4
5.6,7.8`,
			format: mat.Csv,
			expected: &mat.Matrix{
				Row: 2,
				Col: 2,
				Data: [][]float64{
					{1.2, 3.4},
					{5.6, 7.8},
				},
			},
		},
	}

	assert := assert.New(t)

	for _, tc := range testcases {
		r := strings.NewReader(tc.data)
		m, e := mat.Read(r, tc.format)
		if e == nil {
			assert.True(areMatricesEqual(tc.expected, m))
		} else {
			assert.Nil(m)
		}
	}
}
