package patterns

func CountOccurenceOfTextInString(str, searchText string) int {
	var count int
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == rune(searchText[0]) {
			for j := 1; j < len(searchText); j++ {
				if runes[i+j] != rune(searchText[j]) {
					break
				}
				if j == len(searchText)-1 {
					count++
				}
			}
		}
	}
	return count
}
