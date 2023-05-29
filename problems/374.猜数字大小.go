package problems

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 0, n

	for left <= right {
		mid := left + (right-left)/2
		ans := guess(mid)
		if ans == 0 {
			return mid
		} else if ans > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

// @lc code=end

func guess(n int) int {
	return 0
}
