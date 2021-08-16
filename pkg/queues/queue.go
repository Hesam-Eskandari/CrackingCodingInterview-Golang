package queues

import (
	"container/list"
	"testing"
)

type queue struct {
	capacity                 int // zero capacity is equivalent to unlimited capacity
	list                     *list.List
	queueIsEmptyException    *QueueIsEmptyException
	queueIsFullException     *QueueIsFullException
	queueHasNilListException *QueueHasNilListException
}

type Queue interface {
	// Add adds a value to the last input (beginning) of the queue
	Add(value interface{}) error
	// AppendArray appends values of an array (priority with larger index) to the end (least waited values) of the queue
	AppendArray(array []interface{}) (err error)
	// AppendQueue concatenates another queue to the end of primary queue
	AppendQueue(queue Queue) (err error)
	//Clear resets the queue to an empty queue
	Clear(capacity int) (err error)
	// GetList returns the inner doubly linked list of the queue
	GetList() *list.List
	// Len returns the length of the queue
	Len() (length int, err error)
	// LoopAndRun loops over a compatible function from start to end
	LoopAndRun(t *testing.T, start, end, capacity int, value func(in interface{}) (interface{}, error),
		arr []interface{}, function func(t *testing.T, queue Queue, array []interface{}, index int, value interface{},
			capacity int) ([]interface{}, error)) (array []interface{}, err error)
	// Peek returns the first in of the queue
	Peek() (value interface{}, err error)
	// Pop removes the longest waited value (beginning) in the queue and returns its value
	Pop() (value interface{}, err error)
	// RemoveFromBegin removes arbitrary number of occurrences of a value from the early input of the queue
	RemoveFromBegin(value interface{}, howMany int) (removed int, err error)
	// RemoveFromEnd removes arbitrary number of occurrences of a value from the last input of the queue
	RemoveFromEnd(value interface{}, howMany int) (removed int, err error)
	// Replace replaces the queue with another queue
	Replace(queue Queue) (err error)
	// Reverse reverses he queue
	Reverse() (queue Queue, err error)
	// ToArray converts the queue to an array
	ToArray() (array []interface{}, err error)
	// ValidateLength validates the queue is not nil and has the length of at lease one
	ValidateLength(funcName string) (err error)
	// AssertEqualArray verifies the queue and an array have the same values in the same order
	AssertEqualArray(t *testing.T, array []interface{})
	// AssertEqualPeek verifies the value of the peek
	AssertEqualPeek(t *testing.T, value interface{})
	// AssertValueAtIndex verifies the value located in an index from the end of queue
	AssertValueAtIndex(t *testing.T, index int, value interface{})
	// AssertError checks if an error is not nil during testing
	AssertError(t *testing.T, err error)
}

func NewQueue(capacity ...int) Queue {
	var maxLength int
	if len(capacity) == 0 {
		maxLength = 0
	} else {
		maxLength = capacity[0]
	}
	return &queue{
		capacity:                 maxLength, // zero capacity is equivalent to unlimited capacity
		list:                     list.New(),
		queueIsEmptyException:    &QueueIsEmptyException{""},
		queueIsFullException:     &QueueIsFullException{capacity: maxLength},
		queueHasNilListException: &QueueHasNilListException{},
	}
}

// Add adds a value to the last input (beginning) of the queue
func (q *queue) Add(value interface{}) (err error) {
	// check if queue is nil
	if _, err = q.Len(); err != nil {
		return
	}
	if err = q.checkCapacity(); err != nil {
		return
	}

	q.list.PushBack(value)
	return
}

// AppendArray appends values of an array (priority with larger index) to the end (least waited values) of the queue
func (q *queue) AppendArray(array []interface{}) (err error) {
	if _, err = q.Len(); err != nil {
		return
	}
	for index := range array {
		if err = q.checkCapacity(); err != nil {
			return
		}
		q.list.PushFront(array[len(array)-index-1])
	}
	return
}

// AppendQueue concatenates another queue to the end of primary queue
func (q *queue) AppendQueue(queue Queue) (err error) {
	if _, err = q.Len(); err != nil {
		return
	}
	if _, err = queue.Len(); err != nil {
		return
	}
	for node := queue.GetList().Back(); node != nil; {
		if err = q.checkCapacity(); err != nil {
			return
		}
		q.list.PushFront(node)
		node = node.Prev()
	}
	return
}

//Clear resets the queue to an empty queue
func (q *queue) Clear(capacity int) (err error) {
	// check if queue is nil
	if _, err = q.Len(); err != nil {
		return
	}
	q.list = list.New()
	q.capacity = capacity
	return
}

// GetList returns the inner doubly linked list of the queue
func (q *queue) GetList() *list.List {
	return q.list
}

// Len returns the length of the queue
func (q *queue) Len() (length int, err error) {
	if q == nil {
		return length, &QueueIsNilException{}
	}
	if q.list == nil {
		return length, q.queueHasNilListException
	}
	length = q.list.Len()
	return
}

