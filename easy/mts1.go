package main
import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin) // get the input
	response := make([]int64, 0, 3)
    for i:= 0; i < 3; i++ {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "parsing input:", err)
				return
			}
        	response = append(response, num)
		} else {
			fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err().Error())
			return
		}
    }
	if response[0] > 1000 || response[1] > 1000 || response[2] > 10000 {
		fmt.Fprintln(os.Stderr, "input data is out of bounds")
		return
	}
	if response[2] <= 100 {
		fmt.Fprintln(os.Stdout, response[0])
	} else {
		fmt.Fprintln(os.Stdout, response[0] + (response[2] - 100) * response[1])
	}
}