package queue

import "math"

type node struct {
	value *int
	next  *node
}

type queue struct {
	first *node
	last  *node
	Len   int
}

func NewQueue() queue {
	return queue{}
}

func newNode(val int) node {
	return node{value: &val}
}

func (q *queue) Enqueue(val int) int {
	newNode := newNode(val)
	if q.Len == 0 {
		q.first = &newNode
		q.last = &newNode
	} else {
		q.last.next = &newNode
		q.last = &newNode
	}
	q.Len++
	return q.Len
}

func (q *queue) Dequeue() (bool, int) {
	if q.Len == 0 {
		return false, math.MinInt
	}
	oldFirst := q.first
	q.first = oldFirst.next
	q.Len--
	if q.Len == 0 {
		q.last = nil
	}
	return true, *oldFirst.value
}
