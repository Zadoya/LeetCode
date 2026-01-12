// условие задачи - https://leetcode.com/problems/remove-element/description/

package main

// бенчмарк показал, что аппенд самый не эффективный метод
// так как изначальный слайс не меняется, из-за передачи значения в функцию
// то имеет смысл применить метод двух указателей

func removeElement1(nums []int, val int) int {
    j := 0
    for i := range nums {
        if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
    }
    return j
}
