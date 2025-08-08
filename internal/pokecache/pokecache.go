package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(d time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(d)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > d {
				delete(c.cache, k)
			}
		}
		c.mux.Unlock()
	}
}
