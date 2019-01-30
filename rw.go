package matply

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Format is the matrix representation formats
type Format string

const (
	// Csv represents comma separated value format
	Csv Format = "csv"
)

// Read reads a matrix from r
func Read(r io.Reader, format Format) (*Matrix, error) {
	switch format {
	case Csv:
		return ReadCsv(r)
	default:
		return nil, fmt.Errorf("Format %v not supported", format)
	}
}

// ReadCsv reads a matrix from csv formatted data
func ReadCsv(r io.Reader) (*Matrix, error) {
	csv := csv.NewReader(r)
	raw, err := csv.ReadAll()
	if err != nil {
		return nil, err
	}
	data, err := convert(raw)
	if err != nil {
		return nil, err
	}
	return NewMatrix(data)
}

func convert(raw [][]string) ([][]float64, error) {
	row := len(raw)
	res := make([][]float64, row)
	for i := 0; i < row; i++ {
		col := len(raw[i])
		res[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			d, err := stof(raw[i][j])
			if err != nil {
				return nil, err
			}
			res[i][j] = d
		}
	}
	return res, nil
}

func stof(raw string) (float64, error) {
	trimmed := strings.Trim(raw, " ")
	return strconv.ParseFloat(trimmed, 64)
}
