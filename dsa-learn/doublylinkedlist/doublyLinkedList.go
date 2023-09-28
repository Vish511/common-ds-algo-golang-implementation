package doublylinkedlist

import "fmt"

type Node struct {
	Next  *Node
	Prev  *Node
	Value *int
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{}
}

func NewNode(val int) Node {
	return Node{Value: &val}
}

func (dll *DoublyLinkedList) Traverse() {
	curr := dll.Head
	for curr != nil {
		fmt.Printf(" %v ->", *curr.Value)
		curr = curr.Next
	}
}

func (dll *DoublyLinkedList) Push(val int) {
	newNode := NewNode(val)
	if dll.Head == nil {
		dll.Head = &newNode
		dll.Tail = &newNode
	} else {
		dll.Tail.Next = &newNode
		newNode.Prev = dll.Tail
		dll.Tail = &newNode
	}
	dll.Len++
}
func (dll *DoublyLinkedList) Pop() *Node {
	if dll.Len == 0 {
		return nil
	}
	currTail := dll.Tail
	dll.Tail = currTail.Prev
	if dll.Tail != nil {
		dll.Tail.Next = nil
	} else {
		//If tail is null, the list is empty
		dll.Head = nil
	}
	dll.Len--
	return currTail
}

func (dll *DoublyLinkedList) RemoveFront() *Node {
	if dll.Len == 0 {
		return nil
	}
	currHead := dll.Head
	dll.Head = currHead.Next
	dll.Len--

	if dll.Head != nil {
		dll.Head.Prev = nil
	} else {
		dll.Tail = nil
	}
	return currHead
}

func (dll *DoublyLinkedList) AddFront(val int) {
	newNode := NewNode(val)
	if dll.Len == 0 {
		dll.Head = &newNode
		dll.Tail = &newNode
	} else {
		dll.Head.Prev = &newNode
		newNode.Next = dll.Head
		dll.Head = &newNode
	}
	dll.Len++
}

func (dll *DoublyLinkedList) Get(idx int) *Node {
	if idx < 0 || idx >= dll.Len {
		return nil
	}
	var curr *Node
	if idx > (dll.Len / 2) {
		curr = dll.Tail
		currIdx := dll.Len - 1
		for currIdx > idx {
			curr = curr.Prev
			currIdx--
		}
	} else {
		curr = dll.Head
		currIdx := 0
		for currIdx < idx {
			curr = curr.Next
			currIdx++
		}

	}
	return curr
}

func (dll *DoublyLinkedList) Set(idx, val int) bool {
	if idx < 0 || idx >= dll.Len {
		return false
	}
	targetNode := dll.Get(idx)
	targetNode.Value = &val
	return true
}

func (dll *DoublyLinkedList) Insert(idx, val int) bool {
	if idx < 0 || idx > dll.Len {
		return false
	}
	if idx == 0 {
		dll.AddFront(val)
		return true
	}
	if idx == dll.Len {
		dll.Push(val)
		return true
	}
	newNode := NewNode(val)
	currPrev := dll.Get(idx - 1)
	currNext := currPrev.Next
	currPrev.Next = &newNode
	currNext.Prev = &newNode
	return true
}

func (dll *DoublyLinkedList) Remove(idx int) *Node {
	if idx < 0 || idx >= dll.Len {
		return nil
	}
	if idx == 0 {
		return dll.RemoveFront()
	}
	if idx == dll.Len-1 {
		return dll.Pop()
	}
	targetNode := dll.Get(idx)
	targetNode.Next.Prev = targetNode.Prev
	targetNode.Prev.Next = targetNode.Next
	targetNode.Next = nil
	targetNode.Prev = nil
	dll.Len--
	return targetNode
}

func (dll *DoublyLinkedList) Reverse() {
	curr := dll.Tail
	for curr != nil {
		curr.Next = curr.Prev
		curr = curr.Prev
	}
	dll.Head, dll.Tail = dll.Tail, dll.Head
}
