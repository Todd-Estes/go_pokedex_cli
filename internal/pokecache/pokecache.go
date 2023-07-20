package pokecache

import (
	"time"
	"sync"
	"fmt"
)

type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

type Cache struct {
	Cache map[string]cacheEntry
	Mux   *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{ 
		Cache: make(map[string]cacheEntry),
		Mux:   &sync.Mutex{},
	}

	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		Val: val,
	}
	fmt.Println("NEW ENTRY ADDED!!!")
	c.Cache[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mux.Lock()
	c.Mux.Unlock()
	entry, ok := c.Cache[key]
	if ok {
		return entry.Val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration ) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for k, v := range c.Cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Cache, k)
		}
	}
}



	// for {
	// 	for key, cacheEntry := range c.cache {
	// 		if interval < cacheEntry.createdAt.Sub(time.Now()) {
	// 			delete(c.cache, key)
	// 		}
	// 	}
	// }
// }