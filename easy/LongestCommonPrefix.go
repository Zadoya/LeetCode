// условие задачи - https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; len(prefix) > 0 && i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]
	}

	return prefix
}

func main () {
	strs1 := []string{"flower","flow","flight"}
	strs2 := []string{"dog","racecar","car"}

	fmt.Println(longestCommonPrefix(strs1), strs1)
	fmt.Println(longestCommonPrefix(strs2))
}