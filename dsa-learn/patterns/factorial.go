package patterns

//4 - 4*3*2*1
func FactorialIterative(num int) int {
	var fact int = 1
	for i := num; i > 1; i-- {
		fact *= i
	}
	return fact
}

func FactorialRecursive(num int) int {
	if num <= 1 {
		return 1
	}
	return num * FactorialIterative(num-1)
}
