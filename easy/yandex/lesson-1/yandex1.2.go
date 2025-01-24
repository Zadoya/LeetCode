package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int64, 4)

	i := 0
	for i < 4 && scanner.Scan() {
		nums[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		i++
	}
	fmt.Fprintln(os.Stdout, min(nums[0], nums[1]) + 1, min(nums[2], nums[3]) + 1)
}

