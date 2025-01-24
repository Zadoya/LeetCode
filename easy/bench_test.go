package main

import (
	"testing"
)

func addIfUnique(arr []int, value int) {
    for _, v := range arr {
        if v == value {
            return // Уже существует, не добавляем.
        }
    }
	arr = append(arr, value)
    return
}

func BenchmarkArray(b *testing.B) {
	subjectID := []int{1234, 345, 34, 456, 456, 24515, 124545, 45532}
	result    := make([]int, 0, 100) 
	b.ResetTimer()
	for range b.N {
		for _, val := range subjectID { 
			addIfUnique(result, val)
		}
	}
}

func BenchmarkMap(b *testing.B) {
	subjectID := []int{1234, 345, 34, 456, 456, 24515, 124545, 45532}
	result    := make(map[int]struct{}, 100) 
	b.ResetTimer()
	for range b.N {
		for _, val := range subjectID { 
			result[val] = struct{}{}
		}
	}
}