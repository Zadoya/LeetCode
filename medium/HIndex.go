// условие задачи - https://leetcode.com/problems/h-index/description/

package main

import "fmt"

func mergeSort(arr []int, left, right int) []int {
	if left < right {
		middle := left + (right - left) / 2

		up := mergeSort(arr, left, middle)
		down := mergeSort(arr, middle + 1, right)

		u, d := 0, 0
		result := make([]int, 0,(right - left) * 2)
		for u < len(up) && d < len(down) {
			if up[u] < down[d] {
				result = append(result, up[u])
				u++
			} else {
				result = append(result, down[d])
				d++
			}
		}
		if u < len(up) {
			result = append(result, up[u:]...)
		}
		if d < len(down) {
			result = append(result, down[d:]...)
		}
		return result
	}
	return []int{arr[left]}
}

func sort(citations []int) []int {
	return mergeSort(citations, 0, len(citations) - 1)
}

func hIndex(citations []int) int {
	citations = sort(citations)

    i := 0
    for i < len(citations) && citations[i] == 0 {
        i++
    }

    citations = citations[i + 1:]

    h := min(citations[len(citations) - 1], len(citations))

    for i := len(citations) - 1; i > 0; i-- {
        if citations[i] <= len(citations) - i && citations[i] > 0 {
            return citations[i]
        }
    }

    return h
}

func main() {
	fmt.Println(hIndex([]int{0,0,2}))
}