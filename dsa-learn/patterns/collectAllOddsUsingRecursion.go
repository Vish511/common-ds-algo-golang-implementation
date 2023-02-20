package patterns

func GetAllOddNumbersHelperRecursion(arr []int) []int {
	var odds []int
	var getOddsHelper func(nums []int)
	getOddsHelper = func(nums []int) {
		if len(nums) == 0 {
			return
		}
		num := nums[0]
		if num%2 != 0 {
			odds = append(odds, num)
		}
		getOddsHelper(nums[1:len(nums)])
	}
	getOddsHelper(arr)

	return odds
}

func GetAllOddNumbersPureRecursion(arr []int) []int {
	var odds []int
	if len(arr) == 0 {
		return odds
	}

	if arr[0]%2 != 0 {
		odds = append(odds, arr[0])
	}

	odds = append(odds, GetAllOddNumbersPureRecursion(arr[1:len(arr)])...)
	return odds
}
