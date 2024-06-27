// условие задачи - https://leetcode.com/problems/divisor-game/description/
// математическое объяснение этой задачи - https://assets.leetcode.com/users/images/ca8f8cef-a33c-4193-b0ca-0af326901c75_1603578539.455527.png

func divisorGame(n int) bool {
    if n % 2 == 0 {
        return true
    }
    return false
}