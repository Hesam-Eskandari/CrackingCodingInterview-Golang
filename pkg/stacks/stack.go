package Stacks

import "reflect"

type (
	node struct {
		value interface{}
		prev  *node
	}
	Stack struct {
		top    *node
		length int
	}
)

func newStack() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

func (s *Stack) Top() interface{} {
	return s.top.value
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Push(value interface{}) interface{} {
	top := &node{
		value: value,
		prev:  s.top,
	}
	s.top = top
	s.length += 1
	return s.top.value
}

func (s *Stack) Pop() (detachedHead interface{}) {
	if s.top == nil {
		panic("cannot remove the top node of an empty stack")
	}
	detachedHead = s.top.value
	if s.top.prev == nil {
		s.top = nil
	} else {
		s.top = s.top.prev
	}
	s.length -= 1
	return
}

func (s *Stack) ToArray() (arr []interface{}) {
	arr = make([]interface{}, s.length)
	node := s.top
	for index := s.length - 1; index >= 0; index-- {
		arr[index] = node.value
		node = node.prev
	}
	return
}

func (s *Stack) AppendArray(array interface{}) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			s.Push(arr.Index(index).Interface())
		}
	}
}

func (s *Stack) Append(stack *Stack) {
	if s == nil {
		panic("error appending two stacks, the primary stack cannot be nil")
	}
	if stack == nil {
		return
	}
	for node := stack.top; node != nil; {
		s.Push(node.value)
		node = node.prev
	}
}
