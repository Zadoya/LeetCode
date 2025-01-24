package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readData() ([][2]int64, int64) {
	n := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	arr := make([][2]int64, n)
	for i := 0; i < int(n); i++ {
		if scanner.Scan() {
			arr[i][0], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
		if scanner.Scan() {
			arr[i][1], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}

	return arr, n
}

func countLiers(arr [][2]int64, n int64) int {
	set := make(map[int64]int64, len(arr))
	result := 0
	for i := 0; i < len(arr); i++ {
		if arr[i][0] + arr[i][1] == n - 1 {
			set[arr[i][0]] = arr[i][1]
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i][0] == arr[i][1] && len(arr) % 2 == 0 {
			continue
		}
		if first, ok := set[arr[i][1]]; ok && first == arr[i][0] {
			result++
		}
	}
	return result
}	

func main() {
	arr, n := readData()
	fmt.Fprintln(os.Stdout, countLiers(arr, n))
}