// условие задачи - https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])
	islands := 0

    var dfs func(row, col int) 
    dfs = func(row, col int) {
        if row < 0 || col < 0 || row >= rows || col >= cols || grid[row][col] == '0' {
            return
        }
        grid[row][col] = '0'
        dfs(row - 1, col)
        dfs(row + 1, col)
        dfs(row, col - 1)
        dfs(row, col + 1)
    }

    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == '1' {
                dfs(row, col)
                islands++
            }
        }
    }

	return islands
}