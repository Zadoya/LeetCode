package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/json"
)

type Directory struct {
	Name    string			`json:"dir"`
	Files   []string		`json:"files"`
	Folders []Directory		`json:"folders"`
	Counter int 
	Virus   bool
}

func scanning() []string {
	var (
		in         *bufio.Scanner
		t, n       int64
		err        error
		data       []string
	)

	in = bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanLines)
	if in.Scan() {
		t, err = strconv.ParseInt(in.Text(), 10, 64)
		if err != nil {
			fmt.Println("Ошибка парсинга целого числа:", err)
			return nil
		}
	}

	data = make([]string, 0, t)
	for i := 0; i < int(t); i++ {
		if in.Scan() {
			n, err = strconv.ParseInt(in.Text(), 10, 64)
			if err != nil {
				fmt.Println("Ошибка парсинга целого числа:", err)
				return nil
			}
		}		

		var builder strings.Builder
		for j := 0; j < int(n) && in.Scan(); j++ {
			line := in.Text()
			builder.WriteString(line)
		}
		
		data = append(data, builder.String())
	}

	return data
}

func checkFolder(root Directory) int {
	for i := range root.Files {
		if strings.HasSuffix(root.Files[i], ".hack") {
			root.Virus = true
		}
	}
	if root.Virus {
		for i := range root.Folders {
			root.Folders[i].Virus = true
		}
		root.Counter += len(root.Files)
	}
	for i := range root.Folders {
		root.Counter += checkFolder(root.Folders[i])
	}

	return root.Counter
}

func main() {
	var out  *bufio.Writer

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	jsonData := scanning()
	if jsonData == nil {
		return
	}

	for i := range jsonData {
		var root Directory
		err := json.Unmarshal([]byte(jsonData[i]), &root)
		if err != nil {
			fmt.Fprintln(out, "Ошибка парсинга JSON:", err)
			return
		}
		fmt.Fprintln(out, checkFolder(root))
	}
}