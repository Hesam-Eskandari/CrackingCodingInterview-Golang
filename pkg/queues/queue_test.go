package queues

import (
	"math/rand"
	"testing"
	"time"
)

func assertEqualValue(t *testing.T, value1, value2 interface{}) {
	if value1 != value2 {
		t.Errorf("assertEqualValue: first value: %v and second value: %v are not equal", value1, value2)
	}
}

func setupQueue(capacity, length int) (queue Queue, array []interface{}) {
	rand.Seed(time.Now().UnixNano())
	if capacity == -1 {
		capacity = rand.Intn(20) + 20
	}
	arrayCap := length
	if capacity != 0 && capacity < length {
		arrayCap = rand.Intn(capacity/2) + capacity/2
	}
	queue = NewQueue(capacity)
	array = make([]interface{}, 0, arrayCap)
	for index := 0; index < arrayCap; index++ {
		array = append(array, rand.Intn(arrayCap*2))
	}
	if err := queue.AppendArray(array); err != nil {
		panic(err)
	}
	return
}

func TestQueue_AppendArray(t *testing.T) {
	queue, array := setupQueue(0, 5)
	queue.AssertEqualArray(t, array)

}

func TestQueue_ToArray(t *testing.T) {
	queue, _ := setupQueue(0, 100)
	queueArray, _ := queue.ToArray()
	queue.AssertEqualArray(t, queueArray)

}

func TestQueue_Add(t *testing.T) {
	queue, array := setupQueue(100, 80)
	rand.Seed(time.Now().UnixNano())
loop:
	for index := 0; index < 50; index++ {
		value := rand.Intn(100)
		if err := queue.Add(value); err != nil {
			if _, ok := err.(*QueueIsFullException); ok {
				break loop
			}
			queue.AssertError(t, err)
		}
		array = append(array, value)
		queue.AssertEqualArray(t, array)
	}
}

func TestQueue_Peek(t *testing.T) {
	queue, _ := setupQueue(50, 45)
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)
	if err := queue.Add(value); err != nil {
		queue.AssertError(t, err)
	}
	queue.AssertEqualPeek(t, value)
}

func TestQueue_Pop(t *testing.T) {
	length := 45
	queue, array := setupQueue(50, length)
	for index := 0; index < length-1; index++ {
		value, err := queue.Pop()
		if err != nil {
			queue.AssertError(t, err)
		}
		assertEqualValue(t, value, array[index])
	}
}
