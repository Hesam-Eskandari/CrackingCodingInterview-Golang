package linkedlists

import (
	"container/list"
	"fmt"
	"testing"
)

func (l *linkedList) assertEqualArray(t *testing.T, expectedArr []int) {
	for index, item := range l.toList() {
		if item != expectedArr[index] {
			t.Errorf("returned linked list does not match to expected array, returned: %v, expected: %v",
				l.toList(), expectedArr)
			break
		}
	}
}

func (l *linkedList) assertNilElement(t *testing.T, element *list.Element) {
	if element != nil {
		t.Errorf("expected nil returned %v", element.Value)
	}
}

func (l *linkedList) assertElementValue(t *testing.T, element *list.Element, value int) {
	if element == nil {
		panic(fmt.Sprintf("assertElementValue: given element is nil, expected value: %v", value))
	}
	if element.Value != value {
		t.Errorf("expected element with value= %v received %v element", value, element.Value)
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

func setUpKthToLast() (input *linkedList, outValue []int, outValid []bool, k []int) {
	input = newLinkedList().createListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	k = []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	outValid = []bool{false, true, true, true, true, true, true, true, true, true, true, false}
	outValue = []int{0, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	return
}

func TestKthToEnd(t *testing.T) {
	input, outValue, outValid, k := setUpKthToLast()
	for index, value := range k {
		t.Run(fmt.Sprintf("k: %v", value), func(t *testing.T) {
			element := input.kthToLast(value)

			if !outValid[index] {
				input.assertNilElement(t, element)
			} else {
				input.assertElementValue(t, element, outValue[index])
			}
		})
	}
}
