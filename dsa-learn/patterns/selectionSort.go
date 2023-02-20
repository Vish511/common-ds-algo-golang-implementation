package patterns

func SelectionSort(arr []int) []int {
	var smallest int
	for i := 0; i < len(arr); i++ {
		smallest = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		if i != smallest {
			arr[i], arr[smallest] = arr[smallest], arr[i]
		}
	}
	return arr
}
