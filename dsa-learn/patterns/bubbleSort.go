package patterns

import "fmt"

func BubbleSort(arr []int) []int {
	var noSwaps bool //optimization for nearly sorted arrays to prevent checking multiple times
	for i := len(arr) - 1; i >= 0; i-- {
		noSwaps = true
		for j := 0; j < i; j++ {
			fmt.Println(arr[j], arr[j+1])
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}
	return arr
}
