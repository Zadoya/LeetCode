// условие задачи - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package main

func twoSum(numbers []int, target int) []int {
    first := 0
    second := len(numbers) - 1
    for first < second {
        if numbers[first] + numbers[second] == target {
            return []int{first + 1, second + 1} 
        } else if numbers[first] + numbers[second] > target {
            second--
        } else if numbers[first] + numbers[second] < target {
            first++
        }
    }

    return []int{0, 0}
}