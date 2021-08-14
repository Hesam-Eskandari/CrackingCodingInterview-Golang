package linkedlists

import (
	"fmt"
	"testing"
)

func setUpDeleteDuplicates() (inputs []*linkedList, out [][]interface{}) {
	inputs = []*linkedList{
		NewLinkedList().CreateListFromArray([]interface{}{}),
		NewLinkedList().CreateListFromArray([]interface{}{1}),
		NewLinkedList().CreateListFromArray([]interface{}{1, 2, 4, 5, 3, 4, 5}),
		NewLinkedList().CreateListFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}),
	}
	out = [][]interface{}{
		{},
		{1},
		{1, 2, 4, 5, 3},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	return
}

func setUpDeleteDuplicatesBack() (inputs []*linkedList, out [][]interface{}) {
	inputs = []*linkedList{
		NewLinkedList().CreateListFromArray([]interface{}{}),
		NewLinkedList().CreateListFromArray([]interface{}{1}),
		NewLinkedList().CreateListFromArray([]interface{}{1, 2, 4, 5, 3, 4, 5}),
		NewLinkedList().CreateListFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}),
	}
	out = [][]interface{}{
		{},
		{1},
		{1, 2, 3, 4, 5},
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

func TestDeleteDuplicatesBackNoAdditionalDatastructures(t *testing.T) {
	inputs, out := setUpDeleteDuplicatesBack()
	for i, input := range inputs {
		t.Run(fmt.Sprintf("{%v, %v}", i, input), func(t *testing.T) {
			input.DeleteDuplicatesBack(true)
			input.AssertEqualArray(t, out[i])
		})
	}
}

func TestDeleteDuplicatesBackUsingAdditionalDatastructures(t *testing.T) {
	inputs, out := setUpDeleteDuplicatesBack()
	for i, input := range inputs {
		t.Run(fmt.Sprintf("{%v, %v}", i, input), func(t *testing.T) {
			input.DeleteDuplicatesBack(false)
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
	arr := []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}
	newList.Replace(NewLinkedList().CreateListFromArray(arr))
	newList.AssertEqualArray(t, arr)
}
