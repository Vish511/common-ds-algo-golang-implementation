package patterns

func SumRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + SumRange(num-1)
}
