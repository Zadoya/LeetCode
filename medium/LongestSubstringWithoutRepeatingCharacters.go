// условие задачи - https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

// использую скользящее окно

func isUnique(s string) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
    
	return true
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	subString := ""

	for	_, val := range s {
		subString += string(val)
		if isUnique(subString) {
			maxLen = max(maxLen, len(subString))
		} else {
			subString = subString[1:]
		}
	}

	return maxLen
}