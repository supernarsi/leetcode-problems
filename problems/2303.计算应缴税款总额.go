package problems

/*
 * @lc app=leetcode.cn id=2303 lang=golang
 *
 * [2303] 计算应缴税款总额
 */

// @lc code=start
func calculateTax(brackets [][]int, income int) (res float64) {
	brackets = append([][]int{{0, 0}}, brackets...)
	lv := 0
	for i, item := range brackets {
		if item[0] == income {
			lv = i
			break
		} else if item[0] > income {
			lv = i - 1
			break
		}
	}

	for k := 1; k <= lv; k++ {
		res += float64((brackets[k][0]-brackets[k-1][0])*brackets[k][1]) / 100
	}
	if income > brackets[lv][0] {
		res += float64((income-brackets[lv][0])*brackets[lv+1][1]) / 100
	}
	return
}

// @lc code=end
