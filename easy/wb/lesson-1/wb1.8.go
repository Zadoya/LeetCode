package main

import "fmt"

func main() {
	array := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	for i := range array[:12] {
		fmt.Println(i)
		array[len(array) - 1], array[i] = array[i], array[len(array) - 1]
		array = array[:len(array) - 1]
	}

	fmt.Println(array)
}