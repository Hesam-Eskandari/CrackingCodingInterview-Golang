package linkedlists

import (
	"container/list"
	"reflect"
)

type linkedList struct {
	list *list.List
}

type LinkedList interface {
	// CreateListFromArray constructs a doubly linkedList from a given array
	CreateListFromArray(array interface{}) *linkedList
	// DeleteDuplicates removes redundant nodes assuming it's a singly linked list with only head is given
	DeleteDuplicates(noAdditionalDatastructures bool)
	// KthToLast returns the kth element from tail (back) assuming it's a singly linked list with given head
	KthToLast(k int) *list.Element
	// Replace replaces the list inside the linkedList with a new list
	Replace(list *linkedList)
	// ToArray returns an array with values in linked list with same order and size
	ToArray() []interface{}
}

func NewLinkedList() LinkedList {
	return &linkedList{
		list: list.New(),
	}
}

// CreateListFromArray constructs a doubly linkedList from a given array
func (l *linkedList) CreateListFromArray(array interface{}) *linkedList {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			l.list.PushBack(arr.Index(index).Interface())
		}
	}
	return l
}

// ToArray returns an array with values in linked list with same order and size
func (l *linkedList) ToArray() []interface{} {
	var arr []interface{}
	if l == nil {
		return arr
	}
	element := l.list.Front()
	for element != nil {
		arr = append(arr, element.Value)
		element = element.Next()
	}
	return arr
}

// Replace replaces the list inside the linkedList with a new list
func (l *linkedList) Replace(list *linkedList) {
	l.list = list.list
}

// DeleteDuplicates removes redundant nodes assuming it's a singly linked list with only head is given
func (l *linkedList) DeleteDuplicates(noAdditionalDatastructures bool) {
	element := l.list.Front()
	if element == nil {
		return
	}
	if noAdditionalDatastructures {
		// O(n^2) time, O(1) space
		for element.Next() != nil {
			node := element
			for node.Next() != nil {
				if node.Next().Value == element.Value {
					l.list.Remove(node.Next())
				} else {
					node = node.Next()
				}
			}
			element = element.Next()
			if element == nil {
				break
			}
		}
	} else {
		// O(n) time, O(n) space
		hmap := make(map[interface{}]bool)
		hmap[element.Value] = true
		for element.Next() != nil {
			if _, ok := hmap[element.Next().Value]; ok {
				l.list.Remove(element.Next())
			} else {
				hmap[element.Next().Value] = true
				element = element.Next()
			}
		}
	}
}

// KthToLast returns the kth element from tail (back) assuming it's a singly linked list with given head
func (l *linkedList) KthToLast(k int) *list.Element {
	elementAhead := l.list.Front()
	elementDelay := l.list.Front()
	if elementAhead == nil || k < 0 {
		return nil
	}
	if k == 0 {
		k = 1
	}
	count := 0
	for elementAhead != nil {
		count += 1
		if count > k {
			elementDelay = elementDelay.Next()
		}
		elementAhead = elementAhead.Next()
	}
	if count < k {
		return nil
	}
	return elementDelay
}
