// условие задачи - https://leetcode.com/problems/pascals-triangle/description/

package main

func generate(numRows int) [][]int {
	traiangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		traiangle[i] = make([]int, i+1)
		traiangle[i][0], traiangle[i][i] = 1, 1
		for j:= 1; j < i; j++ {
			traiangle[i][j] = traiangle[i-1][j-1] + traiangle[i - 1][j]
		}
	}
	return traiangle
}
