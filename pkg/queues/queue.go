package queues

import (
	"container/list"
	"fmt"
)

type queue struct {
	list     *list.List
	capacity int // zero capacity is equivalent to unlimited capacity
}

type Queue interface {
	// Add adds a value to the last input (beginning) of the queue
	Add(value interface{}) error
	// AppendArray appends values of an array (priority with larger index) to the end (least waited values) of the queue
	AppendArray(array []interface{}) (err error)
	// AppendQueue concatenates another queue to the end of primary queue
	AppendQueue(queue Queue) (err error)
	//Clear resets the queue to an empty queue
	Clear() (err error)
	// GetList returns the inner doubly linked list of the queue
	GetList() *list.List
	// Len returns the length of the queue
	Len() (length int, err error)
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
	// Reverse reverses he queu
	Reverse() (queue Queue, err error)
	// ToArray converts the queue to an array
	ToArray() (array []interface{}, err error)
	// ValidateLength validates the queue is not nil and has the length of at lease one
	ValidateLength(funcName string) (err error)
}

func NewQueue(capacity ...int) Queue {
	var maxLength int
	if len(capacity) == 0 {
		maxLength = 0
	} else {
		maxLength = capacity[0]
	}
	return &queue{
		list:     list.New(),
		capacity: maxLength, // zero capacity is equivalent to unlimited capacity
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
func (q *queue) Clear() (err error) {
	// check if queue is nil
	if _, err = q.Len(); err != nil {
		return
	}
	q.list = list.New()
	return
}

// GetList returns the inner doubly linked list of the queue
func (q *queue) GetList() *list.List {
	return q.list
}

// Len returns the length of the queue
func (q *queue) Len() (length int, err error) {
	if q == nil {
		return length, fmt.Errorf("error: queue cannot be nil")
	}
	length = q.list.Len()
	return
}

// Peek returns the first in of the queue
func (q *queue) Peek() (value interface{}, err error) {
	if err = q.ValidateLength("Peek"); err != nil {
		return
	}
	q.list.Back()
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
		return fmt.Errorf("%s: cannot operate with queue of length zero", funcName)
	}
	return
}

func (q *queue) checkCapacity() (err error) {
	if q.capacity != 0 && q.list.Len() >= q.capacity {
		return fmt.Errorf("reached queue capacity= %v", q.capacity)
	}
	return
}
