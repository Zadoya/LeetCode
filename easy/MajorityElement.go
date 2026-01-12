// уловие задачи - https://leetcode.com/problems/majority-element/

package main

type MajorNum struct {
    Num     int
    Counter int
}

func majorityElement(nums []int) int {
    mn := MajorNum{}
    for i := range nums {
        if mn.Counter == 0 {
            mn.Num = nums[i]
            mn.Counter++
        } else if mn.Num == nums[i] {
			mn.Counter++
		} else {
			mn.Counter--
		}
    }

	return mn.Num
}