package patterns

import (
	"math"
	"strconv"
)

func RadixSort(arr []int) ([]int, error) {
	maxDigits := GetMaxDigit(arr)
	var buckets = make([][]int, 10, 10)
	var err error

	for i := maxDigits - 1; i >= 0; i-- {
		for j := 0; j < len(arr); j++ {
			strNum := string(strconv.Itoa(arr[j]))
			position := i - (maxDigits - len(strNum))
			numAtPosition := 0
			if position >= 0 {
				numAtPosition, err = strconv.Atoi(string(strNum[position]))
				if err != nil {
					return []int{}, err
				}
			}
			buckets[numAtPosition] = append(buckets[numAtPosition], arr[j])
		}
		arr = []int{}
		for b := 0; b < len(buckets); b++ {
			arr = append(arr, buckets[b]...)
			buckets[b] = []int{}
		}
	}
	return arr, nil
}

func GetMaxDigit(arr []int) int {
	if len(arr) < 1 {
		return 0
	}
	var maxNum int
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}

	return int(math.Log10(float64(maxNum))) + 1
}
