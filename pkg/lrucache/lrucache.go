package lrucache

import "errors"

var (
	ErrKeyNotFound     = errors.New("key not found")
	ErrInvalidCapacity = errors.New("capacity must be > 1")
)

// LRUCache - least recently used cache.
// Cache keeps most recently used keys,
// but least recently used can be evicted.
type LRUCache[K comparable, V any] struct {
	capacity uint
	cache    map[K]V
	usages   map[K]uint
}

func NewLRUCache[K comparable, V any](capacity uint) (*LRUCache[K, V], error) {
	if capacity == 0 {
		return nil, ErrInvalidCapacity
	}
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]V, capacity),
		usages:   make(map[K]uint, capacity),
	}, nil
}

var nowCounter uint

func now() uint {
	nowCounter++
	return nowCounter
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if uint(len(c.cache)) == c.capacity {
		c.evict()
	}
	c.cache[key] = value
	c.usages[key] = now()
}

func (c *LRUCache[K, V]) Get(key K) (V, error) {
	val, ok := c.cache[key]
	if !ok {
		return *new(V), ErrKeyNotFound
	}
	c.usages[key] = now()
	return val, nil
}

func (c *LRUCache[K, V]) Clear() {
	c.cache = make(map[K]V, c.capacity)
	c.usages = make(map[K]uint, c.capacity)
}

// TODO: How to speedup eviction of least recently used keys?
func (c *LRUCache[K, V]) evict() {
	var oldest = c.firstUsageKey()
	for k, n := range c.usages {
		if n < c.usages[oldest] {
			oldest = k
		}
	}
	delete(c.usages, oldest)
	delete(c.cache, oldest)
}

func (c *LRUCache[K, V]) firstUsageKey() K {
	for k := range c.usages {
		return k
	}
	panic("usages is empty")
}
