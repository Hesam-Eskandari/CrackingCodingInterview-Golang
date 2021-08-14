package cache

import (
	"container/list"
	"github.com/Data-Structures-Golang/pkg/linkedlists"
	"testing"
)

type lruCache struct {
	capacity int
	List     linkedlists.LinkedList
	hashMap  map[interface{}]*list.Element
}

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

type LRUCache interface {
	// AssertEqualArray tests if values in a cache and an array are equal and are in the same order
	AssertEqualArray(t *testing.T, array []interface{})
	// Clear resets the cache to empty state
	Clear(newCapacity int)
	// Delete removes a given key and its value from the cache
	// returns true if key already existed and false otherwise
	Delete(key interface{}) bool
	// GetLinkedList returns the underlying linked list
	GetLinkedList() linkedlists.LinkedList
	// GetValue retrieves the value corresponding to a key if exists any
	GetValue(key interface{}) (interface{}, bool)
	// Insert adds a new key-value to cache or updates the existing one
	// returns true if key already existed and false otherwise
	Insert(key, value interface{}) bool
	// LoopAndRun loops over a compatible function from start to end
	LoopAndRun(t *testing.T, start, end, capacity int, value func(in interface{}) interface{}, arr []interface{}, function func(
		t *testing.T, cache LRUCache, array []interface{}, index int, value interface{}, capacity int) []interface{}) (array []interface{})
	ToArray() []interface{}
}

func NewLRUCache(capacity int) LRUCache {
	return &lruCache{
		capacity: capacity,
		List:     linkedlists.NewLinkedList(),
		hashMap:  make(map[interface{}]*list.Element),
	}
}

func NewKeyValue(key, value interface{}) KeyValue {
	return KeyValue{
		Key:   key,
		Value: value,
	}
}

// AssertEqualArray tests if values in a cache and an array are equal and are in the same order
func (c *lruCache) AssertEqualArray(t *testing.T, array []interface{}) {
	if c == nil {
		panic("assertEqualArray: nil cache inserted")
	}
	arr := c.ToArray()
	if arr == nil {
		if len(array) > 0 {
			t.Errorf("AssertEqualArray: expected %v, got %v", array, nil)
			return
		}
	}
	if len(arr) != len(array) {
		t.Errorf("assertEqualArray: cache length is %v it's expected to be %v", len(arr), len(array))

		return
	}
	for index := range array {
		if arr[index] != array[index] {
			t.Errorf("assertEqualArray: cache %v does not match to array %v", arr, array)
			return
		}
	}
}

// GetValue retrieves the value corresponding to a key if exists any
func (c *lruCache) GetValue(key interface{}) (interface{}, bool) {
	node, ok := c.hashMap[key]
	if ok {
		lst := c.List.GetList()
		lst.MoveToBack(node)
		return node.Value.(KeyValue).Value, ok
	}
	return nil, ok
}

// Insert adds a new key-value to cache or updates the existing one
// returns true if key already existed and false otherwise
func (c *lruCache) Insert(key, value interface{}) bool {
	lst := c.List.GetList()
	node, ok := c.hashMap[key]
	if ok {
		node.Value = NewKeyValue(key, value)
		lst.MoveToBack(node)
		return ok
	} else if lst.Len() >= c.capacity {
		oldNode := lst.Front()
		c.Delete(oldNode.Value.(KeyValue).Key)
	}
	lst.PushBack(NewKeyValue(key, value))
	c.hashMap[key] = lst.Back()
	return ok
}

// Delete removes a given key and its value from the cache
// returns true if key already existed and false otherwise
func (c *lruCache) Delete(key interface{}) (ok bool) {
	node, ok := c.hashMap[key]
	if ok {
		lst := c.List.GetList()
		lst.Remove(node)
		delete(c.hashMap, key)
	}
	return
}

// Clear resets the cache to empty state
func (c *lruCache) Clear(newCapacity int) {
	c.hashMap = make(map[interface{}]*list.Element)
	c.List = linkedlists.NewLinkedList()
	c.capacity = newCapacity
}

// GetLinkedList returns the underlying linked list
func (c *lruCache) GetLinkedList() linkedlists.LinkedList {
	return c.List
}

// LoopAndRun loops over a compatible function from start to end
func (c *lruCache) LoopAndRun(t *testing.T, start, end, capacity int,
	value func(in interface{}) interface{}, arr []interface{}, function func(t *testing.T, cache LRUCache,
		array []interface{}, index int, value interface{}, capacity int) []interface{}) []interface{} {
	if arr == nil {
		arr = make([]interface{}, 0, capacity)
		c.Clear(capacity)
	}
	for index := start; index < end; index++ {
		val := value(index)
		arr = function(t, c, arr, index, val, capacity)
	}
	return arr
}

// ToArray returns an array with values in linked List with same order and size
func (c *lruCache) ToArray() []interface{} {
	var arr []interface{}
	if c == nil {
		return arr
	}
	element := c.List.GetList().Front()
	for element != nil {
		arr = append(arr, element.Value.(KeyValue).Value)
		element = element.Next()
	}
	return arr
}
