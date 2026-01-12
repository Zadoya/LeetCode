package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	strs []string
	odd  []string
	even []string
}

func scanStr(in *bufio.Reader) (string, error) {
	str, err := in.ReadString('\n')

	return str[:len(str)-1], err
}

func scanInt(in *bufio.Reader) (int, error) {
	line, err := scanStr(in); 
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(line, 10, 64)

	return int(num), err
}

func reading() ([]Data, error) {
	var (
		in         *bufio.Reader
		t, n       int
		str        string
		err        error
		dataset    []Data
	)

	in = bufio.NewReader(os.Stdin)

	if t, err = scanInt(in); err != nil {
		return nil, err
	}

	dataset = make([]Data, 0, t)
	for i := 0; i < t; i++ {
		var data Data
		if n, err = scanInt(in); err != nil {
			return nil, err
		}
		data.strs = make([]string, 0, n)
		for j := 0; j < n; j++ {
			if str, err = scanStr(in); err != nil {
				return nil, err
			}
			data.strs = append(data.strs, str)
		}

		dataset = append(dataset, data)
	}

	return dataset, nil
}

func writing(data any) {
	var out  *bufio.Writer

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	switch data.(type) {
	case int:
		fmt.Fprintln(out, data.(int))	
	case error:
		fmt.Fprintln(out, data.(error))
	}

}

func oddAndEven(str string) (string, string) {
	odd  := make([]byte, 0, len(str)/2+1)
	even := make([]byte, 0, len(str)/2+1)
	turn := true
	for i := 0; i < len(str); i++ {
		if turn {
			even = append(even, str[i])
			turn = false
		} else {
			odd = append(odd, str[i])
			turn = true
		}
	}
	return string(odd), string(even)
}

func factorial(num int) int {
    result := 1
    for i := 2; i <= num; i++ {
        result *= i
    }
    return result
}

func main() {
	dataset, err := reading()
	if err != nil {
		writing(err)
	}

	for _, data := range dataset {
		for _, str := range data.strs {
			odd, even := oddAndEven(str)
			data.odd = append(data.odd, odd)
			data.even = append(data.even, even) 
		}
		counterOdd := make(map[string]int)
		maxOdd := 0
		counterEven := make(map[string]int)
		maxEven := 0
		for i := range data.odd {
			_, okOdd := counterOdd[data.odd[i]]
			_, okEven := counterEven[data.even[i]]
			if okEven && okOdd {
				counterOdd[data.odd[i]]++
				continue
			}
			counterOdd[data.odd[i]]++
			counterEven[data.odd[i]]++
			maxOdd = max(maxOdd, counterOdd[data.odd[i]])
			maxEven = max(maxEven, counterEven[data.even[i]])
		}
		if maxEven == 1 {
			maxEven = 0
		}
		if maxOdd == 1 {
			maxOdd = 0
		}
		writing(maxEven + maxOdd)
		//writing(factorial(maxEven) + factorial(maxOdd))
	}
}