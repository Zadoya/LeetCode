package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadFileData(filename string) [][]int64 {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	size := (int64)(0)
	if scanner.Scan() {
		size, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	nums := make([][]int64, size)
	for i := 0; i < int(size) && scanner.Scan(); i++ {
		str := scanner.Text()
		nums[i] = make([]int64, size)
		for j := 0; j < int(size); j++ {
			switch str[j] {
			case '.':
				nums[i][j] = 0
			case '#':
				nums[i][j] = 1
			}
		}
	}
	return nums
}

func defineFunc(table [][]int64) {
	rectanglesHash := findRectangle(table, 1)
	rectangles := findRectangle(table, 0)
	if len(rectanglesHash) == 1 {
		fmt.Fprintln(os.Stdout, "I")
		return
	}

	if len(rectanglesHash) == 4 && len(rectangles) == 4 || len(rectangles) == 5 {
		fmt.Fprintln(os.Stdout, "O")
		return
	}

	switch len(rectangles) {
	case 1:
		if rectangles[0].TopLeft.X > 0 && rectangles[0].BottomRight.X < int64(len(table) - 1) &&
		   rectangles[0].TopLeft.Y > 0 && rectangles[0].BottomRight.Y < int64(len(table) - 1) {
			fmt.Fprintln(os.Stdout, "O")
		} else if rectangles[0].TopLeft.X == 0 && rectangles[0].BottomRight.X < int64(len(table) - 1) &&
		   		  rectangles[0].TopLeft.Y > 0 && rectangles[0].BottomRight.Y == int64(len(table) - 1) {
			fmt.Fprintln(os.Stdout, "L")
		} else if rectangles[0].TopLeft.X > 0 && rectangles[0].BottomRight.X < int64(len(table) - 1) &&
				  rectangles[0].TopLeft.Y > 0 && rectangles[0].BottomRight.Y == int64(len(table) - 1) {
			fmt.Fprintln(os.Stdout, "C")
	 	}
	case 2:
		if 0 < rectangles[0].TopLeft.Y && rectangles[1].TopLeft.Y  == rectangles[0].TopLeft.Y &&
		   rectangles[0].TopLeft.Y < rectangles[0].BottomRight.Y && rectangles[0].BottomRight.Y < rectangles[1].BottomRight.Y &&
		   rectangles[1].BottomRight.Y == int64(len(table) - 1) && int64(len(table) - 1) == rectangles[1].BottomRight.X &&
		   rectangles[1].BottomRight.X > rectangles[1].TopLeft.X  && rectangles[1].TopLeft.X > rectangles[0].BottomRight.X &&
		   rectangles[0].BottomRight.X > rectangles[0].TopLeft.X && rectangles[0].TopLeft.X > 0 {
			fmt.Fprintln(os.Stdout, "P")
		} else if 0 < rectangles[0].TopLeft.Y && rectangles[1].TopLeft.Y  == rectangles[0].TopLeft.Y &&
			rectangles[0].TopLeft.Y < rectangles[0].BottomRight.Y && rectangles[0].BottomRight.Y == rectangles[1].BottomRight.Y &&
			rectangles[1].BottomRight.Y < int64(len(table) - 1) && int64(len(table) - 1) == rectangles[1].BottomRight.X &&
			rectangles[1].BottomRight.X > rectangles[1].TopLeft.X && rectangles[1].TopLeft.X > rectangles[0].BottomRight.X &&
			rectangles[1].BottomRight.X > rectangles[1].TopLeft.X && rectangles[0].TopLeft.X == 0 {
				fmt.Fprintln(os.Stdout, "H")
		}
	default:
		fmt.Fprintln(os.Stdout, "X")
	}
}

type Rectangle struct {
	TopLeft Point
	BottomRight Point
}

type Point struct {
	X int64
	Y int64
}


func findRectangle(table [][]int64, sign int64) []Rectangle {
	var rectangles []Rectangle

	visited := make([][]bool, len(table))
	for i := range visited {
		visited[i] = make([]bool, len(table))
	}

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			// Ищем верхний левый угол прямоугольника из "."
			if table[i][j] == sign && !visited[i][j] {
				width, height := 0, 0

				// Определяем ширину прямоугольника
				for w := j; w < len(table) && table[i][w] == sign; w++ {
					width++
				}

				// Определяем высоту прямоугольника
				for h := i; h < len(table); h++ {
					allDots := true
					for w := 0; w < width; w++ {
						if table[h][j+w] != sign {
							allDots = false
							break
						}
					}
					if !allDots {
						break
					}
					height++
				}

				// Отмечаем все элементы прямоугольника как посещенные
				for x := i; x < i+height; x++ {
					for y := j; y < j+width; y++ {
						visited[x][y] = true
					}
				}

				// Сохраняем прямоугольник с координатами верхнего левого и нижнего правого углов
				rectangles = append(rectangles, Rectangle{
					TopLeft: Point{(int64)(i), (int64)(j)},
					BottomRight: Point{int64(i + height - 1), int64(j + width - 1)},
				})
			}
		}
	}

	return rectangles
}

func main() {
	table := loadFileData("input.txt")
	defineFunc(table)
}
