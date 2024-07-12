// условие задачи - https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/description

package main

func reverseParentheses(s string) string {
    startBrackets := 0
	endBrackets := 0
	start := 0
	end := len(s) - 1
	startPart := make([]byte, len(s) / 2)
	endPart := make([]byte, len(s) / 2)
	
	for start < end {
		if s[start] == '(' {
			startBrackets++
		}
		if s[end] == ')' {
			endBrackets++
		}
		if startBrackets == endBrackets {
			start++
			end++
		} else if startBrackets < endBrackets {
			start++
		} else {
			end++
		}
	}

	return result
}