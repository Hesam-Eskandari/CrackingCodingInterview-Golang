package Stacks

import (
	"reflect"
)

type (
	node struct {
		value interface{}
		prev  *node
	}
	Stack struct {
		top    *node
		length int
		min    *node
	}
)

// NewStack constructs and returns a new empty stack
func NewStack() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
		min:    nil,
	}
}

// Top returns the value of the top of the stack
func (s *Stack) Top() interface{} {
	if s.top == nil {
		return nil
	}
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
	s.calcMin(value, false)
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
	s.calcMin(detachedHead, true)
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
	stack := NewStack()
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

// Append appends a second stack on top of the original stack
func (s *Stack) Append(stack *Stack) {
	s.AppendReverse(stack.Reverse())
}

func (s *Stack) calcMin(value interface{}, isPop bool) {
	var valueType interface{}
	if s.min != nil {
		valueType = s.min.value
	} else if value != nil {
		s.min = &node{
			value: value,
		}
		return
	} else {
		return
	}
	if isPop && value == s.min.value {
		s.min = s.min.prev
		return
	}
	switch valueType.(type) {
	case int:
		if value.(int) <= s.min.value.(int) {
			s.pushMin(value)
		}
	case int8:
		if value.(int8) <= s.min.value.(int8) {
			s.pushMin(value)
		}
	case int16:
		if value.(int16) <= s.min.value.(int16) {
			s.pushMin(value)
		}
	case int32:
		if value.(int32) <= s.min.value.(int32) {
			s.pushMin(value)
		}
	case int64:
		if value.(int64) <= s.min.value.(int64) {
			s.pushMin(value)
		}
	case float32:
		if value.(float32) <= s.min.value.(float32) {
			s.pushMin(value)
		}
	case float64:
		if value.(float64) <= s.min.value.(float64) {
			s.pushMin(value)
		}
	default:
		return
	}

}

func (s *Stack) pushMin(value interface{}) {
	if s.min == nil {
		s.min = &node{value, nil}
	} else {
		node := &node{
			value: value,
			prev:  s.min,
		}
		s.min = node
	}
}

// SortN replace a stack with a sorted version of it
// only one additional stack is allowed to be used
// only push, pop, top and len methods are allowed to be used
// runs in O(n) of space and O(n^2) of time complexity where n = stack.Len()
// assume that the values are of type int
func (s *Stack) SortN() {
	if s.Top() == nil {
		return
	}
	switch s.Top().(type) {
	case int:
	default:
		return
	}
	stack := NewStack()
	for s.Top() != nil {
		if stack.Top() == nil {
			stack.Push(s.Pop())
			continue
		}
		value := s.Top().(int)
		count := 0
		for stack.Top() != nil && stack.Top().(int) < value {
			s.Push(stack.Pop())
			count += 1
		}
		stack.Push(value)
		for ; count > 0; count-- {
			stack.Push(s.Pop())
		}
		s.Pop()
	}
	for stack.Top() != nil {
		s.Push(stack.Pop())
	}
}
