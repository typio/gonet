package gonet

import (
	"fmt"
	"math/rand"
)

// Matrix represents a two dimensional slice of size (rows, cols)
type Matrix struct {
	rows, cols int
	data       [][]float64
}

///////////////////
// Basic Methods //
///////////////////

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

// PrintM prints the matrix
func (m Matrix) PrintM() {
	fmt.Printf("\nRows: %d, Cols: %d\n", m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Print(m.data[i][j])
			if j < m.cols-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// GetSize returns array of number of rows and number of cols
func (m Matrix) GetSize() [2]int {
	return [2]int{m.rows, m.cols}
}

// Randomize replaces every value in matrix with random number from range [0.0, 1.0)
func (m *Matrix) Randomize() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i][j] += rand.Float64()
		}
	}
}

////////////////////
// Scalar Methods //
////////////////////

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

/////////////////////////
// Elementwise Methods //
/////////////////////////

// AddM sums every element with corresponding element of anouther matrix
func (m *Matrix) AddM(n *Matrix) {
	if m.rows != n.rows || m.cols != n.cols {
		panic("Matricies are not of the same dimensions")
	} else {
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				m.data[i][j] += n.data[i][j]
			}
		}
	}
}

// MultiplyM multiplies every element with corresponding element of anouther matrix
func (m *Matrix) MultiplyM(n *Matrix) {
	if m.rows != n.rows || m.cols != n.cols {
		panic("Matricies are not of the same dimensions")
	} else {
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				m.data[i][j] *= n.data[i][j]
			}
		}
	}
}

// MatrixP returns a new matrix that is the matrix product
func (m *Matrix) MatrixP(n *Matrix) *Matrix {
	if m.rows != n.cols || m.cols != n.rows {
		panic("Matricies are not of the same dimensions, must be [rows, cols] = [cols, rows]")
	} else {
		result := Create(m.rows, n.cols)
		for i := 0; i < result.rows; i++ {
			for j := 0; j < result.cols; j++ {
				var sum float64
				for k := 0; k < m.cols; k++ {
					sum += m.data[i][k] * n.data[k][j]
				}
				result.data[i][j] = sum
			}
		}
		return result
	}
}

// Transpose returns a new matrix with rows and cols swapped
func (m *Matrix) Transpose() *Matrix {
	result := Create(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}
	return result
}
