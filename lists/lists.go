package lists

import "fmt"

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) InsertAtEnd(data interface{}) {
	node := &Node{
		next: nil,
		data: data,
	}

	if l.Tail != nil {
		currentTail := l.Tail
		currentTail.next = node
	}

	l.Tail = node
}

func (l *LinkedList) Insert(data interface{}) {
	node := &Node{
		next: nil,
		data: data,
	}

	if l.Head != nil {
		currentHead := l.Head

		if l.Tail == nil {
			l.Tail = node
		}

		node.next = currentHead
	}

	l.Head = node
}

func (l *LinkedList) Display() {
	items := l.Head

	for items != nil {
		fmt.Printf("%v ->", items.data)
		items = items.next
	}

	fmt.Println()
}
