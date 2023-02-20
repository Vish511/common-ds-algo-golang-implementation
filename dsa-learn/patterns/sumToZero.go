package patterns

//[-3 -2 -1 0 1 2 3]
func PairSumToZero(arr []int) []int {
	var pairSumZero = []int{}
	var left int
	var right int = len(arr) - 1
	var sum int
	for left < right {
		sum = arr[left] + arr[right]
		if sum > 0 {
			right--
		} else if sum < 0 {
			left++
		} else {
			pairSumZero = append(pairSumZero, arr[left], arr[right])
			left++
			right++
		}
	}
	return pairSumZero
}
