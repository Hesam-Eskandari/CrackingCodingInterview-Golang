package arraysandstrings

import "fmt"

/*
Cracking the coding interview, 6th edition page 91

Rotate Matrix
*/

type matrix struct {
	mat [][]int
}

func NewMatrix(mat [][]int) (*matrix, error) {
	if length := len(mat); length != 0 {
		if row := mat[0]; len(row) != length {
			err := fmt.Errorf("[NewMatrix] error creating a new matrix, expected square matrix, received: %v", mat)
			return nil, err
		}
	}
	return &matrix{mat}, nil
}

func (m *matrix) rotateMatrix() {

	for start := 0; start < len(m.mat)/2; start++ {
		end := len(m.mat) - start - 1
		if end <= start {
			break
		}
		for index := range m.mat[start][start:end] {
			top := m.mat[start][start+index]                    // store top row element
			m.mat[start][start+index] = m.mat[end-index][start] // element(top row) <- element(left col)
			m.mat[end-index][start] = m.mat[end][end-index]     // element(left col) <- element(bottom row)
			m.mat[end][end-index] = m.mat[start+index][end]     // element(bottom row) <- element(right col)
			m.mat[start+index][end] = top                       // element(right col) <- element(top row)
		}
	}
}
