// условие задачи - https://leetcode.com/problems/pascals-triangle-ii/description/

package main

// time complexity - O(n), space complexity - O(n)
// через формулу получилось проще
func getRow2(rowIndex int) []int {
	row := make([]int, rowIndex + 1)

    row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i - 1] * (rowIndex - i + 1) / i
	}
	return row
}

// time complexity - O(n^2), space complexity - O(n)
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex + 1)

    row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		tmp := 1
        for j := 1; j <= rowIndex; j++ {
            row[j], tmp = row[j] + tmp, row[j]
        }
	}
	return row
}
