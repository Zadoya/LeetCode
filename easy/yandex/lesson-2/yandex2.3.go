package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	r := (int64)(0)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if scanner.Scan() {
		r, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	d := make([]int64, n)
	for i := 0; i < int(n) && scanner.Scan(); i++ {
		d[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return d, r
}

func findComb(d []int64, r int64) int {
	counter := 0
	right := 0
	for left := 0; left < len(d); left++ {
		for right < len(d) && d[right] - d[left] <= r {
			right++
		}
		counter += len(d) - right
	}
	return counter 
}

func main() {
	d, r := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, findComb(d, r))
}