package stacks

import (
	"testing"
)

func TestStack(t *testing.T) {

	stack := NewStack(5)

	t.Run("check if a stacks is empty", func(t *testing.T) {

		got := stack.IsEmpty()

		if got != true {
			t.Errorf("got %t want %t", true, got)
		}
	})

	t.Run("should be able to push elements into stacks that is not full", func(t *testing.T) {

		err := stack.Push(2)

		if err != nil {
			t.Errorf("push operation failed")
		}
		want := uint(1)

		got := stack.Size()

		if got != want {
			t.Errorf("check size: got %d want %d", got, want)
		}

		if stack.top != 0 {
			t.Errorf("check index: the index of the first element inserted into stacks should be 0 but got %d", stack.top)
		}
	})

	t.Run("should be able to pop elements from stacks", func(t *testing.T) {

		err := stack.Push(2)

		if err != nil {
			t.Errorf("push operation failed")
		}

		got, err := stack.Pop()

		want := 2

		if err != nil {
			t.Errorf("pop operation failed")
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		expected := uint(1)

		actual := stack.Size()

		if actual != expected {
			t.Errorf("check size: got %d want %d", actual, expected)
		}

	})

}
