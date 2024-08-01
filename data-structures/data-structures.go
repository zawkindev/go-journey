package datastructures

import "fmt"

//  --------------- LINKED LIST ------------

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

//  --------------- STACK ------------

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() error {
	if size := len(s.items); size == 0 {
		return fmt.Errorf("stack is empty!")
	}
	s.items = s.items[:len(s.items)-1]
	return nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty!")
	}
	return s.items[0], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}
