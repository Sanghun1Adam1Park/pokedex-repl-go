package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
	done     chan struct{}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		done:     make(chan struct{}),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return append([]byte(nil), entry.val...), true
}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			now := time.Now()
			expired := make([]string, 0)
			c.mu.RLock()
			for key, val := range c.entries {
				if now.Sub(val.createdAt) > c.interval {
					expired = append(expired, key)
				}
			}
			c.mu.RUnlock()

			if len(expired) > 0 {
				c.mu.Lock()
				for _, key := range expired {
					delete(c.entries, key)
				}
				c.mu.Unlock()
			}

		case <-c.done:
			return
		}
	}
}

func (c *Cache) Close() {
	close(c.done)
}
