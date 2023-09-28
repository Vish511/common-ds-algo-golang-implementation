package main

import (
	"dsapatterns/priorityqueue"
	"fmt"
)

func main() {
	//Uncomment to run each algorithm
	// fmt.Println(patterns.IsAnagram("nisha", "shani"))
	// fmt.Println(patterns.PairSumToZero([]int{-4, -3, -4, -1, 0, 6, 7, 8}))
	// fmt.Println(patterns.CountUniqueValues([]int{1, 1, 1, 1, 4, 2, 2, 3, 3, 4, 4, 5, 6, 7, 7, 7, 8}))
	// fmt.Println(patterns.MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4))
	// fmt.Println(patterns.SumRange(3))
	// fmt.Println(patterns.FactorialIterative(3))
	// fmt.Println(patterns.FactorialRecursive(4))
	// fmt.Println(patterns.GetAllOddNumbersHelperRecursion([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	// fmt.Println(patterns.GetAllOddNumbersPureRecursion([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	// fmt.Println(patterns.LinearSearch([]int{1, 2, 6, 7, 8, 3, 4, 10, 5}, 100))
	// fmt.Println(patterns.BinarySearch([]int{1, 5, 10, 15, 30, 70, 90, 120}, 12200))
	// fmt.Println(patterns.CountOccurenceOfTextInString("ghmoortamootathondimoota", "moota"))
	// fmt.Println(patterns.KMPStringSearch("onionionspl", "onions"))
	// fmt.Println(patterns.BubbleSort([]int{37, 45, 29, 8, 12, 88, -3}))
	// fmt.Println(patterns.SelectionSort([]int{-1, 0, 0, -3, 5, 1, 2, 3, 4, -2}))
	// fmt.Println(patterns.InsertionSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	// fmt.Println(patterns.MergeSort([]int{5, 2, 3, 1, 4}))
	// fmt.Println(patterns.QuickSort([]int{3, 2, 1}))
	// fmt.Println(patterns.GetMaxDigit([]int{0, 000, 0000}))
	// arr, err := patterns.RadixSort([]int{10, 6, 20036, 355, 5340, 23, 4560, 5000})
	// fmt.Println(arr, err)
	// fmt.Println(int(math.Log10(423)))
	// sll := patterns.NewSinglyLinkedList()
	// sll.Push(1)
	// sll.Push(2)
	// sll.Push(4)
	// sll.Insert(2, 3)
	// // sll.Push(4)
	// // sll.Push(5)
	// // sll.Reverse()
	// // // sll.Set(2, 2)

	// // // sll.Insert(2, 15)
	// // // sll.Push(5)
	// // // fmt.Println(sll.Shift())
	// // // sll.Unshift(10)
	// // // sll.Unshift(5)
	// sll.Traverse()

	// dll := doublylinkedlist.NewDoublyLinkedList()
	// dll.Push(1)
	// dll.Push(2)

	// dll.Push(3)
	// dll.Push(4)
	// dll.Push(5)
	// dll.Reverse()
	// // dll.Push(10)
	// // dll.Push(15)
	// dll.Traverse()

	// stack := stack.NewStack()
	// stack.Push(1)
	// // stack.Push(2)
	// // stack.Push(3)
	// fmt.Printf("%+v", stack)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Pop())
	// q := queue.NewQueue()
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)
	// q.Enqueue(4)
	// q.Enqueue(5)
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())

	// fmt.Println(q)
	// bst := binarysearchtree.NewBinarySearchTree()
	// bst.Insert(10)
	// bst.Insert(6)
	// bst.Insert(15)
	// bst.Insert(13)
	// bst.Insert(3)
	// bst.Insert(1)
	// bst.Insert(8)
	// bst.Insert(20)
	// fmt.Println(bst.BFS())
	// fmt.Println(bst.DFSPreOrder())
	// fmt.Println(bst.DFSPostOrder())
	// fmt.Println(bst.DFSInOrder())

	// binHeap := binaryheap.NewMaxBinaryHeap()
	// // binHeap.Insert(55)
	// binHeap.Insert(41)
	// binHeap.Insert(39)
	// binHeap.Insert(33)
	// binHeap.Insert(18)
	// binHeap.Insert(27)
	// binHeap.Insert(12)
	// // binHeap.Insert(1)

	// fmt.Println(binHeap.Insert(55))
	// fmt.Println(binHeap.ExtractMax())

	pq := priorityqueue.NewPriorityQueue()

	pq.Enqueue("Headache", 3)
	pq.Enqueue("Fever", 2)
	pq.Enqueue("Fracture", 1)
	pq.Enqueue("Head Injury", 0)

	pq.Enqueue("Heart Attack", 0)
	fmt.Println(pq.Enqueue("General checkup", 4))

	for i := 0; i < 5; i++ {
		_, maxPrty, allItems := pq.Dequeue()
		fmt.Println("Patient ", i+1, "- ", maxPrty)
		fmt.Println("remaining patients ", allItems)
		pq.Enqueue("Heart Attack", 0)

	}

	// fmt.Println(binHeap.Insert(199))

	//    20
	//15
	//
	//
}

// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var count = 0
// var lk sync.Mutex

// func main() {
// 	var wg sync.WaitGroup
// 	fmt.Println("Hello, 世界")
// 	wg.Add(1)
// 	go func() {
// 		Inc()
// 		wg.Done()
// 	}()
// 	Inc()

// 	wg.Wait()

// 	fmt.Println("count is", count)

// }
// func Inc() {
// 	lk.Lock()
// 	count++
// 	lk.Unlock()
// }
