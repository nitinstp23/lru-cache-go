package internal_test

import (
	"github.com/nitinstp23/lru-cache-go/internal"
	"testing"
)

func TestCreateLRUCache(t *testing.T) {
	tt := []struct {
		name        string
		capacity    int
		cacheOutput internal.LRUCache
		errorText   string
	}{
		{"create empty cache with capacity of 2 elements", 2, internal.LRUCache{Capacity: 2}, ""},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			lruCache, err := internal.CreateCache(tc.capacity)

			if err != nil && err.Error() != tc.errorText {
				t.Fatalf("error text should be %v; got %v", err, tc.errorText)
			}

			if lruCache.Capacity != tc.capacity {
				t.Fatalf("cache capacity should be %v; got %v", tc.capacity, lruCache.Capacity)
			}
		})
	}
}
