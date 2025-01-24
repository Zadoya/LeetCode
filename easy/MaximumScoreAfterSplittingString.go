// условие задачи - https://leetcode.com/problems/maximum-score-after-splitting-a-string/

package main

func maxScore(s string) int {
	var (
		zeros, ones int16
	)

	sumZeros := make([]int16, len(s))
    sumOnes  := make([]int16, len(s))

    for i := range s {
		if s[i] == '0' {
			zeros++
		}
		sumZeros[i] = zeros
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			ones++
		}
		sumOnes[i] = ones
	}

	max := int16(0)
	for i := 0; i < len(s) - 1; i++ {
		tmp := sumOnes[i + 1] + sumZeros[i]
		if tmp > max {
			max = tmp
		}
	}
	
	return int(max)
}