package stacks

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	data interface{}
}

type LlStack struct {
	top  *Node
	size int
}

func NewLlStack() *LlStack {

	stack := &LlStack{
		size: 0,
	}

	return stack
}

func (stack *LlStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LlStack) Size() int {
	return stack.size
}

func (stack *LlStack) Push(data interface{}) {
	newNode := &Node{
		next: stack.top,
		data: data,
	}
	stack.top = newNode
	stack.size++
}

func (stack *LlStack) Pop() (interface{}, error) {

	if stack.IsEmpty() {
		return nil, errors.New("stacks is empty")
	}

	data := stack.top.data

	stack.top = stack.top.next
	stack.size--

	return data, nil
}

func (stack *LlStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("stacks is empty")
	}
	data := stack.top.data
	return data, nil
}
