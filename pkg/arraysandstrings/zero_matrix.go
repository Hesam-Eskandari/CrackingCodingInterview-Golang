package arraysandstrings

type zeroMatrix struct {
	matrix [][]int
}

func NewZeroMatrix(matrix [][]int) *zeroMatrix {
	return &zeroMatrix{
		matrix: matrix,
	}
}

func (zm *zeroMatrix) setToZero() {
	zeroRow := false
	for _, value := range zm.matrix[0] {
		if value == 0 {
			zeroRow = true
		}
	}
	zeroCol := false
	for row := range zm.matrix {
		if zm.matrix[row][0] == 0 {
			zeroCol = true
		}
	}
	for row := range zm.matrix {
		for col := range zm.matrix[row] {
			if zm.matrix[row][col] == 0 {
				zm.matrix[0][col] = 0
				zm.matrix[row][0] = 0
			}
		}
	}
	for index, value := range zm.matrix[0] {
		if value == 0 {
			for row := range zm.matrix {
				zm.matrix[row][index] = 0
			}
		}
	}
	for index := range zm.matrix {
		if zm.matrix[index][0] == 0 {
			zm.matrix[index] = make([]int, len(zm.matrix[index]))
		}
	}
	if zeroRow {
		zm.matrix[0] = make([]int, len(zm.matrix[0]))
	}
	if zeroCol {
		for index := range zm.matrix {
			zm.matrix[index][0] = 0
		}
	}
}
