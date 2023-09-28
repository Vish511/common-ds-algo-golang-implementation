package priorityqueue

import "math"

type node struct {
	val      string
	priority int
}

type priorityQueue struct {
	values []node
}

func newNode(val string, priority int) node {
	return node{val: val, priority: priority}
}

func NewPriorityQueue() priorityQueue {
	return priorityQueue{}
}

func (pq *priorityQueue) Enqueue(val string, priority int) []node {
	newNode := newNode(val, priority)
	pq.values = append(pq.values, newNode)
	return pq.bubbleUp()
}

func (pq *priorityQueue) bubbleUp() []node {
	if len(pq.values) <= 1 {
		return pq.values
	}
	currIdx := len(pq.values) - 1
	currElem := pq.values[currIdx]

	for true {
		parentIdx := int(math.Floor(float64((currIdx - 1) / 2)))
		parentElem := pq.values[parentIdx]

		if currElem.priority >= parentElem.priority { //If parent has a higher priority(lower value) then we are at the right place and can break
			break
		}
		pq.values[currIdx], pq.values[parentIdx] = pq.values[parentIdx], pq.values[currIdx]
		currIdx = parentIdx
	}
	return pq.values
}

func (pq *priorityQueue) Dequeue() (bool, node, []node) {
	if len(pq.values) < 1 {
		return false, node{}, []node{}
	}

	var maxPrtyItem = pq.values[0]
	pq.values[0], pq.values[len(pq.values)-1] = pq.values[len(pq.values)-1], pq.values[0]
	pq.values = pq.values[0 : len(pq.values)-1] //Remove the last element from the priority queue as we have already stored it in maxPrtyItem

	return true, maxPrtyItem, pq.bubbleDown()
}

func (pq *priorityQueue) bubbleDown() []node {
	if len(pq.values) <= 1 {
		return pq.values
	}
	currIdx := 0
	currElem := pq.values[0]
	for true {
		leftChildIdx := (2 * currIdx) + 1
		rightChildIdx := (2 * currIdx) + 2
		var leftChild, rightChild node
		swap := -1

		if leftChildIdx < len(pq.values) {
			leftChild = pq.values[leftChildIdx]
			if currElem.priority > leftChild.priority {
				swap = leftChildIdx
			}
		}

		if rightChildIdx < len(pq.values) {
			rightChild = pq.values[rightChildIdx]
			if swap == -1 && currElem.priority > rightChild.priority ||
				swap != -1 && rightChild.priority < leftChild.priority {
				swap = rightChildIdx
			}
		}

		if swap == -1 {
			break
		}
		pq.values[swap], pq.values[currIdx] = pq.values[currIdx], pq.values[swap]
		currIdx = swap

	}
	return pq.values
}
