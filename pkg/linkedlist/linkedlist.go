package linkedlist

import "errors"

var (
	ErrEmptyList    = errors.New("list is empty")
	ErrItemNotFound = errors.New("item not found in list")
	ErrInvalidItem  = errors.New("item can't be nil")
)

type Item[T any] struct {
	val  T
	next *Item[T]
	prev *Item[T]
}

func (i *Item[T]) GetValue() T {
	return i.val
}

func NewItem[T any](val T) *Item[T] {
	return &Item[T]{val: val}
}

type LinkedList[T any] struct {
	head *Item[T]
	tail *Item[T]
	len  uint64
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Len() uint64 {
	return l.len
}

func (l *LinkedList[T]) PushFront(item *Item[T]) {
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		l.head.prev = item
		item.next = l.head
		l.head = item
	}
	l.len++
}

func (l *LinkedList[T]) PushBack(item *Item[T]) {
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		l.tail.next = item
		item.prev = l.tail
		l.tail = item
	}
	l.len++
}

func (l *LinkedList[T]) removeItem(item *Item[T]) error {
	if l.len == 0 {
		return ErrEmptyList
	}
	if item == nil {
		return ErrInvalidItem
	}
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
	l.len--
	item.prev = nil
	item.next = nil
	return nil
}

func (l *LinkedList[T]) PopFront() (*Item[T], error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	item := l.head
	if l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next
	}
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}
	return item, nil
}

func (l *LinkedList[T]) PopBack() (*Item[T], error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}
	item := l.tail
	if l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}
	return item, nil
}

func (l *LinkedList[T]) Front() *Item[T] {
	return l.head
}

func (l *LinkedList[T]) Back() *Item[T] {
	return l.tail
}

func (l *LinkedList[T]) GetAt(idx int) (T, error) {
	item, err := l.findByIndex(uint64(idx))
	if err != nil {
		return *new(T), err
	}
	return item.val, nil
}

func (l *LinkedList[T]) ForEach(cb func(item *Item[T], idx uint64)) {
	var idx uint64
	for item := l.head; item != nil; item = item.next {
		cb(item, idx)
		idx++
	}
}

func (l *LinkedList[T]) Search(predicate func(item *Item[T]) bool) (*Item[T], error) {
	for item := l.head; item != nil; item = item.next {
		if predicate(item) {
			return item, nil
		}
	}
	return nil, ErrItemNotFound
}

func (l *LinkedList[T]) findByIndex(idx uint64) (*Item[T], error) {
	var curr uint64
	return l.Search(func(item *Item[T]) bool {
		if curr == idx {
			return true
		}
		curr++
		return false
	})
}

func (l *LinkedList[T]) Remove(item *Item[T]) error {
	return l.removeItem(item)
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}
