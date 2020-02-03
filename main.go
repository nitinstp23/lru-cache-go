package main

import "github.com/nitinstp23/lru-cache-go/internal"

import "fmt"

func main() {
	lruCache := internal.CreateCache(2)

	fmt.Printf("set value %v at key %v \n", 1, 1)
	lruCache.Put(1, 1)

	fmt.Printf("set value %v at key %v \n", 2, 2)
	lruCache.Put(2, 2)

	// returns 1
	fmt.Printf("value at key %v = %v \n", 1, lruCache.Get(1))

	// evicts key 2
	fmt.Printf("set value %v at key %v \n", 3, 3)
	lruCache.Put(3, 3)

	// returns -1 (not found)
	fmt.Printf("value at key %v = %v \n", 2, lruCache.Get(2))

	fmt.Printf("set value %v at key %v \n", 4, 4)
	// evicts key 1
	lruCache.Put(4, 4)

	// returns -1 (not found)
	fmt.Printf("value at key %v = %v \n", 1, lruCache.Get(1))

	// returns 3
	fmt.Printf("value at key %v = %v \n", 3, lruCache.Get(3))

	// returns 4
	fmt.Printf("value at key %v = %v \n", 4, lruCache.Get(4))
}
