package linked_list

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node // The head of the linked list points to the first element in the list
	size int
}

func (ll *LinkedList) Traverse() {

	if ll.head == nil {
		fmt.Println("Empty list")
	}

	current := ll.head

	for current != nil {
		fmt.Printf("%v---->", current.data)
		current = current.next
	}
}

// Length O(1) for creating a temp variable to hold size
//func (ll *LinkedList) Length() int {
//	return ll.size
//}

func (ll *LinkedList) Length() int {

	size := 0

	current := ll.head

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (ll *LinkedList) InsertBeginning(data interface{}) {

	node := &Node{
		data: data,
	}

	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
}

func (ll *LinkedList) InsertEnd(data interface{}) {

	node := &Node{
		data: data,
	}

	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	ll.size++

}

func (ll *LinkedList) Insert(position int, data interface{}) error {

	// Allow insertion at the beginning and end of the list
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("insert: Index out of bounds")
	}

	newNode := &Node{
		data: data,
	}

	var prev, current *Node

	prev = nil
	current = ll.head

	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}

	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}
	ll.size++

	return nil

}

func (ll *LinkedList) DeleteFirst() (interface{}, error) {

	if ll.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	data := ll.head.data

	ll.head = ll.head.next

	ll.size--

	return data, nil

}

func (ll *LinkedList) DeleteLast() (interface{}, error) {

	if ll.head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	current := ll.head

	var prev *Node

	for current.next != nil {
		prev = current
		current = current.next
	}

	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}

	ll.size--
	return current.data, nil
}

func (ll *LinkedList) Delete(position int) (interface{}, error) {

	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("insert: Index out of bounds")
	}
	var prev, current *Node
	prev = nil
	current = ll.head
	pos := 0
	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}
		if current != nil {
			prev.next = current.next
		}
	}
	ll.size--
	return current.data, nil
}
