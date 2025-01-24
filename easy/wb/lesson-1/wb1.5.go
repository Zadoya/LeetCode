package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readData() (map[string]int, int) {
	n := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	languages := make(map[string]int)
	for i := 0; i < int(n); i++ {
		m := int64(0)
		if scanner.Scan() {
			m, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
		for j := 0; j < int(m); j++ {
			if scanner.Scan() {
				languages[scanner.Text()]++
			}
		}
	}

	return languages, int(n)
}

func main() {
	languages, n := readData()
	countPopular := 0
	popularLanguages := make([]string, 0, len(languages))
	for key, value := range languages {
		if value == n {
			countPopular++
			popularLanguages = append(popularLanguages, key)
		}
	}
	fmt.Fprintln(os.Stdout, countPopular)
	for _, language := range popularLanguages {
		fmt.Fprintln(os.Stdout, language)
	}
	fmt.Fprintln(os.Stdout, len(languages))
	for key, _ := range languages {
		fmt.Fprintln(os.Stdout, key)
	}
}