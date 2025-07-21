package pokecache

import (
	"sync"
	"time"
)

type cacheEntry[T any] struct {
	createdAt time.Time
	val       T
}
type Cache[T any] struct {
	dict map[string]cacheEntry[T]
	mu   *sync.RWMutex
}

func NewCache(interval time.Duration) Cache[any] {
	cache := Cache[any]{
		dict: make(map[string]cacheEntry[any]),
		mu:   &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache[T]) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache[T]) reap(interval time.Duration) {
	timePassed := time.Now().UTC().Add(-interval)
	for key, entry := range c.dict {
		if entry.createdAt.Before(timePassed) {
			delete(c.dict, key)
		}
	}
}

func (c *Cache[T]) Add(key string, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dict[key] = cacheEntry[T]{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.dict[key]
	return entry.val, ok
}