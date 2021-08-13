package cache

import (
	"container/list"
	"github.com/Data-Structures-Golang/pkg/linkedlists"
)

type lruCache struct {
	capacity int
	list     linkedlists.LinkedList
	hashMap  map[interface{}]*list.Element
}

type LRUCache interface {
	// Clear resets the cache to empty state
	Clear(newCapacity int)
	// GetValue retrieves the value corresponding to a key if exists any
	GetValue(key interface{}) (interface{}, bool)
	// Insert adds a new key-value to cache or updates the existing one
	// returns true if key already existed and false otherwise
	Insert(key, value interface{}) bool
	// Delete removes a given key and its value from the cache
	// returns true if key already existed and false otherwise
	Delete(key interface{}) bool
}

func NewLRUCache(capacity int) LRUCache {
	return &lruCache{
		capacity: capacity,
		list:     linkedlists.NewLinkedList(),
		hashMap:  make(map[interface{}]*list.Element),
	}
}

// GetValue retrieves the value corresponding to a key if exists any
func (c *lruCache) GetValue(key interface{}) (interface{}, bool) {
	if node, ok := c.hashMap[key]; ok {
		lst := c.list.GetList()
		lst.MoveToFront(node)
		return node.Value, ok
	}
	return nil, false
}

// Insert adds a new key-value to cache or updates the existing one
// returns true if key already existed and false otherwise
func (c *lruCache) Insert(key, value interface{}) (ok bool) {
	lst := c.list.GetList()
	if node, ok := c.hashMap[key]; ok {
		node.Value = value
		lst.MoveToFront(node)
		return
	} else if lst.Len() < c.capacity {
		lst.Remove(lst.Back())
	}
	lst.PushFront(value)
	c.hashMap[key] = lst.Front()
	return
}

// Delete removes a given key and its value from the cache
// returns true if key already existed and false otherwise
func (c *lruCache) Delete(key interface{}) (ok bool) {
	if node, ok := c.hashMap[key]; ok {
		lst := c.list.GetList()
		lst.Remove(node)
		delete(c.hashMap, key)
	}
	return
}

// Clear resets the cache to empty state
func (c *lruCache) Clear(newCapacity int) {
	c.hashMap = make(map[interface{}]*list.Element)
	c.list = linkedlists.NewLinkedList()
	c.capacity = newCapacity
}
