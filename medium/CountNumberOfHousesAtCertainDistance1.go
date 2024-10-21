//условие задачи - https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-i/description/

package main

// faster and less space
func countOfPairs(n int, x int, y int) []int {
	distances := make([]int, n)

	if x > y {
        x, y = y, x
    }
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			val := min(j - i, 1 + abs(i - x) + abs(j - y))
			distances[val - 1] += 2
		}
	}
	return distances
}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

/* мое решение

type Node struct {
	Value 	 int
	Distance int 
} 

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
	queue := make([]Node, 0, len(*distances))
	queue  = append(queue, Node{i, 0})
	visited := make([]bool, len(*distances))
	visited[i] = true
	
	for len(queue) != 0 {
		nodes := nodes(queue[0].Value, n, x, y)
		distance := queue[0].Distance
		queue = queue[1:]
		for _, node := range nodes {
			if !visited[node] {
				queue = append(queue, Node{Value: node, Distance: distance + 1})
				(*distances)[distance]++
				visited[node] = true
			}
		}
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
*/
