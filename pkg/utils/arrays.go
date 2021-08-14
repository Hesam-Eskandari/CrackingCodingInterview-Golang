package utils

import (
	"fmt"
	"github.com/Data-Structures-Golang/pkg/linkedlists"
)

func ArrayRemoveDuplicates(array []interface{}) ([]interface{}, error) {
	if array == nil {
		return nil, fmt.Errorf("RemoveFromValue: expected an array, received nil")
	}
	list := linkedlists.NewLinkedList().CreateListFromArray(array)
	list.DeleteDuplicatesBack(false)
	return list.ToArray(), nil
}

func ArrayRemoveDuplicatesBack(array []interface{}) ([]interface{}, error) {
	if array == nil {
		return nil, fmt.Errorf("RemoveFromValue: expected an array, received nil")
	}
	list := linkedlists.NewLinkedList().CreateListFromArray(array)
	list.DeleteDuplicatesBack(false)
	return list.ToArray(), nil
}
