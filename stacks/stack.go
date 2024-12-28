package stacks

import "fmt"

type Stack struct {
	array    []interface{}
	top      int
	capacity uint
}

func NewStack(capacity uint) *Stack {

	stack := &Stack{}

	stack.top = -1
	stack.capacity = capacity
	stack.array = make([]interface{}, capacity)

	return stack
}

func (stack *Stack) IsFull() bool {
	return int(stack.capacity) == stack.top+1
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) Size() uint {
	return uint(stack.top) + 1
}

func (stack *Stack) Push(data interface{}) error {

	if stack.IsFull() {
		return fmt.Errorf("stacks overflow")
	}

	stack.top++
	stack.array[stack.top] = data

	return nil
}

func (stack *Stack) Pop() (interface{}, error) {

	if stack.IsEmpty() {
		return nil, fmt.Errorf("stacks is empty")
	}
	data := stack.array[stack.top]
	stack.top--
	stack.array = stack.array[:stack.top]

	return data, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, fmt.Errorf("stacks is empty")
	}
	data := stack.array[stack.top]

	return data, nil
}

func (stack *Stack) Clear() {
	stack.top = -1
	stack.array = nil
}
