// условие задачи - https://leetcode.com/problems/shortest-common-supersequence/

package main

import (
	"fmt"
	"strings"
)
func shortestCommonSupersequence(str1 string, str2 string) string {
	if strings.Compare(str1, str2) == 0 {
		return str1
	}

	if str1 > str2 {
		str1, str2 = str2, str1
	}

	matrix := fillMatrix(str1, str2)

	for i := range matrix {
		fmt.Println(matrix[i])
	}

	lcs := findLCS(matrix, str1)
	

	var ptrStr1, ptrStr2, ptrLcs int
	
	result := make([]byte, 0, len(str1) + len(str2))
	for ptrLcs < len(lcs) {
		switch {
		case lcs[ptrLcs] == str1[ptrStr1] && lcs[ptrLcs] == str2[ptrStr2]:
			result = append(result, lcs[ptrLcs])
			ptrLcs++
			ptrStr1++
			ptrStr2++
		case lcs[ptrLcs] != str1[ptrStr1]:
			result = append(result, str1[ptrStr1])
			ptrStr1++
		case lcs[ptrLcs] != str2[ptrStr2]:
			result = append(result, str2[ptrStr2])
			ptrStr2++
		}
	}

	result = append(result, str1[ptrStr1:]...)
	result = append(result, str2[ptrStr2:]...)

	return string(result)
}

func findLCS(matrix [][] uint32, str1 string) string {
	maxlen  := matrix[len(matrix) - 1][len(matrix[0]) - 1]
	lcs := make([]byte, maxlen)

	j := len(matrix[0]) - 1
	for i := len(matrix) - 1; i > 0; i-- {
		for j > 0 {
			if matrix[i][j] > 0 {
				if matrix[i][j] > matrix[i][j - 1] && matrix[i][j] > matrix[i - 1][j] {
					lcs[matrix[i][j] - 1] = str1[i - 1]
					j--
					break
				} else if matrix[i - 1][j] > matrix[i][j - 1] {
					break
				}
			}
			j--
		}
	}

	return string(lcs)
}

func fillMatrix(str1, str2 string) [][]uint32 {
	matrix := make([][]uint32, len(str1) + 1)
	matrix[0] = make([]uint32, len(str2) + 1)

    for i, letter1 := range str1 {
		matrix[i+1] = make([]uint32, len(str2) + 1)
		for j, letter2 := range str2 {
			if letter2 == letter1 {
				matrix[i+1][j+1] = matrix[i][j] + 1
			} else {
				matrix[i+1][j+1] = max(matrix[i][j+1], matrix[i+1][j])
			}
		}
	}

	return matrix
}

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
	fmt.Println(shortestCommonSupersequence("bbbaaaba", "bbababbb"))
	fmt.Println(shortestCommonSupersequence("aabbabaa", "aabbbbbbaa"))
}