// условие задачи - https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/

package main

func hardestWorker(n int, logs [][]int) int {
	maxTime := 0
	employeer := -1
	prevTime := 0
	for _, task := range logs {
		workTime := task[1] - prevTime
		prevTime = task[1]
		if workTime > maxTime && task[0] <= n - 1 &&  task[0] >= 0 {
			maxTime = workTime
			employeer = task[0]
		} else if workTime == maxTime && employeer > task[0] && task[0] <= n - 1 &&  task[0] >= 0 {
			employeer = task[0]
		}
	}

	return employeer
}