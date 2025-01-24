// условие задачи - https://leetcode.com/problems/roman-to-integer/

package main

func romanToInt(s string) int {
	roman := map[byte]int{
		'I':1,
		'V':5,
		'X':10, 
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	result := 0
	
	for i := len(s) - 1; i >= 0; i-- {
		number := roman[s[i]]
		if i != len(s) - 1 && number < roman[s[i + 1]] {
			result -= number
		} else {
			result += number	
		}
	}

	return result
}