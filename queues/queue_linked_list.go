package queues

import "errors"

type Node struct {
	data interface{}
	next *Node
}

type LlQueue struct {
	front *Node
	rear  *Node
	size  int
}

func (queue *LlQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *LlQueue) Size() int {
	return queue.size
}

func (queue *LlQueue) Enqueue(data interface{}) {

	newNode := &Node{
		data: data,
		next: nil,
	}

	if queue.IsEmpty() {
		queue.front = newNode
	}

	queue.rear = newNode
	queue.size++

}

func (queue *LlQueue) Dequeue() (interface{}, error) {

	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	data := queue.front.data

	queue.front.next = nil
	queue.size--

	return data, nil
}

func (queue *LlQueue) Front() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	data := queue.front.data

	return data, nil
}

func (queue *LlQueue) Back() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return queue.rear.data, nil
}
