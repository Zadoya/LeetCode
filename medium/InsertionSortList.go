package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func generate(m, n int) [][]int {
	matrix := make([][]int, m)
	unique := make(map[int]struct{}, m * n)

	for i := range matrix {
		row := make([]int, n)
		for j := range row {
			for {
				newNum := rand.Intn(10)
				if _, ok := unique[newNum]; ok == false {
					row[j] = newNum
					unique[newNum] = struct{}{}
					break
				}	
			}
			
		}
		matrix[i] = row
	}
	return matrix
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]

	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot - 1)
		quickSort(arr, pivot + 1, right)
	}
}

func intsToStrRange(arr []int) string {
	var str string

	quickSort(arr, 0, len(arr) - 1)
	str += strconv.Itoa(arr[0])
	for i := 1; i < len(arr); i++ {
		r := 0
		for i < len(arr) && arr[i] - arr[i - 1] == 1 {
			r++
			i++
		}
		if r > 0 {
			str += "-"
			str += strconv.Itoa(arr[i - 1])
			if i < len(arr) {
				str += ","
				str += strconv.Itoa(arr[i])
			}
		} else {
			str += ","
			str += strconv.Itoa(arr[i])
		}	
	}
	return str
}

func main() {
	fmt.Println(intsToStrRange([]int{1,3,4,2,9,8,16,11,0}))
}