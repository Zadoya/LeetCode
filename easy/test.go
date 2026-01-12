package main

import (
	"fmt"
	"testing"
	"unsafe"
)
func KeySprintf(subjects, kinds []int) bool {
	orderIndex := make(map[string]int64)
	for i := range subjects {
		key := fmt.Sprintf("%d_%d", subjects[i], kinds[i])
		orderIndex[key]++
	}
	
	return true
}

func KeyStruct(subjects, kinds []int) bool {
	
	type key struct {
		subjectID int
		kindID    int
	}

	orderIndex := make(map[key]int64)
	for i := range subjects {
		k := key{subjects[i], kinds[i]}
		orderIndex[k]++
	}
	
	return true
}

func BenchmarkKeySprintf(b *testing.B) {
	subjects := []int{1,2,3,4,5,6,7,8,9}
	kinds := []int{1,2,3,4,5,6,7,8,9}
	b.ResetTimer()
	for range b.N {
		KeySprintf(subjects, kinds)
	}
}

func BenchmarkKeyStruct(b *testing.B) {
	subjects := []int{1,2,3,4,5,6,7,8,9}
	kinds := []int{1,2,3,4,5,6,7,8,9}
	b.ResetTimer()
	for range b.N {
		KeyStruct(subjects, kinds)
	}
}


func buzz(dp map[int64]struct{}) (map[int64]struct{}, error) {
	dp[1] = struct{}{}
	dp[2] = struct{}{}
	dp[3] = struct{}{}
	dp[4] = struct{}{}

	return dp, nil
}
/*
func main() {
	dp := make(map[int64]struct{})
	dp[10] = struct{}{}
	dp[11] = struct{}{}
	fmt.Println(dp)
	dp, err := buzz(dp)
	if err != nil {
		return
	}
	fmt.Println(dp)
}*/
