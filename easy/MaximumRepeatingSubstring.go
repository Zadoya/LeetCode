// условие задачи - https://leetcode.com/problems/maximum-repeating-substring/description/

func maxRepeating(sequence string, word string) int {
	if len(sequence) == 0 || len(word) == 0 {
		return 0
	}

	maxRep := 0
 i := 0
 j := 0
 counter := 0
	for i <= len(sequence) - len(word) {
	 if sequence[i:i + len(word)] == word {
		 counter++
		 i += len(word)
	 } else {
		 j++
		 i = j
		 maxRep = max(maxRep, counter)
		 counter = 0
	 }	
	}
 maxRep = max(maxRep, counter)
	return maxRep
}