package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func scanLine(in *bufio.Reader, n int) ([]int, bool) {
	var (
		arr []int
	)

	line, err := in.ReadString('\n')
	if err != nil {
		return nil, false
	}

	numStr := strings.Split(line[:len(line) - 1], " ")
	if len(numStr) != n {
		return []int{}, false
	}

	arr = make([]int, 0, n)
	for i := range numStr {
		if strings.HasPrefix(numStr[i], "0") || strings.HasPrefix(numStr[i], "-0") {
			return []int{}, false
		}
		num, err := strconv.ParseInt(numStr[i], 10, 64)
		if err != nil {
			return []int{}, false
		}
		arr = append(arr, int(num))
	}

	return arr, true 
}

func main() {
	var (
		in         *bufio.Reader
		out        *bufio.Writer
		t, n       []int
		a, b       []int
		ok1, ok2   bool
	)

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t, _ = scanLine(in, 1)
	for i := 0; i < t[0]; i++ {
		n, _ = scanLine(in, 1)
		a, ok1 = scanLine(in, n[0])
		b, ok2 = scanLine(in, n[0])
		if !ok1 || !ok2 {
			fmt.Fprintln(out, "no")
			continue
		}
		slices.Sort(a)
		if slices.Equal(a, b) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
