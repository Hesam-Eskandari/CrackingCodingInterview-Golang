package linkedlists

import (
	"container/list"
)

type linkedList struct {
	list *list.List
}

func newLinkedList() *linkedList {
	return &linkedList{
		list: list.New(),
	}
}

func (l *linkedList) createListFromArray(arr []int) *linkedList {
	for _, item := range arr {
		l.list.PushBack(item)
	}
	return l
}

func (l *linkedList) toList() []int {
	var arr []int
	if l == nil {
		return arr
	}
	element := l.list.Front()
	for element != nil {
		arr = append(arr, element.Value.(int))
		element = element.Next()
	}
	return arr
}

// assume it's a singly linked list
func (l *linkedList) deleteDuplicates(noAdditionalDatastructures bool) {
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
		hmap := make(map[int]bool)
		hmap[element.Value.(int)] = true
		for element.Next() != nil {
			if _, ok := hmap[element.Next().Value.(int)]; ok {
				l.list.Remove(element.Next())
			} else {
				hmap[element.Next().Value.(int)] = true
				element = element.Next()
			}
		}
	}
}

func (l *linkedList) kthToLast(k int) *list.Element {
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
