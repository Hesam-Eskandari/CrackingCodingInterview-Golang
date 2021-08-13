package linkedlists

import (
	"container/list"
	"reflect"
	"testing"
)

type linkedList struct {
	List *list.List
}

type LinkedList interface {
	// AssertEqualArray if a values and ordering of a linked list is equivalent to an array
	AssertEqualArray(t *testing.T, expectedArr []int)
	// AssertEqualElementValue checks if an element value is equal to what is expected
	AssertEqualElementValue(t *testing.T, element *list.Element, value int)
	// AssertNilElement checks if an element is nil
	AssertNilElement(t *testing.T, element *list.Element)
	// CreateListFromArray constructs a doubly linkedList from a given array
	CreateListFromArray(array interface{}) *linkedList
	// DeleteDuplicates removes redundant nodes assuming it's a singly linked List with only head is given
	DeleteDuplicates(noAdditionalDatastructures bool)
	// GetList gets the inner list
	GetList() *list.List
	// KthToLast returns the kth element from tail (back) assuming it's a singly linked List with given head
	KthToLast(k int) *list.Element
	// Replace replaces the List inside the linkedList with a new List
	Replace(List *linkedList)
	// ToArray returns an array with values in linked List with same order and size
	ToArray() []interface{}
}

func NewLinkedList() LinkedList {
	return &linkedList{
		List: list.New(),
	}
}

// CreateListFromArray constructs a doubly linkedList from a given array
func (l *linkedList) CreateListFromArray(array interface{}) *linkedList {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			l.List.PushBack(arr.Index(index).Interface())
		}
	}
	return l
}

// ToArray returns an array with values in linked List with same order and size
func (l *linkedList) ToArray() []interface{} {
	var arr []interface{}
	if l == nil {
		return arr
	}
	element := l.List.Front()
	for element != nil {
		arr = append(arr, element.Value)
		element = element.Next()
	}
	return arr
}

// Replace replaces the List inside the linkedList with a new List
func (l *linkedList) Replace(List *linkedList) {
	l.List = List.List
}

// DeleteDuplicates removes redundant nodes assuming it's a singly linked List with only head is given
func (l *linkedList) DeleteDuplicates(noAdditionalDatastructures bool) {
	element := l.List.Front()
	if element == nil {
		return
	}
	if noAdditionalDatastructures {
		// O(n^2) time, O(1) space
		for element.Next() != nil {
			node := element
			for node.Next() != nil {
				if node.Next().Value == element.Value {
					l.List.Remove(node.Next())
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
				l.List.Remove(element.Next())
			} else {
				hmap[element.Next().Value] = true
				element = element.Next()
			}
		}
	}
}

// KthToLast returns the kth element from tail (back) assuming it's a singly linked List with given head
func (l *linkedList) KthToLast(k int) *list.Element {
	elementAhead := l.List.Front()
	elementDelay := l.List.Front()
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

// GetList gets the inner list
func (l *linkedList) GetList() *list.List {
	return l.List
}
