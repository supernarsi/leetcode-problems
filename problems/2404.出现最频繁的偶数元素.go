package problems

/*
 * @lc app=leetcode.cn id=2404 lang=golang
 *
 * [2404] 出现最频繁的偶数元素
 */

// @lc code=start
func mostFrequentEven(nums []int) (ans int) {
	m := make(map[int]int)
	for _, v := range nums {
		if v%2 == 0 {
			m[v]++
		}
	}
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			ans = k
		} else if v == max && k < ans {
			ans = k
		}
	}
	return
}

// @lc code=end
