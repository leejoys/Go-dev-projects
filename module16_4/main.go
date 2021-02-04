package main

import (
	"fmt"
	"sync"
	"time"
)

// это трюк для проверки типа: до тех пор пока InMemoryCache не будет
// реализовывать интерфейс Cache, программа не запустится
var _ Cache = InMemoryCache{}

//CacheEntry ...
type CacheEntry struct {
	settledAt time.Time
	value     interface{}
}

//Cache ...
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

//InMemoryCache ...
type InMemoryCache struct {
	data     map[string]CacheEntry
	expireIn time.Duration
	m        *sync.RWMutex
}

//NewInMemoryCache ...
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		make(map[string]CacheEntry, 20),
		expireIn,
		&sync.RWMutex{},
	}
}

//Set ...
func (c InMemoryCache) Set(key string, value interface{}) {
	entry := CacheEntry{time.Now(), value}
	c.Lock()
	c.data[key] = entry
	c.Unlock()
}

//Get ...
func (c InMemoryCache) Get(key string) interface{} {
	defer c.RUnlock()
	c.RLock()
	entry, ok := c.data[key]

	if !ok {
		fmt.Println("There is no value")
		return nil
	}
	expired := time.Now()
	if expired.Sub(entry.settledAt) > c.expireIn {
		delete(c.data, key)
		fmt.Println("Expired")
		return nil
	}
	return entry.value
}

func main() {
	someCache := NewInMemoryCache(1 * time.Second)
	fmt.Println(someCache.Get("one"))
	someCache.Set("one", 1)
	someCache.Set("two", 2)
	someCache.Set("three", 3)
	someCache.Set("four", 4)
	fmt.Println(someCache.Get("one"))
	fmt.Println(someCache.Get("two"))
	time.Sleep(1 * time.Second)
	fmt.Println(someCache.Get("three"))
	fmt.Println(someCache.Get("four"))
	someCache.Set("five", 5)
	fmt.Println(someCache.Get("five"))
	someCache.Set("one", 1)
	someCache.Set("two", 2)
	time.Sleep(1 * time.Second)
	someCache.Set("three", 3)
	someCache.Set("four", 4)
	fmt.Println(someCache.Get("one"))
	fmt.Println(someCache.Get("two"))
	fmt.Println(someCache.Get("three"))
	fmt.Println(someCache.Get("four"))
}
