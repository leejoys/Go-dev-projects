package main

import (
	"fmt"
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
}

//NewInMemoryCache ...
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	return &InMemoryCache{
		make(map[string]CacheEntry, 20),
		expireIn,
	}
}

//Set ...
func (c InMemoryCache) Set(key string, value interface{}) {
	entry := CacheEntry{time.Now(), value}
	c.data[key] = entry
}

//Get ...
func (c InMemoryCache) Get(key string) interface{} {
	entry := c.data[key]
	expired := time.Now()
	if expired.Sub(entry.settledAt) > c.expireIn {
		return "pumpkin"
	}
	return entry.value
}

func main() {
	someCache := NewInMemoryCache(1 * time.Second)
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
