package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData() []int64 {
	var a []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	if scanner.Scan() {
		a = strings.Split(scanner.Text(), " ")
	}
	arr := make([]int64, len(a))

	for i := 0; i < len(a); i++ {
		arr[i], _ = strconv.ParseInt(a[i], 10, 64)
	}
	return arr
}

func getUniq(arr []int64) []int64 {
	set := make(map[int64]int, len(arr))
	for i := 0; i < len(arr); i++ {
		set[arr[i]]++
	}
	result := make([]int64, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if amount := set[arr[i]]; amount == 1 {
			result = append(result, arr[i])
		}
	}
	return result
}	

func main() {
	arr := readData()
	result := getUniq(arr)
	for i := 0; i < len(result); i++ {
		fmt.Fprint(os.Stdout, result[i])
		if i != len(result) - 1 {
			fmt.Fprint(os.Stdout, " ")
		}
	}
	fmt.Fprintln(os.Stdout)
}