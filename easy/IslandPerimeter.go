// Условие задачи - https://leetcode.com/problems/island-perimeter/description

// можно обойтись подсчетом только верхнего и левого соседа(нижнего и правого), 
// умножив результат на 2

package main

func checkNeighbour(grid *[][]int, i, j int) int {
	counter := 0
	
	if j + 1 < len((*grid)[0]){
		counter += (*grid)[i][j + 1]
	}
	if j > 0 {
		counter += (*grid)[i][j - 1]
	}
	if i > 0 {
		counter += (*grid)[i - 1][j]
	}
	if i + 1 < len((*grid)){
		counter += (*grid)[i + 1][j]
	}
	return counter
}

func islandPerimeter(grid [][]int) int {
    perimeter := 0
    neighbours := 0
	for i, row := range grid {
		for j := range row {
			if row[j] == 1 {
				neighbours += checkNeighbour(&grid, i, j)
				perimeter++
			}
		}
	}
	return perimeter * 4 - neighbours
}