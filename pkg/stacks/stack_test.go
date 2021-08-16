package Stacks

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func (s *stack) assertEqualArray(t *testing.T, array interface{}) {
	if s == nil {
		panic("assertEqualArray: expected a newStack, received nil")
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		top := s.top
		if s.length != arr.Len() {
			t.Errorf("assertEqualArray error: newStack and array don't have the same length")
		} else {
			for index := arr.Len() - 1; index >= 0; index-- {

				if top.value != arr.Index(index).Interface() {
					t.Errorf("assertEqualArray error: the element %v of newStack with value %v is not "+
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

func (s *stack) assertEqual(t *testing.T, newStack *stack) {
	if s == nil || newStack == nil {
		panic("assertEqual: expected a newStack, received nil")
	}
	if s.length != newStack.length {
		t.Errorf("assertEqual: two stacks don't have the same length. \n"+
			"newStack 1: %v \n newStack 2: %v", s.ToArray(), newStack.ToArray())
	}
}

func (s *stack) assertEqualMin(t *testing.T, value interface{}) {
	if s == nil {
		panic("assertEqualMin: expected a newStack, received nil")
	}
	if (s.min != nil && value == nil) || (s.min == nil && value != nil) {
		t.Errorf("assertEqualMin: single nil error,  %v is not equal to %v", s.min, value)
	} else if s.min != nil && s.min.value != value {
		t.Errorf("assertEqualMin: separate values error, %v is not equal to %v", s.min.value, value)
	}
}

func (s *stack) assertEqualTop(t *testing.T, value interface{}) {
	if s == nil {
		panic("assertEqualTop: expected a newStack, received nil")
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

func (s *stack) assertLength(t *testing.T, length int) {
	if s == nil {
		panic("assertEqualArray: expected a newStack, received nil")
	}
	if s.length != length {
		t.Errorf("expected the newStack to be of length %v, but its length is %v", length, s.length)
	}
}

func (s *stack) raiseError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}

func setUpStack(size ...int) (newStack *stack, expectedArray []int) {
	newStack = NewStack()
	rand.Seed(time.Now().UnixNano())
	var capacity int
	if len(size) == 0 {
		capacity = rand.Intn(20) + 10
	} else {
		capacity = size[0]
	}
	expectedArray = make([]int, 0, capacity)
	for index := 0; index < capacity; index++ {
		expectedArray = append(expectedArray, rand.Intn(1000))
	}
	newStack.AppendArray(expectedArray)
	return
}

func TestStack_ToArray(t *testing.T) {
	newStack, expectedArray := setUpStack(10)
	newStack.assertEqualArray(t, expectedArray)

}

func TestStack_Len(t *testing.T) {
	newStack, expectedArray := setUpStack()
	newStack.assertLength(t, len(expectedArray))
}

func TestStack_Top(t *testing.T) {
	newStack, expectedArray := setUpStack()
	newStack.assertEqualTop(t, expectedArray[len(expectedArray)-1])
}

func TestStack_Push(t *testing.T) {
	newStack, expectedArray := setUpStack()
	add := 999
	newStack.Push(add)
	newStack.assertEqualTop(t, add)
	newStack.assertLength(t, len(expectedArray)+1)
	newStack.assertEqualArray(t, append(expectedArray, add))
}

func TestStack_Pop(t *testing.T) {
	newStack, expectedArray := setUpStack()
	newStack.Pop()
	var newTopValue interface{}
	if len(expectedArray) > 1 {
		newTopValue = expectedArray[len(expectedArray)-2]
	} else {
		newTopValue = nil
	}
	newStack.assertEqualTop(t, newTopValue)
}

func TestStack_Reverse(t *testing.T) {
	newStack, initialArray := setUpStack()
	reversedArray := make([]interface{}, len(initialArray))
	for index, value := range initialArray {
		reverseIndex := len(reversedArray) - index - 1
		reversedArray[reverseIndex] = value
	}
	newStack = newStack.Reverse()
	newStack.assertLength(t, len(initialArray))
	newStack.assertEqualTop(t, initialArray[0])
	newStack.assertEqual(t, newStack.Reverse())
	newStack.assertEqualArray(t, reversedArray)
}

func TestStack_AppendArray(t *testing.T) {
	newStack, initialArray := setUpStack()
	_, arrayToAppend := setUpStack()
	newStack.AppendArray(arrayToAppend)
	newStack.assertLength(t, len(initialArray)+len(arrayToAppend))
	newStack.assertEqualTop(t, arrayToAppend[len(arrayToAppend)-1])
	newStack.assertEqualArray(t, append(initialArray, arrayToAppend...))
}

func TestStack_AppendReverse(t *testing.T) {
	firstStack, firstArray := setUpStack()
	secondStack, secondArray := setUpStack()
	if err := firstStack.AppendReverse(secondStack); err != nil {
		firstStack.raiseError(t, err)
	}
	firstStack.assertLength(t, len(firstArray)+len(secondArray))
	firstStack.assertEqualTop(t, secondArray[0])
}

func TestStack_Min(t *testing.T) {
	newStack := NewStack()
	_, arr := setUpStack()
	min := arr[0]
	oldMin := min
	for index, value := range arr {
		t.Run(fmt.Sprintf("before push. index: %v, value: %v", index, value), func(t *testing.T) {
			if index == 0 {
				newStack.assertEqualMin(t, nil)
			} else {
				newStack.assertEqualMin(t, min)
			}
		})
		newStack.Push(value)
		if value < min {
			min = value
		}
		t.Run(fmt.Sprintf("after push. index: %v, value: %v", index, value), func(t *testing.T) {
			newStack.assertEqualMin(t, min)
		})
		newStack.Pop()
		t.Run(fmt.Sprintf("after pop. index: %v, value: %v", index, value), func(t *testing.T) {
			if index == 0 {
				newStack.assertEqualMin(t, nil)
			} else {
				newStack.assertEqualMin(t, oldMin)
			}
		})
		newStack.Push(value)
		oldMin = min
	}
}

func TestStack_SortN(t *testing.T) {
	newStack, arr := setUpStack(1000)
	sort.Ints(arr)
	newStack.SortN()
	newStack.assertEqualArray(t, arr)
}

func TestStack_Sort(t *testing.T) {
	newStack, arr := setUpStack(50000)
	sort.Ints(arr)
	newStack = newStack.Sort()
	newStack.assertEqualArray(t, arr)

	arrayFloat := []float64{12, 13.1, 0.001, 10, 0}
	stackFloat := NewStack()
	stackFloat.AppendArray(arrayFloat)
	sort.Float64s(arrayFloat)
	newStack.Sort()
	newStack.assertEqualArray(t, arr)
}
