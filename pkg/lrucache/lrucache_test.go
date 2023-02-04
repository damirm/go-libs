package lrucache_test

import (
	"testing"

	"github.com/damirm/go-libs/pkg/lrucache"
)

func TestBasicUsage(t *testing.T) {
	cache := lrucache.NewLRUCache[int, int](2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	_, err := cache.Get(1)
	if err != lrucache.ErrKeyNotFound {
		t.Errorf("key still exists, but it's not expected")
	}

	cache.Clear()
	cache.Put(1, 1)
	cache.Put(2, 2)
	val, err := cache.Get(1)
	if err != nil || val != 1 {
		t.Errorf("unexpected value: %d, but expected: %d, error: %#v", val, 1, err)
	}
	cache.Put(3, 3)

	_, err = cache.Get(2)
	if err != lrucache.ErrKeyNotFound {
		t.Errorf("key still exists, but it's not expected")
	}
}
