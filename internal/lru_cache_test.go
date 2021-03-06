package internal_test

import (
	"testing"

	"github.com/nitinstp23/lru-cache-go/internal"
)

func TestCreateCache(t *testing.T) {
	tt := []struct {
		name        string
		capacity    int
		cacheOutput internal.LRUCache
	}{
		{"create empty cache with capacity of 2 elements", 2, internal.LRUCache{Capacity: 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lruCache := internal.CreateCache(tc.capacity)

			if lruCache.Capacity != tc.capacity {
				t.Fatalf("cache capacity should be %v; got %v", tc.capacity, lruCache.Capacity)
			}

			if lruCache.Size != 0 {
				t.Fatalf("cache store length should be 0; got %v", lruCache.Size)
			}
		})
	}
}

func TestPutCache(t *testing.T) {
	tt := []struct {
		name      string
		capacity  int
		inputs    [][2]int
		cacheSize int
	}{
		{"puts two keys in the cache with capacity of 2", 2, [][2]int{[2]int{1, 2}, [2]int{3, 4}}, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lruCache := internal.CreateCache(tc.capacity)

			for _, input := range tc.inputs {
				key, val := input[0], input[1]

				lruCache.Put(key, val)
			}

			if lruCache.Size != tc.cacheSize {
				t.Fatalf("cache size should be %v; got %v", tc.cacheSize, lruCache.Size)
			}
		})
	}
}

func TestGetCache(t *testing.T) {
	tt := []struct {
		name      string
		capacity  int
		inputs    [][2]int
		inputKey  int
		outputVal int
	}{
		{"get the value at the specified key in the cache", 2, [][2]int{[2]int{1, 2}}, 1, 2},
		{"return -1 if the specified key does not exist in the cache", 2, [][2]int{}, 1, -1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lruCache := internal.CreateCache(tc.capacity)

			for _, input := range tc.inputs {
				key, val := input[0], input[1]

				lruCache.Put(key, val)
			}

			cacheVal := lruCache.Get(tc.inputKey)

			if cacheVal != tc.outputVal {
				t.Fatalf("value at key %v should be %v; got %v", tc.inputKey, tc.outputVal, cacheVal)
			}
		})
	}
}
