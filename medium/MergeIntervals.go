// условие задачи - https://leetcode.com/problems/merge-intervals/description

package main

import (
	"sort"
)

// O(n) - память
/*
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newIntervals := make([][]int, 0)
	for i := 0; i < len(intervals); {
		start := intervals[i][0]
		end := intervals[i][1]
		for ; i < len(intervals) && end >= intervals[i][0]; i++ {
			end = max(end, intervals[i][1])
		}
		newIntervals = append(newIntervals, []int{start, end})
	}

	return newIntervals
}
*/

// O(1) - память
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); {
		if intervals[i-1][1] >= intervals[i][0] {
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}

	return intervals
}