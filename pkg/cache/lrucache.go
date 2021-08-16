package cache

import (
	"container/list"
	"testing"
)

type lruCache struct {
	capacity int
	list     *list.List
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
	// GetList returns the underlying linked list
	GetList() *list.List
	// GetValue retrieves the value corresponding to a key if exists any
	GetValue(key interface{}) (interface{}, bool)
	// Insert adds a new key-value to cache or updates the existing one
	// returns true if key already existed and false otherwise
	Insert(key, value interface{}) bool
	// LoopAndRun loops over a compatible function from start to end
	LoopAndRun(t *testing.T, start, end, capacity int, value func(in interface{}) interface{}, arr []interface{}, function func(
		t *testing.T, cache LRUCache, array []interface{}, index int, value interface{}, capacity int) []interface{}) (array []interface{})
	// ToArrayKeys returns an array with keys in linked list with same order and size
	ToArrayKeys() []interface{}
	// ToArrayValues returns an array with values in linked list with same order and size
	ToArrayValues() []interface{}
}

func NewLRUCache(capacity int) LRUCache {
	return &lruCache{
		capacity: capacity,
		list:     list.New(),
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
	arr := c.ToArrayKeys()
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
		lst := c.list
		lst.MoveToBack(node)
		return node.Value.(KeyValue).Value, ok
	}
	return nil, ok
}

// Insert adds a new key-value to cache or updates the existing one
// returns true if key already existed and false otherwise
func (c *lruCache) Insert(key, value interface{}) bool {
	lst := c.list
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
	var node *list.Element
	node, ok = c.hashMap[key]
	if ok {
		lst := c.list
		lst.Remove(node)
		delete(c.hashMap, key)
	}
	return
}

// Clear resets the cache to empty state
func (c *lruCache) Clear(newCapacity int) {
	c.hashMap = make(map[interface{}]*list.Element)
	c.list = list.New()
	c.capacity = newCapacity
}

// GetList returns the underlying linked list
func (c *lruCache) GetList() *list.List {
	return c.list
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

// ToArrayKeys returns an array with keys in linked list with same order and size
func (c *lruCache) ToArrayKeys() []interface{} {
	var arr []interface{}
	if c == nil {
		return arr
	}
	element := c.list.Front()
	for element != nil {
		arr = append(arr, element.Value.(KeyValue).Key)
		element = element.Next()
	}
	return arr
}

// ToArrayValues returns an array with values in linked list with same order and size
func (c *lruCache) ToArrayValues() []interface{} {
	var arr []interface{}
	if c == nil {
		return arr
	}
	element := c.list.Front()
	for element != nil {
		arr = append(arr, element.Value.(KeyValue).Value)
		element = element.Next()
	}
	return arr
}
