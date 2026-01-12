// условие задачи - https://leetcode.com/problems/length-of-last-word/

package main

func lengthOfLastWord(s string) int {
    var slen int

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            slen++
        } else if slen > 0 {
            return slen
        }
    }

    return slen
}

// using "strings" package
/*
func lengthOfLastWordSrtings(s string) int {
	slice := strings.Fields(s)
	return len(slice[len(slice) - 1])
}
*/