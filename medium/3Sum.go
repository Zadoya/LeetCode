// условие задачи - https://leetcode.com/problems/3sum/description

package main

import (
	"fmt"
	"slices"
)
// time complexity = O(n^2), space = O(1)
func threeSum(nums []int) [][]int {
    res := make([][]int, 0, len(nums))

    slices.Sort(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }

		left  := i + 1
		right := len(nums) - 1
		for left < right {
			switch {
			case nums[left] + nums[right] + nums[i] == 0:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for left < right && nums[left] == nums[left - 1] {
					left++
				}
				for left < right && nums[right] == nums[right + 1] {
					right--
				}
			case nums[left] + nums[right] + nums[i] > 0:
				right--
			case nums[left] + nums[right] + nums[i] < 0:
				left++
			}
		}
	}
	return res
}

// попробовать через 2Sum
func twoSum(nums []int, target int) []int {
	seem := make(map[int]int, len(nums))
	for i := range nums {
		diff := target - nums[i]
		if j, ok := seem[diff]; ok {
			return []int{j, i}
		} else {
			seem[nums[i]] = i
		}
	}
	return nil
}


func main() {
	fmt.Println(threeSum([]int{-2,-1,1,2}))
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4,-2,-3,3,0,4}))
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum([]int{0,0,0,0,0,0,0}))
}