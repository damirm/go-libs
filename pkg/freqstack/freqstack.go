package freqstack

import (
	"container/list"
	"errors"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type FreqStack[T comparable] struct {
	maxFreq uint64
	freqs   map[T]uint64
	lists   map[uint64]*list.List
}

func NewFreqStack[T comparable]() *FreqStack[T] {
	return &FreqStack[T]{
		freqs: make(map[T]uint64),
		lists: make(map[uint64]*list.List),
	}
}

func (s *FreqStack[T]) Push(val T) {
	s.freqs[val]++
	freq := s.freqs[val]
	if _, ok := s.lists[freq]; !ok {
		s.lists[freq] = list.New()
	}
	s.lists[freq].PushBack(val)
	if freq > s.maxFreq {
		s.maxFreq = freq
	}
}

func (s *FreqStack[T]) Pop() (T, error) {
	if len(s.lists) == 0 {
		return *new(T), ErrEmptyStack
	}
	stack := s.lists[s.maxFreq]
	item := stack.Back()
	stack.Remove(item)
	if stack.Len() == 0 {
		delete(s.lists, s.maxFreq)
		s.maxFreq--
	}
	val := item.Value.(T)
	s.freqs[val]--
	if s.freqs[val] == 0 {
		delete(s.freqs, val)
	}
	return val, nil
}
