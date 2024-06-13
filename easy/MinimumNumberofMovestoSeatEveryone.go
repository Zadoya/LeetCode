// условие задачи - https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description

import (
	"math"
	"sync"
	"sort"
)
/*
func partition(arr *[]int, left, right int) int {
	pivot := (*arr)[right]

	i := left
	for j := left; j < right; j++ {
		if (*arr)[j] < pivot {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}
	(*arr)[i], (*arr)[right] = (*arr)[right], (*arr)[i]
	return i
}

func quickSort(arr *[]int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot - 1)
		quickSort(arr, pivot + 1, right)
	}
}
*/

func minMovesToSeat(seats []int, students []int) int {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		//quickSort(&seats, 0, len(seats) - 1)
		sort.Ints(seats)
		wg.Done()
	}()
	go func() {
		//quickSort(&students, 0, len(students) - 1)
		sort.Ints(students)
		wg.Done()
	}()
	wg.Wait()
	
	result := 0.0
	for i := range seats {
		result += math.Abs(float64(seats[i] - students[i]))
	}
	return int(result)
}