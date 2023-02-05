package ringbuffer

import "errors"

var (
	ErrEmptyBuffer  = errors.New("buffer is empty")
	ErrBufferIsFull = errors.New("buffer is full")
)

type RingBuffer[T comparable] struct {
	capacity   uint
	head, tail uint
	items      []T
	full       bool
}

func NewRingBuffer[T comparable](capacity uint) *RingBuffer[T] {
	return &RingBuffer[T]{
		capacity: capacity,
		items:    make([]T, capacity),
	}
}

func (b *RingBuffer[T]) Put(value T) error {
	if b.full {
		return ErrBufferIsFull
	}
	b.items[b.tail] = value
	b.tail++
	b.tail = b.tail % b.capacity
	if b.tail == b.head {
		b.full = true
	}
	return nil
}

func (b *RingBuffer[T]) IsFull() bool {
	return b.full
}

func (b *RingBuffer[T]) Size() uint {
	if b.full {
		return b.capacity
	}
	if b.tail >= b.head {
		return b.tail - b.head
	}
	return b.capacity + b.tail - b.head
}

func (b *RingBuffer[T]) Get() (T, error) {
	if b.head == b.tail && !b.full {
		return *new(T), ErrEmptyBuffer
	}
	res := b.items[b.head]
	b.head++
	b.full = false
	return res, nil
}
