package patterns

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	groupByStr1 := make(map[string]int)
	for _, v := range str1 {
		char := string(v)
		groupByStr1[char] += 1
	}

	for _, v := range str2 {
		char := string(v)
		if _, ok := groupByStr1[char]; !ok {
			return false
		}
	}
	return true
}
