package datastructures

import "fmt"

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(value any) {
	newNode := &Node{value: value}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Delete(value any) error {
	switch {
	case list.head == nil:
		return fmt.Errorf("List is empty")
	case list.head.value == value:
		list.head = list.head.next
	default:
		current := list.head
		for current.next != nil && current.next.value != value {
			current = current.next
		}
		if current.next == nil {
			return fmt.Errorf("value not found")
		} else {
			current.next = current.next.next
		}
	}
	return nil
}
