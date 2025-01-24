package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
)

func loadFileData(filename string) ([]int64, int64) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	n := (int64)(0)
	k := (int64)(0)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if scanner.Scan() {
		k, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	a := make([]int64, n)
	for i := 0; i < int(n) && scanner.Scan(); i++ {
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return a, k
}

func countDays(d []int64, r int64) int {
	ans := 0
	slices.Sort(d)

	left, right := 0, 0
	for left < len(d) && right < len(d) {
		if d[right] - d[left] <= r {
			ans = max(ans, right - left + 1)
			right++
		} else {
			left++
		}
	}
	return ans
}

func main() {
	a, k := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, countDays(a, k))
}