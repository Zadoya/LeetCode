// условие задачи - https://leetcode.com/problems/maximize-happiness-of-selected-children/description

func maximumHappinessSum(happiness []int, k int) int64 {
	var sum int64


	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	for i := 0; i < k; i++ {
		if happiness[i] - i > 0 {
			sum += int64(happiness[i] - i)
		} else {
			return sum
		}
	}
	return sum
}