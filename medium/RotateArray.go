// условие задачи - https://leetcode.com/problems/rotate-array/

package main

/*
func rotate(nums []int, k int)  {
	if k = k % len(nums); len(nums) == 1 || k == 0 {
		return
	}
    copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
*/

func rotate(nums []int, k int)  {
	if k = k % len(nums); len(nums) == 1 || k == 0 {
		return
	}

	reverse(&nums, 0, len(nums) - 1)
	reverse(&nums, 0, k - 1)
	reverse(&nums, k, len(nums) - 1)
}

func reverse(nums *[]int, start, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		end--
		start++
	}
}
