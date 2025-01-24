package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readData() [][2]int64 {
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

	return arr
}

func countSumBoxes(arr [][2]int64) map[int64]int64 {
	set := make(map[int64]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		set[arr[i][0]] += arr[i][1]
	}
	return set
}

func main() {
	arr := readData()
	m := countSumBoxes(arr)
	
	i := 0
	for key, value := range m {
		arr[i][0] = key
		arr[i][1] = value
		i++
	}
	arr = arr[:i]
	
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(os.Stdout, arr[i][0], arr[i][1])
	}
}