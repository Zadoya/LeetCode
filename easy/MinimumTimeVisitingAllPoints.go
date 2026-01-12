// условие задачи - https://leetcode.com/problems/minimum-time-visiting-all-points/description/

package main

func minTimeToVisitAllPoints(points [][]int) int {
    sum := 0
    yDiff := 0
    xDiff := 0

    for i := 1; i < len(points); i++ {
        yDiff = abs(points[i][0] - points[i-1][0])
        xDiff = abs(points[i][1] - points[i-1][1])
        sum += max(yDiff, xDiff)
    }

    return sum
}

func abs(point int) int {
    if point < 0 {
        return point * -1
    }

    return point
}