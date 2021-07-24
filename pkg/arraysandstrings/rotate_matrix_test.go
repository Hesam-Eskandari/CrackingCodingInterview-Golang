package arraysandstrings

import (
	"testing"
)

type testRotateMatrix struct {
	message  string
	mat      *matrix
	expected *matrix
}

func newTestRotateMatrix(t *testing.T, message string, matrix1 [][]int, expected [][]int) *testRotateMatrix {
	mat1, err := NewMatrix(matrix1)
	if err != nil {
		t.Errorf("[NewTestRotateMatrix]: error creating test case, error: %s", err.Error())
	}
	mat2, err := NewMatrix(expected)
	if err != nil {
		t.Errorf("[NewTestRotateMatrix]: error creating test case, error: %s", err.Error())
	}
	return &testRotateMatrix{
		message:  message,
		mat:      mat1,
		expected: mat2,
	}
}

func (tm *testRotateMatrix) assertEqual(t *testing.T) {
	if len(tm.mat.mat) != len(tm.expected.mat) {
		t.Errorf("the size of the original matrix and the expected test case are not equal,"+
			"matrix length: %v by %v, expectex test matrix size: %v by %v",
			len(tm.mat.mat), len(tm.mat.mat), len(tm.expected.mat), len(tm.expected.mat))
	}
	for row := range tm.mat.mat {
		for col := range tm.mat.mat[row] {
			if tm.mat.mat[row][col] != tm.expected.mat[row][col] {
				t.Errorf("unexpected matrix returned. Expected: %v, returned: %v", tm.expected.mat, tm.mat.mat)
			}
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	inputs := []*testRotateMatrix{
		newTestRotateMatrix(t, "empty matrix", [][]int{}, [][]int{}),
		newTestRotateMatrix(t, "one element matrix", [][]int{{1}}, [][]int{{1}}),
		newTestRotateMatrix(t, "two by two matrix", [][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}),
		newTestRotateMatrix(t, "three by three", [][]int{{0, 1, 2}, {3, 0, 1}, {2, 4, 7}}, [][]int{{2, 3, 0}, {4, 0, 1}, {7, 1, 2}}),
	}
	for _, input := range inputs {
		t.Run(input.message, func(t *testing.T) {
			input.mat.rotateMatrix()
			input.assertEqual(t)
		})
	}
}
