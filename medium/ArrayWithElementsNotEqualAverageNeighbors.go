// условие задачи - https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/description/

// time - O(Nlogn)
// space - O(1)
func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums) - 1; i += 2 {
        nums[i], nums[i + 1] = nums[i + 1], nums[i]
    } 
    return nums
}

/* 
// соседи должны быть либо больше, либо меньше
// time  - O(n)
// space - O(1) 
func rearrangeArray(nums []int) []int {
	i := 1
	c := false
	for i < len(nums)-1 {
		if c {
			if nums[i] > nums[i-1] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
			if nums[i] > nums[i+1] {
				nums[i+1], nums[i] = nums[i], nums[i+1]
			}
		} else {
			if nums[i] < nums[i-1] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
			if nums[i] < nums[i+1] {
				nums[i+1], nums[i] = nums[i], nums[i+1]
			}
		}
		i++
		c = !c
	}
	return nums
}
*/