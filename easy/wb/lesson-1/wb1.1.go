package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() ([]int64, []int64) {
	var a, b []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	if scanner.Scan() {
		a = strings.Split(scanner.Text(), " ")
	}
	if scanner.Scan() {
		b = strings.Split(scanner.Text(), " ")
	}
	arr1 := make([]int64, len(a))
	arr2 := make([]int64, len(b))

	for i := 0; i < len(a); i++ {
		arr1[i], _ = strconv.ParseInt(a[i], 10, 64)
	}
	for i := 0; i < len(b); i++ {
		arr2[i], _ = strconv.ParseInt(b[i], 10, 64)
	}

	return arr1, arr2
}

func countSameNums(arr1, arr2 []int64) int {
	counter := 0
	set := make(map[int64]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		set[arr1[i]]++
	}
	for i := 0; i < len(arr2); i++ {
		if _, ok := set[arr2[i]]; ok {
			counter++
			delete(set, arr2[i])
		}
	}
	return counter
}	

func main() {
	arr1, arr2 := readData()
	fmt.Fprintln(os.Stdout, countSameNums(arr1, arr2))
}