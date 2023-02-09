package freqstack

import (
	"errors"

	"github.com/damirm/go-libs/pkg/linkedlist"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type FreqStack[T comparable] struct {
	maxFreq uint64
	freqs   map[T]uint64
	// TODO: Use container/list instead.
	stacks map[uint64]*linkedlist.LinkedList[T]
}

func NewFreqStack[T comparable]() *FreqStack[T] {
	return &FreqStack[T]{
		freqs:  make(map[T]uint64),
		stacks: make(map[uint64]*linkedlist.LinkedList[T]),
	}
}

func (s *FreqStack[T]) Push(val T) {
	s.freqs[val]++
	freq := s.freqs[val]
	if _, ok := s.stacks[freq]; !ok {
		s.stacks[freq] = linkedlist.NewLinkedList[T]()
	}
	item := linkedlist.NewItem(val)
	s.stacks[freq].PushBack(item)
	if freq > s.maxFreq {
		s.maxFreq = freq
	}
}

func (s *FreqStack[T]) Pop() (T, error) {
	if len(s.stacks) == 0 {
		return *new(T), ErrEmptyStack
	}
	stack := s.stacks[s.maxFreq]
	item, err := stack.PopBack()
	if err != nil {
		return *new(T), err
	}
	if stack.Len() == 0 {
		delete(s.stacks, s.maxFreq)
		s.maxFreq--
	}
	val := item.GetValue()
	s.freqs[val]--
	if s.freqs[val] == 0 {
		delete(s.freqs, val)
	}
	return val, nil
}
