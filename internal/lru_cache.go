package internal

// LRUCache defines cache attributes
type LRUCache struct {
	Capacity int
}

// CreateCache creates cache with specified capacity
func CreateCache(capacity int) (LRUCache, error) {
	return LRUCache{Capacity: capacity}, nil
}
