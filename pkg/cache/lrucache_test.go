package cache

import (
	"github.com/Data-Structures-Golang/pkg/utils"
	"math/rand"
	"reflect"
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

func insertEachIteration(t *testing.T, cache LRUCache, array []interface{}, index int, value interface{}, capacity int) []interface{} {
	existed := cache.Insert(index, value)
	var err error
	if existed {
		if array, err = utils.ArrayRemoveDuplicatesBack(array); err != nil {
			panic(err)
		}
	}
	array = append(array, value)
	if len(array) >= capacity {
		array = array[len(array)-capacity:]
	}
	cache.AssertEqualArray(t, array)
	return array
}

func getValueEachIteration(t *testing.T, cache LRUCache, array []interface{}, index int, value interface{}, capacity int) []interface{} {
	_, ok := cache.GetValue(index)
	if ok {
		var err error
		array = append(array, value)
		if array, err = utils.ArrayRemoveDuplicatesBack(array); err != nil {
			panic(err)
		}
	}
	cache.AssertEqualArray(t, array)
	return array
}

func TestLruCache_Insert(t *testing.T) {
	capacity := 10
	cache := setupLRUCache(capacity)
	makeTwice := func(in interface{}) interface{} { return reflect.ValueOf(2 * in.(int)).Interface() }
	cache.LoopAndRun(t, 0, capacity+50, capacity, makeTwice, nil, insertEachIteration)
}

func TestLruCache_GetValue(t *testing.T) {
	capacity := 10
	cache := setupLRUCache(capacity)
	makeTwice := func(in interface{}) interface{} { return reflect.ValueOf(2 * in.(int)).Interface() }
	arr := cache.LoopAndRun(t, 0, capacity+5, capacity, makeTwice, nil, insertEachIteration)
	arr = cache.LoopAndRun(t, 5, capacity, capacity, makeTwice, arr, getValueEachIteration)
	cache.LoopAndRun(t, 0, 2*capacity+50, capacity, makeTwice, arr, insertEachIteration)
}
