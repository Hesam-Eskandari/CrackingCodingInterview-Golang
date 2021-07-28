package linkedlists

import (
	"fmt"
	"testing"
)

func (l *linkedList) assertEqualArray(t *testing.T, expectedArr []int) {
	for index, item := range l.toList() {
		if item != expectedArr[index] {
			t.Errorf("terurned linked list does not match to expected array, returned: %v, expected: %v",
				l.toList(), expectedArr)
			break
		}
	}
}

func setUpDeleteDuplicates() (inputs []*linkedList, out [][]int) {
	inputs = []*linkedList{
		newLinkedList().createListFromArray([]int{}),
		newLinkedList().createListFromArray([]int{1}),
		newLinkedList().createListFromArray([]int{1, 2, 4, 5, 3, 4, 5}),
		newLinkedList().createListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
	}
	out = [][]int{
		{},
		{1},
		{1, 2, 4, 5, 3},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	return
}

func TestDeleteDuplicatesNoAdditionalDatastructures(t *testing.T) {
	inputs, out := setUpDeleteDuplicates()
	for i, input := range inputs {
		t.Run(fmt.Sprintf("{%v, %v}", i, input), func(t *testing.T) {
			input.deleteDuplicates(true)
			input.assertEqualArray(t, out[i])
		})
	}
}

func TestDeleteDuplicatesUsingAdditionalDatastructures(t *testing.T) {
	inputs, out := setUpDeleteDuplicates()
	for i, input := range inputs {
		t.Run(fmt.Sprintf("{%v, %v}", i, input), func(t *testing.T) {
			input.deleteDuplicates(false)
			input.assertEqualArray(t, out[i])
		})
	}
}
