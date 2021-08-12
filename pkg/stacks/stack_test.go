package Stacks

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func (s *Stack) assertEqualArray(t *testing.T, array interface{}) {
	if s == nil {
		panic("assertEqualArray: expected a stack, received nil")
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		top := s.top
		if s.length != arr.Len() {
			t.Errorf("assertEqualArray error: stack and array don't have the same length")
		} else {
			for index := arr.Len() - 1; index >= 0; index-- {

				if top.value != arr.Index(index).Interface() {
					t.Errorf("assertEqualArray error: the element %v of stack with value %v is not "+
						"equal to item %v from array with value %v", arr.Len()-index, top.value, index, arr.Index(index))
					break
				}
				top = top.prev
			}
		}
	default:
		panic(fmt.Sprintf("expected an array, received %v", reflect.TypeOf(array).Kind()))
	}
}

func (s *Stack) assertEqual(t *testing.T, stack *Stack) {
	if s == nil || stack == nil {
		panic("assertEqual: expected a stack, received nil")
	}
	if s.length != stack.length {
		t.Errorf("assertEqual: two stacks don't have the same length. \n"+
			"stack 1: %v \n stack 2: %v", s.ToArray(), stack.ToArray())
	}
}

func (s *Stack) assertEqualMin(t *testing.T, value interface{}) {
	if s == nil {
		panic("assertEqualMin: expected a stack, received nil")
	}
	if (s.min != nil && value == nil) || (s.min == nil && value != nil) {
		t.Errorf("assertEqualMin: single nil error,  %v is not equal to %v", s.min, value)
	} else if s.min != nil && s.min.value != value {
		t.Errorf("assertEqualMin: separate values error, %v is not equal to %v", s.min.value, value)
	}
}

func (s *Stack) assertEqualTop(t *testing.T, value interface{}) {
	if s == nil {
		panic("assertEqualTop: expected a stack, received nil")
	}
	if s.top != nil && value != s.top.value {
		t.Errorf("assertEqualTop: expected the top value to be %v but"+
			" it's equal to %v", value, s.top.value)
	}
	if s.top == nil && value != nil {
		t.Errorf("assertEqualTop: expected the top value to be %v but"+
			" it's equal to %v", value, s.top.value)
	}
}

func (s *Stack) assertLength(t *testing.T, length int) {
	if s == nil {
		panic("assertEqualArray: expected a stack, received nil")
	}
	if s.length != length {
		t.Errorf("expected the stack to be of length %v, but its length is %v", length, s.length)
	}
}

func setUpStack() (stack *Stack, expectedArray []int) {
	stack = NewStack()
	rand.Seed(time.Now().UnixNano())
	expectedArray = make([]int, 0, rand.Intn(20)+10)
	for index := 0; index < cap(expectedArray); index++ {
		expectedArray = append(expectedArray, rand.Intn(100))
	}
	stack.AppendArray(expectedArray)
	return
}

func TestStack_ToArray(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertEqualArray(t, expectedArray)

}

func TestStack_Len(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertLength(t, len(expectedArray))
}

func TestStack_Top(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertEqualTop(t, expectedArray[len(expectedArray)-1])
}

func TestStack_Push(t *testing.T) {
	stack, expectedArray := setUpStack()
	add := 999
	stack.Push(add)
	stack.assertEqualTop(t, add)
	stack.assertLength(t, len(expectedArray)+1)
	stack.assertEqualArray(t, append(expectedArray, add))
}

func TestStack_Pop(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.Pop()
	var newTopValue interface{}
	if len(expectedArray) > 1 {
		newTopValue = expectedArray[len(expectedArray)-2]
	} else {
		newTopValue = nil
	}
	stack.assertEqualTop(t, newTopValue)
}

func TestStack_Reverse(t *testing.T) {
	stack, initialArray := setUpStack()
	reversedArray := make([]interface{}, len(initialArray))
	for index, value := range initialArray {
		reverseIndex := len(reversedArray) - index - 1
		reversedArray[reverseIndex] = value
	}
	stack = stack.Reverse()
	stack.assertLength(t, len(initialArray))
	stack.assertEqualTop(t, initialArray[0])
	stack.assertEqual(t, stack.Reverse())
	stack.assertEqualArray(t, reversedArray)
}

func TestStack_AppendArray(t *testing.T) {
	stack, initialArray := setUpStack()
	arrayToAppend := []int{21, 22, 23, 24, 25, 26, 27, 28, 29}
	stack.AppendArray(arrayToAppend)
	stack.assertLength(t, len(initialArray)+len(arrayToAppend))
	stack.assertEqualTop(t, arrayToAppend[len(arrayToAppend)-1])
	stack.assertEqualArray(t, append(initialArray, arrayToAppend...))
}

func TestStack_AppendReverse(t *testing.T) {
	firstStack, firstArray := setUpStack()
	secondStack, secondArray := setUpStack()
	firstStack.AppendReverse(secondStack)
	firstStack.assertLength(t, len(firstArray)+len(secondArray))
	firstStack.assertEqualTop(t, secondArray[0])
}

func TestStack_Min(t *testing.T) {
	stack := NewStack()
	arr := []int{4, 2, 6, 3, 2, 4, 7, 1, 8, 0, 2}
	min := arr[0]
	oldMin := min
	for index, value := range arr {
		t.Run(fmt.Sprintf("before push. index: %v, value: %v", index, value), func(t *testing.T) {
			if index == 0 {
				stack.assertEqualMin(t, nil)
			} else {
				stack.assertEqualMin(t, min)
			}
		})
		stack.Push(value)
		if value < min {
			min = value
		}
		t.Run(fmt.Sprintf("after push. index: %v, value: %v", index, value), func(t *testing.T) {
			stack.assertEqualMin(t, min)
		})
		stack.Pop()
		t.Run(fmt.Sprintf("after pop. index: %v, value: %v", index, value), func(t *testing.T) {
			if index == 0 {
				stack.assertEqualMin(t, nil)
			} else {
				stack.assertEqualMin(t, oldMin)
			}
		})
		stack.Push(value)
		oldMin = min
	}
}
