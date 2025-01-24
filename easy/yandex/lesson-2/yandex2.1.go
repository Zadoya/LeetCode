package main

import (
	"bufio"
	"fmt"
	"os"
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
	size := (int64)(0)
	if scanner.Scan() {
		size, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	nums := make([]int64, size)
	for i := 0; i < int(size) && scanner.Scan(); i++ {
		nums[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return nums
}

func prefixSum(nums []int64) []int64 {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i - 1]
	}
	return nums
}

func main() {
	nums := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, prefixSum(nums))
}