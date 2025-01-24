package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadFileData(filename string) (string, int64) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	n := (int64)(0)
	c := (int64)(0)
	s := ""
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if scanner.Scan() {
		c, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if scanner.Scan() {
		s = scanner.Text()
	}

	return s[:n], c
}

func cenzor(s string, c int) string {
	countA := 0
	countB := 0

	left := 0
	prev := 0
	for right := 0; right < len(s); right++ {
		switch s[right] {
		case 'A':
			if countB > 0 {
				countA, countB = 0, 0
			}
			countA++
		case 'B':
			if countA > 0 {
				countB++
			}
		}
		if prev < (countA * countB) {
			prev = countA * countB
		}
		if countA * countB > c {
			switch s[left] {
			case 'A':
				countA--
			case 'B':
				countB--
			}
			left++
		}
	}
	
}

func main() {
	s, c := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, cenzor(s, int(c)))
}