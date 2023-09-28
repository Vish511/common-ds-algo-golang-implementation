// Stack implementation using singly linked list
package stack

import "math"

type node struct {
	value *int
	next  *node
}

type stack struct {
	first *node
	last  *node
	Len   int
}

func newNode(val int) node {
	return node{value: &val}
}

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(val int) int {
	newNode := newNode(val)
	if s.Len == 0 {
		s.first = &newNode
		s.last = &newNode
	} else {
		currFirst := s.first
		s.first = &newNode
		newNode.next = currFirst
	}
	s.Len++
	return s.Len
}

func (s *stack) Pop() (bool, int) {
	if s.Len == 0 {
		return false, math.MinInt
	}
	currFirst := s.first
	s.first = currFirst.next
	currFirst.next = nil
	s.Len--
	if s.Len == 0 {
		s.last = nil
	}
	return true, *currFirst.value
}
