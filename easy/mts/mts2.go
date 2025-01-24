package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin) // get the input
    var (
		str 	string
		counter uint8
	)
    if scanner.Scan() {
        str = scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
	for _, simbol := range str {
		switch {
		case (simbol == 'm' || simbol == 'M') && counter == 0:
			counter++
		case (simbol == 't' || simbol == 'T') && counter == 1:
			counter++
		case (simbol == 's' || simbol == 'S') && counter == 2:
			counter++
		}
	}
	if counter == 3 {
		fmt.Fprintln(os.Stdout, "1")
	} else {
		fmt.Fprintln(os.Stdout, "0")
	}
}