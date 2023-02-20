package patterns

//Write a program to calculate maximum sum of n consecutive elements of an array
//Eg - Input Array:[1,2,5,2,8,1,5], n:2 --> Output Max Sum:10
func MaxSubArraySum(arr []int, n int) int {
	var maxSum int
	var tempSum int

	for i := 0; i < n; i++ {
		tempSum += arr[i]
	}
	for j := n; j < len(arr); j++ {
		tempSum = tempSum - arr[j-n] + arr[j]
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}

	return maxSum
}
