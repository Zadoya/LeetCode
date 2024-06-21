// условие задачи - https://leetcode.com/problems/search-a-2d-matrix/description/

func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    cols := len(matrix[0])
    
    left := 0
    right := rows * cols - 1
    for left <= right {
        middle := int((left + right) / 2)
        current := matrix[middle / len(matrix[0])][middle % len(matrix[0])]
        if target > current {
            left = middle + 1
        } else if target < current {
            right = middle - 1      
        } else {
            return true
        }
    }
    return false
}