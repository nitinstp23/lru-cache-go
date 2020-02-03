package internal

// LRUCache defines cache attributes
type LRUCache struct {
	Capacity int
	Store    map[int]int
	Size     int
}

// CreateCache creates cache with specified capacity
func CreateCache(capacity int) (LRUCache, error) {
	lruCache := LRUCache{Capacity: capacity, Store: map[int]int{}}

	return lruCache, nil
}

// Put sets the specified key and value to the cache
func (lruCache *LRUCache) Put(key, val int) error {
	lruCache.Store[key] = val

	lruCache.Size++

	return nil
}
