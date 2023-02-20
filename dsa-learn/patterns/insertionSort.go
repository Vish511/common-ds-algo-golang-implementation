package patterns

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currentElem := arr[i]
		for j := i - 1; j >= 0 && currentElem < arr[j]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return arr
}
