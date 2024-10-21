package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var (
		sum, n int64
		err error
	)
	if n, err = readNums(scanner, [2]int64{1, 1000}, 1); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if sum, err = readNums(scanner, [2]int64{0, 100_000_000}, n); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if sum % 2 == 0 {
		fmt.Fprintln(os.Stdout, "1")
	} else {
		fmt.Fprintln(os.Stdout, "0")
	}
}

func readNums(scanner *bufio.Scanner, numsRange[2]int64, n int64) (int64, error) {
	var result int64

	for i := 0; i < int(n); i++ {
		if scanner.Scan() {
			ai, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				return 0, fmt.Errorf("parsing input:", err)
			}
			if ai < numsRange[0] || ai > numsRange[1] {
				return 0, fmt.Errorf("uncorrect input data:", err)
			}
			result += ai	
		} else {
			return 0, fmt.Errorf("reading standard input:", scanner.Err().Error())
		}
	}

	return result, nil
}