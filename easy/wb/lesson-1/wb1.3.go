package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readData() ([]int64, int64) {
	num := int64(0)
	arr := make([]int64, 3)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < 3; i++ {
		if scanner.Scan() {
			arr[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	} 
	if scanner.Scan() {
		num, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return arr, num
}

func openCalculator(arr []int64, num int64) int {
	set := make(map[int64]struct{}, len(arr))
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = struct{}{}
	}
	add := make(map[int64]struct{})
	for num > 0 {
		if _, ok := set[num % 10]; !ok {
			add[num % 10] = struct{}{}
		}
		num /= 10
	}
	return len(add)
}	

func main() {
	arr, num := readData()
	fmt.Fprintln(os.Stdout, openCalculator(arr, num))
}