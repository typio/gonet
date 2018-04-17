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
	arr := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			arr[i][j] = 0
		}
	}
	return &Matrix{rows, cols, arr}
}

// FromArray creates a matrix from an array
func FromArray(arr []float64) *Matrix {
	matrix := Create(len(arr), 1)
	for i := 0; i < len(arr); i++ {
		matrix.data[i][0] = arr[i]
	}
	return matrix
}

// Read returns the 2D slice in the Matrix struct
func (m Matrix) Read() [][]float64 {
	return m.data
}

// PrintM prints the matrix
func (m Matrix) PrintM() {
	fmt.Printf("Rows: %d, Cols: %d\n", m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Print(m.data[i][j])
			if j < m.cols-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("\n")
	}
}

// GetSize returns array of number of rows and number of cols
func (m Matrix) GetSize() [2]int {
	return [2]int{m.rows, m.cols}
}

// Randomize replaces every value in matrix with random number from range [0.0, 1.0)
func (m *Matrix) Randomize() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i][j] += rand.Float64()*2 - 1
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

// AddM returns matrix of every element plus the corresponding element of anouther matrix
func (m *Matrix) AddM(n *Matrix) *Matrix {
	result := Create(m.rows, m.cols)
	if m.rows != n.rows || m.cols != n.cols {
		panic("Matricies are not of the same dimensions")
	} else {
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				result.data[i][j] = m.data[i][j] + n.data[i][j]
			}
		}
	}
	return result
}

// SubtractM returns matrix of every element minus the corresponding element of anouther matrix
func (m *Matrix) SubtractM(n *Matrix) *Matrix {
	result := Create(m.rows, m.cols)
	if m.rows != n.rows || m.cols != n.cols {
		panic("Matricies are not of the same dimensions")
	} else {
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				result.data[i][j] = m.data[i][j] - n.data[i][j]
			}
		}
	}
	return result
}

// MultiplyM returns matrix of every element times the corresponding element of anouther matrix
func (m *Matrix) MultiplyM(n *Matrix) *Matrix {
	result := Create(m.rows, m.cols)
	if m.rows != n.rows || m.cols != n.cols {
		panic("Matricies are not of the same dimensions")
	} else {
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				result.data[i][j] = m.data[i][j] + n.data[i][j]
			}
		}
	}
	return result
}

// MapM applies a function to every element in matrix
func (m *Matrix) MapM(fn func(n float64) float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i][j] = fn(m.data[i][j])
		}
	}
}

// MapNM returns new matrix with function applied to every element in matrix
func (m *Matrix) MapNM(fn func(n float64) float64) *Matrix {
	result := Create(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = fn(m.data[i][j])
		}
	}
	return result
}

// MatrixP returns a new matrix that is the matrix product
func (m *Matrix) MatrixP(n *Matrix) *Matrix {
	if m.cols != n.rows {
		panic("Matricies are not of correct dimensions for matrix multiplication")
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
