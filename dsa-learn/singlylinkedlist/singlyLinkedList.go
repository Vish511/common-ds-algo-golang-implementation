package singlylinkedlist

import "fmt"

type Node struct {
	Next  *Node
	Value *int
}

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewNode(val int) Node {
	newNode := Node{nil, &val}
	return newNode
}

func NewSinglyLinkedList() SinglyLinkedList {
	newList := SinglyLinkedList{nil, nil, 0}
	return newList
}

func (sll *SinglyLinkedList) Traverse() {
	var curr = sll.Head
	for curr != nil {
		fmt.Printf("%v->", *curr.Value)
		curr = curr.Next
	}
}

func (sll *SinglyLinkedList) Push(val int) {
	node := NewNode(val)
	if sll.Len == 0 {
		sll.Head = &node
		sll.Tail = &node
	} else {
		sll.Tail.Next = &node
		sll.Tail = &node
	}
	sll.Len++
}

func (sll *SinglyLinkedList) Pop() *Node {
	if sll.Len == 0 {
		return nil
	}
	curr := sll.Head
	newTail := sll.Head
	for curr.Next != nil {
		newTail = curr
		curr = curr.Next
	}
	val := curr
	newTail.Next = nil
	sll.Tail = newTail
	sll.Len--

	if sll.Len == 0 {
		sll.Head = nil
		sll.Tail = nil
	}
	return val
}

func (sll *SinglyLinkedList) RemoveFront() *Node {
	if sll.Len == 0 {
		return nil
	}
	currHead := sll.Head
	sll.Head = currHead.Next //New Head
	sll.Len--
	if sll.Len == 0 {
		sll.Head = nil
		sll.Tail = nil
	}
	return currHead
}

func (sll *SinglyLinkedList) AddFront(val int) {
	node := NewNode(val)
	currHead := sll.Head
	sll.Head = &node
	sll.Head.Next = currHead
	sll.Len++
}

func (sll *SinglyLinkedList) Insert(idx, val int) bool {
	if idx < 0 || idx > sll.Len {
		return false
	}
	if idx == 0 {
		sll.AddFront(val)
		return true
	}
	if idx == sll.Len {
		sll.Push(val)
		return true
	}
	newNode := NewNode(val)
	curr := sll.Get(idx - 1)
	next := curr.Next
	curr.Next = &newNode
	newNode.Next = next
	sll.Len++
	return true
}

func (sll *SinglyLinkedList) Remove(idx int) *Node {
	if idx < 0 || idx >= sll.Len {
		return nil
	}
	if idx == 0 {
		return sll.RemoveFront()
	}
	if idx == sll.Len-1 {
		return sll.Pop()
	}
	currPrev := sll.Get(idx - 1) //Get the node before the node to be removed
	currNext := sll.Get(idx + 1) //Get the node next to the one to be removed

	removedNode := currPrev.Next
	removedNode.Next = nil
	currPrev.Next = currNext
	sll.Len--
	return removedNode
}

func (sll *SinglyLinkedList) Get(idx int) *Node {
	if idx < 0 || idx >= sll.Len {
		return nil
	}
	curr := sll.Head
	curIdx := 0
	for curIdx < idx {
		curIdx++
		curr = curr.Next
	}
	return curr
}

func (sll *SinglyLinkedList) Set(idx, val int) bool {
	if idx < 0 || idx >= sll.Len {
		return false
	}
	targetNode := sll.Get(idx)
	targetNode.Value = &val
	return true
}

func (sll *SinglyLinkedList) Reverse() {
	var curr, node, prev *Node = sll.Head, sll.Head, nil
	for node != nil {
		node = node.Next
		curr.Next = prev
		prev = curr
		curr = node
	}
	sll.Head, sll.Tail = sll.Tail, sll.Head
}
