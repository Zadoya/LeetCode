// уловие задачи - https://leetcode.com/problems/two-sum/description/

// using two pointers -- O(n lon n)
func twoSumPointers(nums []int, target int) []int {
    sortedNums := make([][2]int, len(nums))
    for i, num := range nums {
        sortedNums[i] = [2]int{num, i}
    }

    sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i][0] < sortedNums[j][0]
	})

	left, right := 0, len(nums) - 1

	for left < right {
		sum := sortedNums[left][0] + sortedNums[right][0]
		if target == sum {
			return []int{sortedNums[left][1], sortedNums[right][1]}
		} else if target < sum {
			right--
		} else {
			left++
		}
	}
    return nil
}

// using hash table -- O(n)
func twoSumHash(nums []int, target int) []int {
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

