// условие задачи - 

package main

func isSubsequence(s string, t string) bool {
    ptr := 0
    for i := range t {
        if ptr < len(s) && t[i] == s[ptr] {
                ptr++
        }
    }
    if ptr == len(s) {
        return true
    }
    return false
}