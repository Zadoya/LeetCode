// условие задачи - https://leetcode.com/problems/search-insert-position/description

package main

func searchInsert(nums []int, target int) int {
	middle := len(nums) / 2
	switch {
	case target < nums[middle]:
		if len(nums) == 1 {
			return 0
		}
		return searchInsert(nums[:middle], target)
	case target > nums[middle]:
		if len(nums) == 1 {
			return 1
		}
		return middle + searchInsert(nums[middle:], target)
	default:
		return middle
	}
}

func main() {
	println(searchInsert([]int{3,5,6,10}, 1))
}