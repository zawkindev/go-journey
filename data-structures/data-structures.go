package datastructures

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
