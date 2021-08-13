package linkedlists

import (
	"container/list"
	"fmt"
	"testing"
)

// AssertEqualArray if a values and ordering of a linked list is equivalent to an array
func (l *linkedList) AssertEqualArray(t *testing.T, expectedArr []int) {
	for index, item := range l.ToArray() {
		if item != expectedArr[index] {
			t.Errorf("returned linked list does not match to expected array, returned: %v, expected: %v",
				l.ToArray(), expectedArr)
			break
		}
	}
}

// AssertNilElement checks if an element is nil
func (l *linkedList) AssertNilElement(t *testing.T, element *list.Element) {
	if element != nil {
		t.Errorf("expected nil returned %v", element.Value)
	}
}

// AssertEqualElementValue checks if an element value is equal to what is expected
func (l *linkedList) AssertEqualElementValue(t *testing.T, element *list.Element, value int) {
	if element == nil {
		panic(fmt.Sprintf("AssertEqualElementValue: given element is nil, expected value: %v", value))
	}
	if element.Value != value {
		t.Errorf("expected element with value= %v received %v element", value, element.Value)
	}
}

func setUpDeleteDuplicates() (inputs []*linkedList, out [][]int) {
	inputs = []*linkedList{
		NewLinkedList().CreateListFromArray([]int{}),
		NewLinkedList().CreateListFromArray([]int{1}),
		NewLinkedList().CreateListFromArray([]int{1, 2, 4, 5, 3, 4, 5}),
		NewLinkedList().CreateListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
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
			input.DeleteDuplicates(true)
			input.AssertEqualArray(t, out[i])
		})
	}
}

func TestDeleteDuplicatesUsingAdditionalDatastructures(t *testing.T) {
	inputs, out := setUpDeleteDuplicates()
	for i, input := range inputs {
		t.Run(fmt.Sprintf("{%v, %v}", i, input), func(t *testing.T) {
			input.DeleteDuplicates(false)
			input.AssertEqualArray(t, out[i])
		})
	}
}

func setUpKthToLast() (input *linkedList, outValue []int, outValid []bool, k []int) {
	input = NewLinkedList().CreateListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	k = []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	outValid = []bool{false, true, true, true, true, true, true, true, true, true, true, false}
	outValue = []int{0, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	return
}

func TestKthToEnd(t *testing.T) {
	input, outValue, outValid, k := setUpKthToLast()
	for index, value := range k {
		t.Run(fmt.Sprintf("k: %v", value), func(t *testing.T) {
			element := input.KthToLast(value)

			if !outValid[index] {
				input.AssertNilElement(t, element)
			} else {
				input.AssertEqualElementValue(t, element, outValue[index])
			}
		})
	}
}

func TestLinkedList_Replace(t *testing.T) {
	newList := NewLinkedList().CreateListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	newList.Replace(NewLinkedList().CreateListFromArray(arr))
	newList.AssertEqualArray(t, arr)
}