// LoopAndRun loops over a compatible function from start to end
func (q *queue) LoopAndRun(t *testing.T, start, end, capacity int, value func(in interface{}) (interface{}, error),
	arr []interface{}, function func(t *testing.T, queue Queue, array []interface{}, index int, value interface{},
		capacity int) ([]interface{}, error)) (array []interface{}, err error) {
	if arr == nil {
		arr = make([]interface{}, 0, capacity)
		if err = q.Clear(capacity); err != nil {
			return
		}
	}
	var val interface{}
	for index := start; index < end; index++ {
		val, err = value(index)
		if err != nil {
			return
		}
		arr, err = function(t, q, arr, index, val, capacity)
		if err != nil {
			return
		}
	}
	return arr, nil
}

// Peek returns the first in of the queue
func (q *queue) Peek() (value interface{}, err error) {
	if err = q.ValidateLength("Peek"); err != nil {
		return
	}
	value = q.list.Back().Value
	return
}

// Pop removes the longest waited value (beginning) in the queue
func (q *queue) Pop() (value interface{}, err error) {
	if err = q.ValidateLength("Pop"); err != nil {
		return
	}
	value = q.list.Front().Value
	q.list.Remove(q.list.Front())
	return
}

// RemoveFromBegin removes arbitrary number of occurrences of a value from the early input of the queue
func (q *queue) RemoveFromBegin(value interface{}, howMany int) (removed int, err error) {
	if err = q.ValidateLength("RemoveFromBegin"); err != nil {
		return
	}
	if howMany == 0 || howMany > q.list.Len() {
		howMany = q.list.Len()
	}
	for node := q.list.Back(); node != nil; {
		if node.Value == value {
			q.list.Remove(node)
			removed += 1
		}
		if removed >= howMany {
			break
		}
		node = node.Prev()
	}
	return
}

// RemoveFromEnd removes arbitrary number of occurrences of a value from the last input of the queue
func (q *queue) RemoveFromEnd(value interface{}, howMany int) (removed int, err error) {
	if err = q.ValidateLength("RemoveFromEnd"); err != nil {
		return
	}
	if howMany == 0 || howMany > q.list.Len() {
		howMany = q.list.Len()
	}
	for node := q.list.Front(); node != nil; {
		if node.Value == value {
			q.list.Remove(node)
			removed += 1
		}
		if removed >= howMany {
			break
		}
		node = node.Next()
	}
	return
}

// Replace replaces the queue with another queue
func (q *queue) Replace(queue Queue) (err error) {
	q.list = queue.GetList()
	return
}

// Reverse reverses he queue
func (q *queue) Reverse() (queue Queue, err error) {
	// check if queue is nil
	length, err := q.Len()
	if err != nil {
		return
	} else if length == 0 {
		return q, nil
	}
	for node := q.list.Front(); node != nil; {
		queue.GetList().PushFront(node)
		node = node.Next()
	}
	return
}

// ToArray converts the queue to an array
func (q *queue) ToArray() (array []interface{}, err error) {
	// check if queue is nil
	if _, err = q.Len(); err != nil {
		return
	}
	for node := q.list.Front(); node != nil; {
		array = append(array, node.Value)
		node = node.Next()
	}
	return
}

// ValidateLength validates the queue is not nil and has the length of at lease one
func (q *queue) ValidateLength(funcName string) (err error) {
	var length int
	length, err = q.Len()
	if err != nil {
		return
	}
	if length == 0 {
		q.queueIsEmptyException.funcName = funcName
		return q.queueIsEmptyException
	}
	return
}

func (q *queue) checkCapacity() (err error) {
	if q.capacity != 0 && q.list.Len() >= q.capacity {
		return q.queueIsFullException
	}
	return
}

// 	AssertEqualArray verifies the queue and an array have the same values in the same order
func (q *queue) AssertEqualArray(t *testing.T, array []interface{}) {
	// validate queue
	if err := q.ValidateLength("AssertEqualArray"); err != nil {
		t.Errorf(err.Error())
	}
	for index, node := len(array)-1, q.GetList().Back(); node != nil; index-- {
		if node.Value != array[index] {
			queueArray, _ := q.ToArray()
			t.Errorf("queue %v does not match to array %v", queueArray, array)
			break
		}
		node = node.Prev()
	}
}

// AssertEqualPeek verifies the value of the peek
func (q *queue) AssertEqualPeek(t *testing.T, value interface{}) {
	// validate queue
	if err := q.ValidateLength("AssertEqualPeek"); err != nil {
		t.Errorf(err.Error())
	}
	peek, _ := q.Peek()
	if peek != value {
		t.Errorf("AssertEqualPeek: peek of the queue is expected to be %v but it is equal to %v", value, peek)
	}
}

// AssertValueAtIndex verifies the value located in an index from the end of queue
func (q *queue) AssertValueAtIndex(t *testing.T, index int, value interface{}) {
	// validate queue
	if err := q.ValidateLength("AssertValueAtIndex"); err != nil {
		t.Errorf(err.Error())
	}
	if index > q.list.Len()-1 {
		t.Errorf("AssertValueAtIndex: index error: index %v should be less than queue length: %v", index, q.list.Len())
	}
	array, _ := q.ToArray()
	val := array[index]
	if val != value {
		t.Errorf("AssertValueAtIndex: value of the queue at index %v is expected"+
			" to be %v but it is equal to %v", index, value, val)
	}
}

// AssertError checks if an error is not nil during testing
func (q *queue) AssertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}
