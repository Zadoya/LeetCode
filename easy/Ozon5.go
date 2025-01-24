package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type wirehouse [][]byte

func scanMap(in *bufio.Reader, n int64, m int64) (wirehouse, int, int, error) {
	var (
		a, b int
		line string
	)
	result := make([][]byte, 0, n)

	for i := 0; i < int(n); i++ {
		if _, err := fmt.Fscan(in, &line); err != nil {
			return nil, a, b, err
		}
		if len(line) != int(m) {
			return nil, a, b, fmt.Errorf("incorrect number of arguments in string")
		}
		result = append(result, []byte(line))
		if j := strings.Index(line, "A"); j != -1 {
			a = i * int(m) + j
		} 
		if j := strings.Index(line, "B"); j != -1 {
			b = i * int(m) + j
		} 
	}
	return result, a, b, nil
}


func pathToTheLeft(wirehouse *wirehouse, place int, width int, path byte) {
	if place == 0 {
		return
	}
	rightMax := place % width
	if rightMax % 2 == 1 {
		rightMax--
		for i := 0; i <= rightMax; i++ {
			(*wirehouse)[0][i] = path
		}
		for i := 0; i <= place / width; i++ {
			(*wirehouse)[i][rightMax] = path
		}
	} else {
		for i := 0; i < rightMax; i++ {
			(*wirehouse)[0][i] = path
		}
		for i := 0; i < place / width; i++ {
			(*wirehouse)[i][rightMax] = path
		}
	}
}

func pathToTheRight(wirehouse *wirehouse, place int, height int, width int, path byte) {
	if place == height * width - 1 {
		return
	}
	leftMax := place % width
	if leftMax % 2 == 1 {
		leftMax++
		for i := width - 1; i >= leftMax; i-- {
			(*wirehouse)[height - 1][i] = path
		}
		for i := height - 1; i >= place / width; i-- {
			(*wirehouse)[i][leftMax] = path
		}
	} else {
		for i := width - 1; i > leftMax; i-- {
			(*wirehouse)[height - 1][i] = path
		}
		for i := height - 1; i > place / width; i-- {
			(*wirehouse)[i][leftMax] = path
		}
	}
}

func scanningDataset() ([]wirehouse, error) {
	var (
		in         *bufio.Reader
		t, n, m    int64
		err        error
		dataset    []wirehouse
	)

	in = bufio.NewReader(os.Stdin)

	if _, err = fmt.Fscan(in, &t); err != nil {
		return nil, err
	}

	dataset = make([]wirehouse, 0, t)
	for i := 0; i < int(t); i++ {
		var data wirehouse
		if _, err = fmt.Fscan(in, &n, &m); err != nil {
			return nil, err
		}

		data, a, b, err := scanMap(in, n, m)
		if err != nil {
			return nil, err
		}
		if a < b {
			pathToTheLeft(&data, a, int(m), 'a')
			pathToTheRight(&data, b, int(n), int(m), 'b')
		} else {
			pathToTheLeft(&data, b, int(m), 'b')
			pathToTheRight(&data, a, int(n), int(m), 'a')
		}

		dataset = append(dataset, data)
	}

	return dataset, nil
}

func main() {
	dataset, err := scanningDataset()
	if err != nil {
		fmt.Println("Scanning error:", err)
	}

	var out  *bufio.Writer

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, data := range dataset {
		for _, line := range data {
			fmt.Fprintln(out, string(line))
		}
	}
}