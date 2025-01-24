// условие задачи - https://leetcode.com/problems/valid-parentheses/description/

package main

func handler(b byte, stack *[]byte) bool {
	if len(*stack) != 0 && b == (*stack)[len(*stack)-1] {
		*stack = (*stack)[:len(*stack)-1]
		return true
	}
	return false
}

func isValid(s string) bool {
    openedBracket := make([]byte, 0, len(s))                 

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
            openedBracket = append(openedBracket, ')')
		case '{':
            openedBracket = append(openedBracket, '}')
		case '[':
            openedBracket = append(openedBracket, ']')
		case ']':
			if !handler(']', &openedBracket) {
				return false
			}
		case ')':
			if !handler(')', &openedBracket) {
				return false
			}
		case '}':
			if !handler('}', &openedBracket) {
				return false
			}
		}
	}
	return len(openedBracket) == 0
}

// with map 
/*
func isValid(s string) bool {
	openedBracket := make([]byte, 0, len(s)) 
	bracketsMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if bracket, ok := bracketsMap[s[i]]; ok {
			if len(openedBracket) != 0 && openedBracket[len(openedBracket) - 1] == bracket {
				openedBracket = openedBracket[:len(openedBracket) - 1]
			} else {
				return false
			}
		} else {
			openedBracket = append(openedBracket, s[i])
		}
	}
	return len(openedBracket) == 0
}
*/