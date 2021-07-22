package arraysandstrings

import "testing"

type testZeroMatrix struct {
	message    string
	zeroMatrix *zeroMatrix
	expected   *zeroMatrix
}

func (tzm *testZeroMatrix) assertTrue(t *testing.T) {
	for row := range tzm.zeroMatrix.matrix {
		for col := range tzm.zeroMatrix.matrix[row] {
			if tzm.zeroMatrix.matrix[row][col] != tzm.expected.matrix[row][col] {
				t.Errorf("expected: %v, returned: %v", tzm.expected, tzm.zeroMatrix)
			}
		}
	}
}

func TestSetToZero(t *testing.T) {
	inputs := []*testZeroMatrix{
		{"two by two", NewZeroMatrix([][]int{{2, 0}, {0, 2}}), NewZeroMatrix([][]int{{0, 0}, {0, 0}})},
		{"three by three 1", NewZeroMatrix([][]int{{1, 1, 2}, {2, 0, 1}, {1, 1, 2}}), NewZeroMatrix([][]int{{1, 0, 2}, {0, 0, 0}, {1, 0, 2}})},
		{"three by three 2", NewZeroMatrix([][]int{{1, 0, 2}, {1, 0, 2}, {1, 0, 2}}), NewZeroMatrix([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})},
	}
	for _, input := range inputs {
		t.Run(input.message, func(t *testing.T) {
			input.zeroMatrix.setToZero()
			input.assertTrue(t)
		})
	}
}
