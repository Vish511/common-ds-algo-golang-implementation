package patterns

//Kmuth-Morris-Pratt Algorithm

//Text: "onionionspl"   Pattern: "onions"
//LPS [0 0 0 1 2 0]
func KMPStringSearch(text, pattern string) bool {
	var lps = calculateLPS(pattern)
	var textRunes = []rune(text)
	var patternRunes = []rune(pattern)
	var i, j int

	for i < len(textRunes) {
		for j < len(patternRunes) {
			if textRunes[i] == patternRunes[j] {
				//Check if everything is a match at the end
				if j == len(patternRunes)-1 {
					return true
				}
				i++
				j++
			} else if textRunes[i] != patternRunes[j] {
				if j > 0 {
					j = lps[j-1]
				} else {
					i += 1
				}
			}
		}
	}
	return false
}

//Calculate longest prefix that is also a suffix
//e.g for the pattern - onions, the LPS would be [0 0 0 1 2 0]
func calculateLPS(pattern string) []int {
	if pattern == "" {
		return nil
	}
	var lps = make([]int, len(pattern))
	lps[0] = 0
	var left int
	var right int = 1
	runes := []rune(pattern)
	for right < len(runes) {
		if runes[right] == runes[left] {
			lps[right] = left + 1
			left++
		} else {
			lps[right] = 0
		}
		right++
	}
	return lps
}
