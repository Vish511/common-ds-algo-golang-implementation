package patterns

//Implement a function called countUniqueValues
//which accepts a sorted array and counts the unique values in the array
//There can be negative values in the array

//Eg [1,1,1,1,2] - 2
func CountUniqueValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var left int
	var right int = left + 1
	for right < len(arr) {
		if arr[left] != arr[right] {
			arr[left+1] = arr[right]
			left++
		}
		right++
	}
	return left + 1
}
