package queues

import "fmt"

type QueueIsFullException struct {
	capacity int
}

func (exp QueueIsFullException) Error() string {
	return fmt.Sprintf("reached queue capacity: %v", exp.capacity)
}

type QueueIsEmptyException struct {
	funcName string
}

func (exp QueueIsEmptyException) Error() string {
	return fmt.Sprintf("method %s cannot operate with queue of length zero", exp.funcName)
}

type QueueIsNilException struct{}

func (exp QueueIsNilException) Error() string {
	return fmt.Sprintf("error: queue cannot be nil")
}

type QueueHasNilListException struct{}

func (exp QueueHasNilListException) Error() string {
	return fmt.Sprintf("error: list in the queue cannot be nil")
}
