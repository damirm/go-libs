package lfucache

// LFUCache - least frequently used cache.
// Cache keeps most frequently used keys,
// but least frequently used will be evicted.
type LFUCache[K comparable, V comparable] struct {
	capacity int
	cache    map[K]V
}

func (c *LFUCache[K, V]) Put(key K, val V) {
	panic("not implemented")
}

func (c *LFUCache[K, V]) Get(key K) (V, error) {
	panic("not implemented")
}
