package internal

// LRUCache defines cache attributes
type LRUCache struct {
	Capacity int
	Size     int
	store    map[int]*cacheItem
	head     *cacheItem
	tail     *cacheItem
}

type cacheItem struct {
	key   int
	value int
	prev  *cacheItem
	next  *cacheItem
}

// CreateCache creates cache with specified capacity
func CreateCache(capacity int) LRUCache {
	lruCache := LRUCache{Capacity: capacity, store: map[int]*cacheItem{}}

	head := cacheItem{prev: nil}
	tail := cacheItem{next: nil}

	head.next = &tail
	tail.prev = &head

	lruCache.head = &head
	lruCache.tail = &tail

	return lruCache
}

// Put the specified key and value to the cache
func (lruCache *LRUCache) Put(key, val int) {
	if lruCache.hasKey(key) {
		// if key exists in store
		// update the value at the specified key
		// move it to front of the list

		item := lruCache.store[key]
		item.value = val

		lruCache.moveToFront(item)
	} else {
		// if key doesn't exist in store
		// write it to the store
		// move it to the front
		// if size exceeds capacity, remove last cache item
		// increment current size

		newItem := &cacheItem{key: key, value: val}
		newItem = lruCache.addToFront(newItem)

		// fmt.Printf("adding item %v \n", newItem)
		lruCache.store[key] = newItem
		lruCache.Size++

		if lruCache.Size > lruCache.Capacity {
			// fmt.Println("removing last item")
			lastItem := lruCache.removeLastItem()
			// fmt.Printf("deleting last item %v \n", lastItem)
			delete(lruCache.store, lastItem.key)
		}
	}
}

// Get the value at the specified key
func (lruCache *LRUCache) Get(key int) int {
	if !lruCache.hasKey(key) {
		return -1
	}

	item := lruCache.store[key]
	lruCache.moveToFront(item)

	return item.value
}

func (lruCache *LRUCache) hasKey(key int) bool {
	if _, ok := lruCache.store[key]; ok {
		return true
	}

	return false
}

func (lruCache *LRUCache) addToFront(item *cacheItem) *cacheItem {
	firstItem := lruCache.head.next

	firstItem.prev = item
	item.next = firstItem
	item.prev = lruCache.head
	lruCache.head.next = item

	return item
}

func (lruCache *LRUCache) moveToFront(item *cacheItem) {
	lruCache.removeFromList(item)
	lruCache.addToFront(item)
}

func (lruCache *LRUCache) removeFromList(item *cacheItem) {
	prevItem := item.prev
	nextItem := item.next

	prevItem.next = nextItem
	nextItem.prev = prevItem
}

func (lruCache *LRUCache) removeLastItem() *cacheItem {
	lastItem := lruCache.tail.prev

	lruCache.tail.prev = lastItem.prev
	lruCache.Size--

	return lastItem
}
