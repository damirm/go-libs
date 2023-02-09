package lrucache

import (
	"errors"

	"github.com/damirm/go-libs/pkg/linkedlist"
)

var (
	ErrKeyNotFound     = errors.New("key not found")
	ErrInvalidCapacity = errors.New("capacity must be > 1")
)

type cachedItem[K comparable, V any] struct {
	key   K
	value V
	item  *linkedlist.Item[K]
}

// LRUCache - least recently used cache.
// Cache keeps most recently used keys,
// but least recently used can be evicted.
type LRUCache[K comparable, V any] struct {
	capacity uint
	values   map[K]cachedItem[K, V]
	keys     *linkedlist.LinkedList[K]
}

func NewLRUCache[K comparable, V any](capacity uint) (*LRUCache[K, V], error) {
	if capacity == 0 {
		return nil, ErrInvalidCapacity
	}
	return &LRUCache[K, V]{
		capacity: capacity,
		values:   make(map[K]cachedItem[K, V]),
		keys:     linkedlist.NewLinkedList[K](),
	}, nil
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if ci, ok := c.values[key]; ok {
		if c.keys.Front() != ci.item {
			c.keys.Remove(ci.item)
			c.keys.PushFront(ci.item)
		}
	} else {
		item := linkedlist.NewItem(key)
		ci = cachedItem[K, V]{
			key:   key,
			value: value,
			item:  item,
		}
		c.values[key] = ci
		c.keys.PushFront(item)
	}
	if uint(c.keys.Len()) > c.capacity {
		c.evict()
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, error) {
	if ci, ok := c.values[key]; ok {
		c.keys.Remove(ci.item)
		c.keys.PushFront(ci.item)
		return ci.value, nil
	}
	return *new(V), ErrKeyNotFound
}

func (c *LRUCache[K, V]) Clear() {
	c.values = make(map[K]cachedItem[K, V])
	c.keys.Clear()
}

func (c *LRUCache[K, V]) evict() {
	k, _ := c.keys.PopBack()
	if k != nil {
		delete(c.values, k.GetValue())
	}
}
