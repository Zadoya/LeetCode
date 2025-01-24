package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func loadFileData(filename string) []int64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	n := (int64)(0)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	a := make([]int64, n)
	for i := 0; i < int(n) && scanner.Scan(); i++ {
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return a
}

func median(a []int64) []int64 {
	var left, right int

	slices.Sort(a)
	result := make([]int64, len(a))

	if len(a) % 2 == 0 {
		left = len(a) / 2 - 1
		right = len(a) / 2
	} else {
		left = len(a) / 2
		right = len(a) / 2
	}
	for i := 0; left >= 0 && right < len(a); i++ {
		switch {
		case left == right:
			result[i] = a[left]
			left--
			right++
		case a[left] <= a[right]:
			result[i] = a[left]
			left--
			i++
			result[i] = a[right]
			right++
		}
	}
	return result
}

func main() {
	a := loadFileData("input.txt")
	a = median(a)
	for i := 0; i < len(a); i++ {
		fmt.Fprintf(os.Stdout,"%v ", a[i])
	}
	fmt.Fprintln(os.Stdout)
}