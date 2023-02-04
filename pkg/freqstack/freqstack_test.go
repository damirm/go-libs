package freqstack_test

import (
	"testing"

	"github.com/damirm/go-libs/pkg/freqstack"
)

func TestBasicUsage(t *testing.T) {
	for _, tc := range []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{3, 2, 1},
		},
		{
			input:  []int{1, 2, 2, 3},
			output: []int{2, 3, 2, 1},
		},
		{
			input:  []int{1, 2, 2, 3, 3, 1, 1, 4},
			output: []int{1, 1, 3, 2, 4, 3, 2, 1},
		},
	} {
		stack := freqstack.NewFreqStack[int]()
		for _, n := range tc.input {
			stack.Push(n)
		}

		for i, n := range tc.output {
			val, err := stack.Pop()
			if err != nil || val != n {
				t.Errorf("expected %d but got %d:%d, err: %#v", n, i, val, err)
			}
		}
	}
}

func TestInvalidUsage(t *testing.T) {
	stack := freqstack.NewFreqStack[int]()
	_, err := stack.Pop()
	if err != freqstack.ErrEmptyStack {
		t.Errorf("expected freqstack.ErrEmptyStack, but got: %#v", err)
	}
}
