package gonet

// Matrix represents a two dimensional slice of size (rows, cols)
type Matrix struct {
	rows, cols int
	data       [][]float64
}

// Create matrix and fill with 0's
func Create(rows, cols int) *Matrix {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = 0
		}
	}
	return &Matrix{rows: rows, cols: cols, data: matrix}
}

// Read returns the 2D slice in the Matrix struct
func (m Matrix) Read() [][]float64 {
	return m.data
}

// GetSize returns array of number of rows and number of cols
func (m Matrix) GetSize() [2]int {
	return [2]int{m.rows, m.cols}
}

//////////////////////
// Scalar Functions //
//////////////////////

// Add sums every element with a float64 parameter
func (m *Matrix) Add(n float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i][j] += n
		}
	}
}

// Multiply multiplies every element with a float64 parameter
func (m *Matrix) Multiply(n float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i][j] *= n
		}
	}
}
