package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var (
		in   *bufio.Reader
		out  *bufio.Writer
		t    int
	)
	
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscanln(in, &t)
	
	a := make([]string, t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(in, &a[i])
	}

	for i := 0; i < t; i++ {
		if len(a[i]) == 1 {
			fmt.Fprintln(out, 0)
			continue
		}
		minNum := byte(58)
		place  := -1
		for j := 0; j < len(a[i]); j++ {
			if j > 0 && a[i][j] > a[i][j-1] {
				place = j - 1
				break
			}
			if a[i][j] <= minNum {
				minNum = a[i][j]
				place = j
			}
		}
		if place >= 0 {
			a[i] = a[i][:place] + a[i][place+1:]
		}
		fmt.Fprintln(out, a[i])
	}
}
