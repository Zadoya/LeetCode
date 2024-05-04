// условие задачи - https://leetcode.com/problems/boats-to-save-people/description

func numRescueBoats(people []int, limit int) int {
	left := 0
	right := len(people) - 1
	res := 0
    
    sort.Ints(people)

	for left := 0,;right >= left {
		if people[right]+people[left] <= limit {
			left++
		}
		right--
		res++
	}

	return res
}