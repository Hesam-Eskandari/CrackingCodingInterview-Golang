package cache

import (
	"fmt"
	"math/rand"
	"testing"
)

func setupLRUCache(arr ...int) LRUCache {
	var capacity int
	if len(arr) == 0 {
		capacity = rand.Intn(10) + 10
	} else {
		capacity = arr[0]
	}
	return NewLRUCache(capacity)

}

func TestLruCache_Insert(t *testing.T) {
	capacity := 10
	cache := setupLRUCache(capacity)
	arr := make([]int, 0, 2*capacity)
	for index := 0; index < 2*capacity; index++ {
		cache.Insert(index, 2*index)
		arr = append(arr, 2*index)
		linkedList := cache.GetLinkedList()
		var start int
		if index < capacity {
			start = 0
		} else {
			start = index - capacity + 1
		}
		fmt.Println("here", linkedList.ToArray())
		linkedList.AssertEqualArray(t, arr[start:index+1])
	}
}
