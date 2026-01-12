//	условие задачи - https://leetcode.com/problems/array-partition

package main

func partition(nums []int, left, right int) ([]int, int) {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return nums, i
}

func qsortInternal(nums []int, left, right int) []int {
	if left < right {
		nums, pivot := partition(nums, left, right)
		nums = qsortInternal(nums, left, pivot - 1)
		nums = qsortInternal(nums, pivot + 1, right)
	}
	return nums
}

func qsort(nums []int) []int {
	return qsortInternal(nums, 0, len(nums) - 1)
}

func arrayPairSum(nums []int) int {
    nums = qsort(nums)  //  slices.Sort(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
