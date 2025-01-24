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

func multiSum(a []int64) int64 {

	for i := 0; i < len(a); {
		if a[i] == 0 {
			a = append(a[:i], a[i + 1:]...)
		} else {
			i++
		}
	}
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			for k := j + 1; k < len(a); k++ {
				sum = (sum + (a[i] * a[j] * a[k])) % 1_000_000_007
			}
		}
	}
	return sum
}

func multiSum3(a []int64) int64 {

	for i := 0; i < len(a); {
		if a[i] == 0 {
			a = append(a[:i], a[i + 1:]...)
		} else {
			i++
		}
	}

	first := 0
	second := 1
	sum := int64(0)
	for second < len(a) {
		multiPair := a[first] * a[second]
		for third := 0; third < len(a); third++ {
			if third < first - 1 || third > second {
				sum = (sum + (multiPair * a[third])) % 1_000_000_007
			}
		}
		first++
		second++
	}
	return sum
}

func main() {
	a := loadFileData("input.txt")
	fmt.Fprintln(os.Stdout, multiSum(a))
	fmt.Fprintln(os.Stdout, multiSum3(a))
}