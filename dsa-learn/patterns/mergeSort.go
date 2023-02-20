package patterns

import (
	"math"
)

func MergeHelper(arr1, arr2 []int) []int {
	var merged []int
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

// [5,2,3,1,4]
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var mid = int(math.Ceil((float64(len(arr)) / float64(2))))
	var left = MergeSort(arr[:mid])
	var right = MergeSort(arr[mid:])
	return MergeHelper(left, right)
}
