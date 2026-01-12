// условие задачи - https://leetcode.com/problems/relative-sort-array/

package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
    countingMap := make(map[int]int, len(arr2))
    for i := 0; i < len(arr1); i++ {
        countingMap[arr1[i]]++
    }
    j := 0
	i := 0
    for i < len(arr1) {
		if j == len(arr2) {
			break
		}
        friq := countingMap[arr2[j]]
		for ; friq > 0; friq-- {
            arr1[i] = arr2[j]
			i++
        }
        delete(countingMap, arr2[j])
    	j++
    }

    lastPart := make([]int, 0, len(arr1) - len(arr2))
    for key, value := range countingMap {
		for i := 0; i < value; i++ {
			lastPart = append(lastPart, key)
		}
    }

	sort.Ints(lastPart)

	arr1 = append(arr1[:i], lastPart[:len(lastPart)]...)
	return arr1
}
