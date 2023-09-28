package binaryheap

import (
	"fmt"
	"math"
)

//Binary heap is a tree in which a node can have atmost 2 children.
//It is compact meaning it is completely filled before moving to next level
//Max Binary Heap - In which the parent node is greater than the children
//Min Binary Heap - In which the children are greater than the parent
//Note: No ordering between sibling nodes

//Formula to be known
//Given a parent, to find the children, the formula is:
//Left child : 2n+1
//Right child: 2n+2

//Given a child, to find the parent:
//(n-1)/2

type binaryHeap struct {
	values []int
}

func NewMaxBinaryHeap() binaryHeap {
	return binaryHeap{}
}

func (bh *binaryHeap) Insert(val int) []int {
	bh.values = append(bh.values, val)

	return bh.BubbleUp()
}

func (bh *binaryHeap) BubbleUp() []int {
	if len(bh.values) <= 1 {
		return bh.values
	}
	currIdx := len(bh.values) - 1
	currElem := bh.values[currIdx]

	for currIdx > 0 {
		parentIdx := int(math.Floor(float64((currIdx - 1) / 2)))
		parentElem := bh.values[parentIdx]
		if currElem <= parentElem {
			break
		}
		//Else we swap
		bh.values[currIdx], bh.values[parentIdx] = bh.values[parentIdx], bh.values[currIdx]
		currIdx = parentIdx
	}
	return bh.values
}

func (bh *binaryHeap) ExtractMax() (bool, int, []int) {
	if len(bh.values) == 0 {
		return false, math.MinInt, []int{}
	}
	max := bh.values[0]
	bh.values[0], bh.values[len(bh.values)-1] = bh.values[len(bh.values)-1], bh.values[0]
	bh.values = bh.values[0 : len(bh.values)-1]
	bh.BubbleDown()
	return true, max, bh.values
}

func (bh *binaryHeap) BubbleDown() []int {
	fmt.Println(bh.values)
	if len(bh.values) <= 1 {
		return bh.values
	}
	currIdx := 0
	currElem := bh.values[0]
	for true {
		leftChildIdx := int(math.Floor(float64((2*currIdx + 1))))
		rightChildIdx := int(math.Floor(float64((2*currIdx + 2))))
		var leftElem, rightElem int
		swap := -1
		if leftChildIdx < len(bh.values) {
			leftElem = bh.values[leftChildIdx]
			if leftElem > currElem {
				swap = leftChildIdx
			}
		}

		if rightChildIdx < len(bh.values) {
			if (swap == -1 && rightElem > currElem) || (swap != -1 && rightElem > leftElem) {
				swap = rightChildIdx
			}
		}
		if swap == -1 {
			break
		}
		bh.values[currIdx], bh.values[swap] = bh.values[swap], bh.values[currIdx]
		currIdx = swap

	}
	return bh.values
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
