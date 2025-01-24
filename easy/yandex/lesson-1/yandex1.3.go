package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	i = [][]int64{{1}}
	c = [][]int64{{1,1},{1,0},{1,1}}
	o = [][]int64{{1,1,1},{1,0,1},{1,1,1}}
	l = [][]int64{{1,0},{1,1}}
	h = [][]int64{{1,0,1},{1,1,1},{1,0,1}}
	p = [][]int64{{1,1,1},{1,0,1},{1,1,1},{1,0,0}}
)

func loadFileData(filename string) [][]int64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	size := (int64)(0)
	if scanner.Scan() {
		size, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	nums := make([][]int64, size)
	for i := 0; i < int(size) && scanner.Scan(); i++ {
		str := scanner.Text()
		nums[i] = make([]int64, size)
		for j := 0; j < int(size); j++ {
			switch str[j] {
			case '.':
				nums[i][j] = 0
			case '#':
				nums[i][j] = 1
			}
		}
	}
	return nums
}

func compresseTable(table [][]int64) [][]int64 {
	prevStr := 0
	for i := 0; i < len(table); {
		currStr := 1
		emptyStr := true
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 1 {
				currStr++
				emptyStr = false
			}
			currStr = currStr << 1
		}
		if prevStr == currStr || emptyStr {
			table = append(table[:i], table[i + 1:]...)
		} else {
			i++
		}
		prevStr = currStr
	}
	return table
}

func invertTable(table [][]int64) [][]int64 {
	newTable := make([][]int64, len(table[0]))
	for i := 0; i < len(table); i++ {
		newRaw := make([]int64, len(table))
		for j := 0; j < len(table[i]); j++ {
			if i != j {
				newRaw[j] = table[j][i]
			}
		}
		newTable[i] = newRaw
	}
	return table
}

func defineSign(table [][]int64) {
	table = compresseTable(table)
	table = invertTable(table)
	//table = compresseTable(table)
	//table = invertTable(table)
	fmt.Println(table)
}

func main() {
	table := loadFileData("input.txt")
	defineSign(table)
}
