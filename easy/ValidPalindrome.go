// условие задачи - https://leetcode.com/problems/valid-palindrome

package main

func isPalindrome(s string) bool {
    start := 0
    end := len(s) - 1

    for start < end {
        if !isAlfabet(s[start]) {
            start++
        }
        if !isAlfabet(s[end]) {
            end--
        }
        for start < end && isAlfabet(s[start]) && isAlfabet(s[end]) {
            if toLower(s[start]) == toLower(s[end]) {
                start++
                end--
            } else {
                return false
            }
        }
    }

    return true
}

func isAlfabet(letter byte) bool {
    return letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' || letter >= '0' && letter <= '9'
}

func toLower(letter byte) byte {
    if letter >= 'A' && letter <= 'Z' {
        return letter + 32
    }
    return letter
}