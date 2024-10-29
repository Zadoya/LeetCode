package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	coordinates := make([]int64, 0, 6)

	i := 0
	for scanner.Scan() && i < 6 {
		coordinates[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		if coordinates[i] < -100 || coordinates[i] > 100 {
			fmt.Fprintf(os.Stderr, "uncorrect input data")
		}
		i++
	}
	fmt.Fprintln(os.Stdout, whichSide(coordinates))
}

func whichSide(coordinates []int64) string {
	if coordinates[5] > coordinates[1] && coordinates[5] < coordinates[3] {
		if coordinates[4] < coordinates[0] {
			return "W"
		} else if coordinates[4] > coordinates[2] {
			return "E"
		}
	}
	if coordinates[4] > coordinates[0] && coordinates[4] < coordinates[2] {
		if coordinates[5] < coordinates[1] {
			return "S"
		} else if coordinates[5] > coordinates[3] {
			return "N"
		}
	}

	if coordinates[4] < coordinates[0] && coordinates[5] < coordinates[1] {
		return "SW"
	}
	if coordinates[4] < coordinates[0] && coordinates[5] > coordinates[3] {
		return "NW"
	}
	if coordinates[4] > coordinates[0] && coordinates[5] < coordinates[1] {
		return "SE"
	}
	if coordinates[4] > coordinates[0] && coordinates[5] > coordinates[3] {
		return "NE"
	}
	return ""
}