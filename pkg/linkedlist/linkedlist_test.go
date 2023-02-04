package linkedlist_test

import (
	"testing"

	"github.com/damirm/go-libs/pkg/linkedlist"
)

func TestAppendPop(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()

	if err := list.Remove(0); err != linkedlist.ErrItemNotFound {
		t.Errorf("expected linkedlist.ErrInvalidItem, but got: %#v", err)
	}

	testValues := []int{1, 2, 3, 4, 5, 6, 7, 9}
	for _, val := range testValues {
		list.Append(val)
	}

	if list.Len() != uint64(len(testValues)) {
		t.Errorf("expected list len=%d, but got %d", len(testValues), list.Len())
	}

	list.ForEach(func(val int, idx uint64) {
		val, err := list.GetAt(int(idx))
		if val != testValues[idx] || err != nil {
			t.Errorf("failed to get element at %d: expected %d, but got %d", idx, testValues[idx], val)
		}
	})

	for i := len(testValues) - 1; i >= 0; i-- {
		res, err := list.Pop()
		if res != testValues[i] || err != nil {
			t.Errorf("got invalid list.Pop result: %d but expected %d, error: %#v", res, testValues[i], err)
		}
	}

	if list.Len() != 0 {
		t.Errorf("expected list len=0, but got %d", list.Len())
	}

	_, err := list.Pop()
	if err != linkedlist.ErrEmptyList {
		t.Errorf("expected linkedlist.ErrEmptyList error, but got: %#v", err)
	}
}

func TestRemove(t *testing.T) {
	list := linkedlist.NewLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.Remove(2)
	if err != nil {
		t.Errorf("failed to remove item, got error: %#v", err)
	}

	if list.Len() != 2 {
		t.Errorf("invalid list len: expected 2, but got %#v", err)
	}

	list.ForEach(func(val int, idx uint64) {
		if val == 2 {
			t.Errorf("got removed element")
		}
	})
}
