package queues

import "testing"

func TestQueue(t *testing.T) {

	//queue := &LlQueue{}

	queue := NewQueue(5)

	queue.Enqueue(2)
	queue.Enqueue(4)

	want := 2

	got, err := queue.Front()

	if err != nil {
		t.Errorf("enqueuing failed")
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	if queue.size != 2 {
		t.Errorf("expect the size of queue to be 2 after enqueuing two values")
	}

	want = 4

	got, _ = queue.Back()

	if got != want {
		t.Errorf("expect the last element in queue to be %d got %d", want, got)
	}

}
