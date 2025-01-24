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
	size := (int64)(0)
	luckyNum := (int64)(0)
	if scanner.Scan() {
		size, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if scanner.Scan() {
		luckyNum, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	nums := make([]int64, size)
	for i := 0; i < int(size) && scanner.Scan(); i++ {
		nums[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	return nums, luckyNum
}

func numsSum(nums []int64, luckyNum int64) int {
	counter := 0
	currSum := int64(0)
	right := 0
	for left := 0; left < len(nums); left++ { // может left???
		for currSum < luckyNum && right < len(nums) {
			currSum += nums[right]
			right++
		}
		if currSum == luckyNum {
			counter++
		}
		currSum -= nums[left]
	}
	return counter 
}

func main() {
	nums, luckyNum := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, numsSum(nums, luckyNum))
}