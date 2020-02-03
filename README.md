### LRU Cache

Design and implement a data structure for Least Recently Used (LRU) cache.
It should support the following operations: get and put.

* `get(key)` - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.

* `put(key, value)` - Set or insert the value if the key is not already present.
When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

**Example:**

<!-- create a cache with capacity of 2 items -->
```go
lruCache := internal.CreateCache(2)

lruCache.Put(1, 1);
lruCache.Put(2, 2);
lruCache.Get(1);       // returns 1
lruCache.Put(3, 3);    // evicts key 2
lruCache.Get(2);       // returns -1 (not found)
lruCache.Put(4, 4);    // evicts key 1
lruCache.Get(1);       // returns -1 (not found)
lruCache.Get(3);       // returns 3
lruCache.Get(4);       // returns 4
```
