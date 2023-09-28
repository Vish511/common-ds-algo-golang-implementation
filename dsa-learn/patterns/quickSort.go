package patterns

// Quick sort using recursion
func QuickSort(arr []int) []int {
	return quickSortHelper(arr, 0, len(arr))
}

func quickSortHelper(arr []int, left, right int) []int {
	if left < right {
		pivotIdx := pivotHelper(arr, left, right)
		quickSortHelper(arr, left, pivotIdx)
		quickSortHelper(arr, pivotIdx+1, right-1)
	}
	return arr
}

func pivotHelper(arr []int, left, right int) int {
	pivot := arr[left]
	var pivotIndex int
	for i := left + 1; i < right; i++ {
		if arr[i] < pivot {
			pivotIndex++
			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
		}
	}
	arr[pivotIndex], arr[left] = arr[left], arr[pivotIndex]
	return pivotIndex
}
