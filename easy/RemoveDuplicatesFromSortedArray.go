// условие задачи - https://leetcode.com/problems/remove-duplicates-from-sorted-array/description

package main

func removeDuplicates(nums []int) int {
    ptr1 := 0
	for ptr2 := 0; ptr2 < len(nums); ptr2++ {
		if nums[ptr1] != nums[ptr2] {
			ptr1++
			nums[ptr1] = nums[ptr2]
		}
	}
	return ptr1 + 1
}
