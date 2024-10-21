// условие задачи - https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

func letterCombinations(digits string) []string {
	buttonMap := map[string]string {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

    combinations := make([][]byte, 0, len(digits) * 4)
	for _, digit := range digits {
		
	}
}