package gonet

// Matrix represents a two dimensional slice of size (rows, cols)
type Matrix struct {
	rows, cols int
	data       [][]int
}

// Create matrix and fill with 0's
func Create(rows, cols int) *Matrix {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = 0
		}
	}
	return &Matrix{rows: rows, cols: cols, data: matrix}
}

func (m Matrix) Read() [][]int {
	return m.data
}
