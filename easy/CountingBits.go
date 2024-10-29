// условие задачи - https://leetcode.com/problems/counting-bits/description/

package main

func countBits(n int) []int {
    result := make([]int, 0, n)
	result = append(result, 0)
	for i := 1; i <= n; i++ {
		result = append(result, result[i&(i-1)] + 1)
	}
	return result
}
