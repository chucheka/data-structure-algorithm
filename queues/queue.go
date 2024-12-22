package queues

import (
	"bytes"
	"errors"
	"fmt"
)

type Queue struct {
	size     int
	capacity int
	front    int
	rear     int
	array    []interface{}
}

func NewQueue(capacity int) *Queue {

	queue := &Queue{
		size: 0,
	}

	queue.array = make([]interface{}, capacity)
	queue.front, queue.capacity = 0, capacity

	return queue
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) IsFull() bool {
	return queue.size == queue.capacity
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Enqueue(data interface{}) error {
	if queue.IsFull() {
		return errors.New("queue is full")
	}

	queue.rear = (queue.front + queue.size) % queue.capacity

	queue.array[queue.rear] = data
	queue.size++

	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {

	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	data := queue.array[queue.front]

	queue.front = (queue.front + 1) % queue.capacity

	queue.size--

	return data, nil
}

func (queue *Queue) Front() (interface{}, error) {

	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return queue.array[queue.front], nil
}

func (queue *Queue) Back() (interface{}, error) {
	return queue.array[queue.rear], nil
}

func (queue *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := queue.front
	for i := 0; i < queue.size; i++ {
		result.WriteString(fmt.Sprintf("%v", queue.array[j]))
		if i < queue.size-1 {
			result.WriteByte(' ')
		}
		j = (j + 1) % queue.capacity
	}
	result.WriteByte(']')
	return result.String()
}
