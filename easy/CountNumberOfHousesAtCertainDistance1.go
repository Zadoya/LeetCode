//условие задачи - https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/description/

package main

import (
	"fmt"
)

func countOfPairs(n int, x int, y int) []int {
	distances := make([]int, n)

	if x > y {
        x, y = y, x
    }
	for i := 0; i < n; i++ {
		BFS(n, x, y, i, &distances) 
	}
	return distances
}

func BFS(n, x, y, i int, distances *[]int) {
	queue := make([]int, 0, len(*distances))
	queue  = append(queue, i)
	visited := make([]bool, len(*distances))
	visited[i] = true
	distance := 0
	
	for len(queue) != 0 {
		nodes := nodes(queue[0], n, x, y)
		queue = queue[1:]
		for _, node := range nodes {
			if !visited[node] {
				queue = append(queue, node)
				(*distances)[distance]++
				visited[node] = true
			}
		}
		distance++
	}
}

func nodes(i, n, x, y int) []int {
	nodes := make([]int, 0, 3)
	if i == x - 1 {
		nodes = append(nodes, y - 1)
	}
	if i == y - 1 {
		nodes = append(nodes, x - 1)
	}
	if i != n - 1 {
		nodes = append(nodes, i + 1)	
	}
	if i != 0 {
		nodes = append(nodes, i - 1)
	}
	return nodes
}


func main() {
	fmt.Println(countOfPairs(5, 2, 4))
}