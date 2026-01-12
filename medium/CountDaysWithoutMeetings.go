// условие задачи - https://leetcode.com/problems/count-days-without-meetings/description

package main

import (
//	"sort"
)

/*
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for i := 0; i < len(meetings) - 1; {
		next := i + 1
		if meetings[i][1] >= meetings[next][0] {
			meetings[i][1] = max(meetings[i][1], meetings[next][1])
			meetings = append(meetings[:next], meetings[next + 1:]...)
		} else {
			i++
		}
	}
	diff := 0
	for i := 0; i < len(meetings) - 1; i++ {
		diff += meetings[i + 1][0] - meetings[i][1] - 1
	}
	return meetings[0][0] + diff + daysmeetings[len(meetings) - 1][1] - 1
}
*/

func countDays(days int, meetings [][]int) int {
	diffsMap := make([]int, days + 1)

	for _, metting := range meetings {
		diffsMap[metting[0] - 1]++ 
		diffsMap[metting[1]]--
	}
	freedays := 0
	if diffsMap[0] == 0 {
		freedays++
	}

	for i := 1; i < len(diffsMap) - 1; i++ {
		diffsMap[i] += diffsMap[i - 1]
		if diffsMap[i] == 0 {
			freedays++
		}
	}
	return freedays
}

// нужно объединить все интервалы, дальше проверить, сколько всего интервалов и покрывают ли они начало и конец,
// ответ получится начало + конец + кол-во интервалов - 1
func main() {
	println(countDays(8, [][]int{{3,4},{4,8},{2,5},{3,8}}))
}