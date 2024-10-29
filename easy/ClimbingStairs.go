// условие задачи - https://leetcode.com/problems/climbing-stairs/description/

package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	curr := 2 
	prev := 1	
	for i := 3; i <= n; i++ {
		curr, prev = curr + prev, curr
	}
    return curr
}
