// условие задачи - https://leetcode.com/problems/summary-ranges/description/

package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
    var (
		strRange string
		isRange  bool
	)

    result := make([]string, 0, len(nums))
    for i := 0; i < len(nums); i++ {
		if i == len(nums) - 1 {
            if isRange {
                strRange += strconv.Itoa(nums[i])
            } else {
                strRange = strconv.Itoa(nums[i])
            }
			result = append(result, strRange)
			break
		}
        if !isRange {
            if nums[i]+1 != nums[i+1] {
                result = append(result, strconv.Itoa(nums[i]))
            } else {
                isRange = true
				strRange = strconv.Itoa(nums[i]) + "->"
            }
        } else {
			if nums[i]+1 != nums[i+1] {
				isRange = false
				strRange += strconv.Itoa(nums[i])
                result = append(result, strRange)
            }
		}
    }
	
	return result
}
