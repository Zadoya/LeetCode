// условие задачи - https://leetcode.com/problems/shopping-offers/

package main

import (
	"fmt"
	"math"
	"sort"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	bitSpetial := make([][2]int, len(special))
	for i := range special {
		for j := 0; j < len(special[i]) - 1; j++ {
			bitSpetial[i][0] = bitSpetial[i][0] * int(math.Pow10(j)) + special[i][j]
		}
		bitSpetial[i][1] = special[i][len(special[i]) - 1]
	}
	for i := range price {
		bitSpetial = append(bitSpetial, [2]int{int(math.Pow10(i)), price[i]})
	}

	bitNeeds := 0
	for i := range needs {
		bitNeeds = bitNeeds * int(math.Pow10(i)) + needs[i]
	}

	minSum := math.MaxInt
	for i := range bitSpetial {
		tmp := bitNeeds
		sum := 0
		for j := i; j < len(bitSpetial); j++ {
			for tmp >= bitSpetial[j][0] {
				tmp -= bitSpetial[j][0]
				sum += bitSpetial[j][1]
			}
		}
		if tmp == 0 && sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

func main() {
	fmt.Println(shoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}))
}