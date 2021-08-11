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

// Top returns the value of the top of the stack
func (s *Stack) Top() interface{} {
	return s.top.value
}

// Len returns the number of nodes in the stack
func (s *Stack) Len() int {
	return s.length
}

// Push adds a node with given value to the stack. Also moves the top to the new added node
func (s *Stack) Push(value interface{}) interface{} {
	top := &node{
		value: value,
		prev:  s.top,
	}
	s.top = top
	s.length += 1
	return s.top.value
}

// Pop removes the top node and returns its value. It also assign the previous node as the new top
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

// ToArray returns an array of values in the stack. The last index of array corresponds to the top of the stack
func (s *Stack) ToArray() (arr []interface{}) {
	arr = make([]interface{}, s.length)
	node := s.top
	for index := s.length - 1; index >= 0; index-- {
		arr[index] = node.value
		node = node.prev
	}
	return
}

// AppendArray pushes the item values of the array starting from index zero to the stack
func (s *Stack) AppendArray(array interface{}) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			s.Push(arr.Index(index).Interface())
		}
	}
}

// Reverse returns a new stack that has values of the original stack reversed
func (s *Stack) Reverse() *Stack {
	stack := newStack()
	for node := s.top; node != nil; {
		stack.Push(node.value)
		node = node.prev
	}
	return stack
}

func (s *Stack) AppendReverse(stack *Stack) {
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

func (s *Stack) Append(stack *Stack) {
	s.AppendReverse(stack.Reverse())
}
