package lrucache

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
)

type LRUCache[K comparable, V comparable] struct {
	capacity int
	cache    map[K]V
	usages   map[K]uint
	oldest   K
}

func NewLRUCache[K comparable, V comparable](capacity int) *LRUCache[K, V] {
	if capacity < 1 {
		capacity = 1
	}
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]V, capacity),
		usages:   make(map[K]uint, capacity),
	}
}

var nowCounter uint

func now() uint {
	nowCounter++
	return nowCounter
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if len(c.cache) == c.capacity {
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

// TODO: How to speedup eviction of least frequently used keys?
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
