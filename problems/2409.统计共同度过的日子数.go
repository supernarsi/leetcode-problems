package problems

import "strconv"

/*
 * @lc app=leetcode.cn id=2409 lang=golang
 *
 * [2409] 统计共同度过的日子数
 */

// @lc code=start
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	datesOfMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	prefixSum := make([]int, 1)
	for _, date := range datesOfMonths {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+date)
	}

	arriveAliceDay := calculateDayOfYear(arriveAlice, prefixSum)
	leaveAliceDay := calculateDayOfYear(leaveAlice, prefixSum)
	arriveBobDay := calculateDayOfYear(arriveBob, prefixSum)
	leaveBobDay := calculateDayOfYear(leaveBob, prefixSum)

	return max2409(0, min2409(leaveAliceDay, leaveBobDay)-max2409(arriveAliceDay, arriveBobDay)+1)
}

func calculateDayOfYear(day string, prefixSum []int) int {
	month, _ := strconv.Atoi(day[:2])
	date, _ := strconv.Atoi(day[3:])
	return prefixSum[month-1] + date
}

func max2409(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2409(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
