// условие задачи - https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/

package main

func checkPowersOfThree(n int) bool {
    for n > 0 {
        if n % 3 == 2 {
            return false
        }
        n = n / 3
    }

    return true
}
