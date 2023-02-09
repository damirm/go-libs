package lrucache

import (
	"container/list"
	"errors"
)

var (
	ErrKeyNotFound     = errors.New("key not found")
	ErrInvalidCapacity = errors.New("capacity must be > 1")
)

type cachedItem[K comparable, V any] struct {
	key   K
	value V
	el    *list.Element
}

// LRUCache - least recently used cache.
// Cache keeps most recently used keys,
// but least recently used can be evicted.
type LRUCache[K comparable, V any] struct {
	capacity uint
	cache    map[K]cachedItem[K, V]

	// least recently used keys always in front of list.
	keys *list.List
}

func NewLRUCache[K comparable, V any](capacity uint) (*LRUCache[K, V], error) {
	if capacity == 0 {
		return nil, ErrInvalidCapacity
	}
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]cachedItem[K, V]),
		keys:     list.New(),
	}, nil
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if ci, ok := c.cache[key]; ok {
		c.keys.MoveToFront(ci.el)
	} else {
		el := c.keys.PushFront(key)
		ci = cachedItem[K, V]{
			key:   key,
			value: value,
			el:    el,
		}
		c.cache[key] = ci
		if uint(c.keys.Len()) > c.capacity {
			c.evict()
		}
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, error) {
	if ci, ok := c.cache[key]; ok {
		c.keys.MoveToFront(ci.el)
		return ci.value, nil
	}
	return *new(V), ErrKeyNotFound
}

func (c *LRUCache[K, V]) Clear() {
	c.cache = make(map[K]cachedItem[K, V])
	c.keys.Init()
}

func (c *LRUCache[K, V]) evict() {
	tail := c.keys.Back()
	if tail != nil {
		delete(c.cache, tail.Value.(K))
		c.keys.Remove(tail)
	}
}
