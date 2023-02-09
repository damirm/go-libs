package linkedlist_test

import (
	"testing"

	"github.com/damirm/go-libs/pkg/linkedlist"
)

func TestPushPop(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()

	if err := list.Remove(linkedlist.NewItem(0)); err != linkedlist.ErrEmptyList {
		t.Errorf("expected linkedlist.ErrEmptyList, but got: %#v", err)
	}

	testValues := []int{1, 2, 3, 4, 5, 6, 7, 9}
	for _, val := range testValues {
		list.PushBack(linkedlist.NewItem(val))
	}

	if list.Len() != uint64(len(testValues)) {
		t.Errorf("expected list len=%d, but got %d", len(testValues), list.Len())
	}

	list.ForEach(func(item *linkedlist.Item[int], idx uint64) {
		val, err := list.GetAt(int(idx))
		if val != testValues[idx] || err != nil {
			t.Errorf("failed to get element at %d: expected %d, but got %d", idx, testValues[idx], val)
		}
	})

	for i := len(testValues) - 1; i >= 0; i-- {
		res, err := list.PopBack()
		if res == nil || res.GetValue() != testValues[i] || err != nil {
			t.Errorf("got invalid list.PopBack result: %#v but expected %d, error: %#v", res, testValues[i], err)
		}
	}

	if list.Len() != 0 {
		t.Errorf("expected list len=0, but got %d", list.Len())
	}

	_, err := list.PopBack()
	if err != linkedlist.ErrEmptyList {
		t.Errorf("expected linkedlist.ErrEmptyList error, but got: %#v", err)
	}

	testValues = []int{1, 2, 3, 4, 5, 6, 7, 9}
	for _, val := range testValues {
		list.PushFront(linkedlist.NewItem(val))
	}
	for i := len(testValues) - 1; i >= 0; i-- {
		res, err := list.PopFront()
		if res == nil || res.GetValue() != testValues[i] || err != nil {
			t.Errorf("got invalid list.PopFront result: %#v but expected %d, error: %#v", res, testValues[i], err)
		}
	}
	testValues = []int{1, 2, 3, 4, 5, 6, 7, 9}
	for _, val := range testValues {
		list.PushBack(linkedlist.NewItem(val))
	}
	for i := 0; i < len(testValues); i++ {
		res, err := list.PopFront()
		if res == nil || res.GetValue() != testValues[i] || err != nil {
			t.Errorf("got invalid list.PopFront result: %#v but expected %d, error: %#v", res, testValues[i], err)
		}
	}
}

func TestRemove(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()

	list.PushBack(linkedlist.NewItem(1))
	i2 := linkedlist.NewItem(2)
	list.PushBack(i2)
	list.PushBack(linkedlist.NewItem(3))

	err := list.Remove(i2)
	if err != nil {
		t.Errorf("failed to remove item, got error: %#v", err)
	}

	if list.Len() != 2 {
		t.Errorf("invalid list len: expected 2, but got %#v", err)
	}

	list.ForEach(func(item *linkedlist.Item[int], idx uint64) {
		if item.GetValue() == 2 {
			t.Errorf("got removed element")
		}
	})
}
