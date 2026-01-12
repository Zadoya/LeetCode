package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name  string
	Price int
}

type Data struct {
	Products map[int][]string
	CheckStr string
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

func scanArray(in *bufio.Reader, size int) (Product, error) {
	line, err := scanStr(in); 
	if err != nil {
		return Product{}, err
	}

	strs := strings.Split(line, " ")
	if len(strs) != int(size) {
		return Product{}, fmt.Errorf("incorrect number of arguments in string")
	}

	var result Product
	if len(strs[0]) > 10 {
		return Product{}, fmt.Errorf("incorrect len of product name")
	}
	result.Name = strs[0]

	num, err := strconv.ParseInt(strs[1], 10, 64)
	if err != nil {
		return Product{}, err
	}
	result.Price = int(num)

	return result, nil
}

func isAlfabet(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] < 'a' || str[i] > 'z' {
			return false
		} 
	}
	return true
}

func parseCSVString(line string) (map[int]string, error) {
	productsStr := strings.Split(line, ",")

	result := make(map[int]string, len(productsStr))
	for _, productStr := range productsStr {
		product := strings.Split(productStr, ":")
		if len(product) == 1 {
			return nil, fmt.Errorf("incorrect product format in valid string")
		}

		if len(product[0]) > 10 || !isAlfabet(product[0]) {
			return nil, fmt.Errorf("incorrect len of product name in valid string")
		}
		if len(product[1]) == 0 || product[1][0] < 49 || product[1][0] > 57 {
			return nil, fmt.Errorf("incorrect product price in valid string")
		}
		price, err := strconv.ParseInt(product[1], 10, 64)
		if err != nil {
			return nil, err
		}
		if name, ok := result[int(price)]; ok && name == product[0] { 
			return nil, fmt.Errorf("incorrect product price in valid string")
		}
		result[int(price)] = product[0]
	}

	return result, nil
}

func reading() ([]Data, error) {
	var (
		in         *bufio.Reader
		t, n       int
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
		data.Products = make(map[int][]string, n)
		for j := 0; j < n; j++ {
			product, err := scanArray(in, 2)
			if err != nil {
				return nil, err
			}
			data.Products[product.Price] = append(data.Products[product.Price], product.Name) 
		}

		if data.CheckStr, err = scanStr(in); err != nil {
			return nil, err
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
	case string:
		fmt.Fprintln(out, data.(string))
	case error:
		fmt.Fprintln(out, data.(error))
	}

}

func main() {
	dataset, err := reading()
	if err != nil {
		writing(err)
	}

	for _, data := range dataset {
		result, err := parseCSVString(data.CheckStr)
		if err != nil {
			writing("NO")
			continue
		}
		if len(result) != len(data.Products) {
			writing("NO")
			continue
		}
		var trigger bool
		/*
		for key, values := range data.Products {
			name, ok := result[key]; 
			if !ok {
				break
			}
			for _, value := range values {
				if name == value {
					trigger = true
					break
				}
			}
		}
		*/
		for key, name := range result {
			values, ok := data.Products[key]; 
			if !ok {
				break
			}
			for _, value := range values {
				if name == value {
					trigger = true
					break
				}
			}
		}
		if trigger {
			writing("NO")
		} else {
			writing("YES")
		}
	}
}