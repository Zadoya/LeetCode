// Условие задачи - https://leetcode.com/problems/sort-colors/description

// решение через count sort
func sortColors(nums []int)  {
	m := make([]int, 3)

	for _, num := range nums {
		m[num]++
	}

	j := 0
	for i := 0; i < 3; i++ {
		for ; m[i] > 0; m[i]-- {
			nums[j] = i
			j++
		}
	}
}

// оптимизированный способ
/*
func sortColors(nums []int)  {
    l, r := 0, len(nums)-1
    for i := 0; i < len(nums); i++ {
        for i >= l && i <= r && (nums[i] == 0 || nums[i] == 2) {
            if nums[i] == 0 {
                nums[i], nums[l] = nums[l], nums[i]
                l++
            } else if nums[i] == 2 {
                nums[i], nums[r] = nums[r], nums[i]
                r--
            }
        }
    }
}
*/