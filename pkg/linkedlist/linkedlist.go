package linkedlist

import "errors"

var (
	ErrEmptyList    = errors.New("list is empty")
	ErrItemNotFound = errors.New("item not found in list")
	ErrInvalidItem  = errors.New("item can't be nil")
)

type Item[T comparable] struct {
	val  T
	next *Item[T]
	prev *Item[T]
}

func (i *Item[T]) GetValue() T {
	return i.val
}

func newItem[T comparable](val T) *Item[T] {
	return &Item[T]{val: val}
}

type LinkedList[T comparable] struct {
	head *Item[T]
	tail *Item[T]
	len  uint64
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Len() uint64 {
	return l.len
}

func (l *LinkedList[T]) Append(val T) {
	item := newItem(val)
	l.appendItem(item)
}

func (l *LinkedList[T]) appendItem(item *Item[T]) {
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		// NOTE: tail can't be nil if head is not nil.
		l.tail.next = item
		item.prev = l.tail
		l.tail = item
	}
	l.len++
}

func (l *LinkedList[T]) removeItem(item *Item[T]) (T, error) {
	if l.len == 0 {
		return *new(T), ErrEmptyList
	}
	if item == nil {
		return *new(T), ErrInvalidItem
	}
	defer func() {
		l.len--
	}()
	if item == l.head {
		l.head = item.next
	}
	if item == l.tail {
		if l.tail.prev != nil {
			l.tail.prev.next = nil
		}
		l.tail = item.prev
	} else {
		item.prev.next = item.next
	}
	return item.val, nil
}

func (l *LinkedList[T]) Pop() (T, error) {
	return l.removeItem(l.tail)
}

func (l *LinkedList[T]) GetAt(idx int) (T, error) {
	item, err := l.findByIndex(uint64(idx))
	if err != nil {
		return *new(T), err
	}
	return item.val, nil
}

func (l *LinkedList[T]) ForEach(cb func(val T, idx uint64)) {
	var idx uint64
	for item := l.head; item != nil; item = item.next {
		cb(item.val, idx)
		idx++
	}
}

func (l *LinkedList[T]) search(predicate func(item *Item[T]) bool) (*Item[T], error) {
	for item := l.head; item != nil; item = item.next {
		if predicate(item) {
			return item, nil
		}
	}
	return nil, ErrItemNotFound
}

func (l *LinkedList[T]) findByVal(val T) (*Item[T], error) {
	return l.search(func(item *Item[T]) bool {
		return item.val == val
	})
}

func (l *LinkedList[T]) findByIndex(idx uint64) (*Item[T], error) {
	var curr uint64
	return l.search(func(item *Item[T]) bool {
		if curr == idx {
			return true
		}
		curr++
		return false
	})
}

func (l *LinkedList[T]) Remove(val T) error {
	item, err := l.findByVal(val)
	if err != nil {
		return err
	}
	_, err = l.removeItem(item)
	return err
}
