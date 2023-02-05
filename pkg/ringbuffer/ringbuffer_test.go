package ringbuffer_test

import (
	"testing"

	"github.com/damirm/go-libs/pkg/ringbuffer"
)

func TestBasicUsage(t *testing.T) {
	buf := ringbuffer.NewRingBuffer[int](3)
	_, err := buf.Get()
	if err != ringbuffer.ErrEmptyBuffer {
		t.Errorf("expected ringbuffer.ErrEmptyBuffer, but got: %#v", err)
	}

	if buf.IsFull() {
		t.Error("expected empty buffer, but it is full")
	}

	if buf.Size() != 0 {
		t.Errorf("expected empty buffer, but got: %d", buf.Size())
	}

	if err := buf.Put(1); err != nil {
		t.Errorf("failed to put element to buffer: %#v", err)
	}
	if buf.Size() != 1 {
		t.Errorf("invalid buffer size, expected 1, but got: %d", buf.Size())
	}
	if err := buf.Put(2); err != nil {
		t.Errorf("failed to put element to buffer: %#v", err)
	}
	if buf.Size() != 2 {
		t.Errorf("invalid buffer size, expected 2, but got: %d", buf.Size())
	}
	if err := buf.Put(3); err != nil {
		t.Errorf("failed to put element to buffer: %#v", err)
	}
	if buf.Size() != 3 {
		t.Errorf("invalid buffer size, expected 3, but got: %d", buf.Size())
	}
	if !buf.IsFull() {
		t.Error("expected full buffer, but it is not")
	}
	if err := buf.Put(4); err == nil {
		t.Error("expected error while pushing new element, because buffer is full")
	}

	if val, err := buf.Get(); val != 1 || err != nil {
		t.Errorf("invalid value, expected 1, but got: %d, err: %#v", val, err)
	}
	if buf.Size() != 2 {
		t.Errorf("invalid buffer size, expected 2, but got: %d", buf.Size())
	}
	if err := buf.Put(4); err != nil {
		t.Errorf("failed to put element to buffer: %#v", err)
	}
	if val, err := buf.Get(); val != 2 || err != nil {
		t.Errorf("invalid value, expected 2, but got: %d, err: %#v", val, err)
	}
}
