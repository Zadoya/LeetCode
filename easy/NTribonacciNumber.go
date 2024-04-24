// Условие задачи - https://leetcode.com/problems/n-th-tribonacci-number/

func tribonacci(n int) int { 
    tab := make([]int, n + 3)
	tab[0] = 0
	tab[1] = 1
	tab[2] = 1
	for i:= 3; i <= n; i++ {
		tab[i] = tab[i - 1] + tab[i - 2] + tab [i - 3]
	}
    return tab[n]
}