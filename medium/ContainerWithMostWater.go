// условие задачи - https://leetcode.com/problems/container-with-most-water/description/

// Time  - O(n)
// Space - O(1)
func maxArea(height []int) int {
	maxSquare := 0
	left := 0
	right := len(height) - 1
	for left < right {
		currHeight := min(height[left], height[right])
		square := currHeight * (right - left)
		maxSquare = max(square, maxSquare)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxSquare
}
