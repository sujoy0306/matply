package matply

import "fmt"

// Matrix holds the matrix dimension and cell data
type Matrix struct {
	Row  int
	Col  int
	Data [][]float64
}

// NewMatrix creates a new matrix from 2D data source
func NewMatrix(d [][]float64) (*Matrix, error) {
	row := len(d)
	if row == 0 {
		return nil, fmt.Errorf("empty data")
	}
	col := len(d[0])
	data := make([][]float64, row)
	for i := 0; i < row; i++ {
		if len(d[i]) != col {
			return nil, fmt.Errorf("row length mismatch")
		}
		data[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			data[i][j] = d[i][j]
		}
	}
	return &Matrix{
		row,
		col,
		data,
	}, nil
}
