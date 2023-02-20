package patterns

import "math"

func BinarySearch(arr []int, num int) int {
	var left int = 0
	var right int = len(arr) - 1

	for left < right {
		var mid int = int(math.Round(float64(left+right) / 2))
		if arr[mid] == num {
			return mid
		} else if num > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
