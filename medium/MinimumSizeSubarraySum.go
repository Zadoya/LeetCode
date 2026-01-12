// условие задачи - https://leetcode.com/problems/minimum-size-subarray-sum/

package main

func minSubArrayLen(target int, nums []int) int {
	start, end, sum := 0, 0, 0
	answer := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			answer = min(answer, end - start + 1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if answer > len(nums) {
		return 0
	}
	return answer
}