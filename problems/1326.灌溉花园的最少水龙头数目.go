package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=1326 lang=golang
 *
 * [1326] 灌溉花园的最少水龙头数目
 */

// @lc code=start
func MinTaps(n int, ranges []int) (ans int) {
	// 到达0~n每个坐标所需要的最小点数
	firstNum := 0
	if ranges[0] > 0 {
		firstNum = 1
	}
	minPointsNum := map[int]int{0: firstNum}

	for i, cover := range ranges {
		if cover <= 0 {
			continue
		}
		// 左边界
		l := i - cover
		// 右边界
		r := min1326(i+cover, n)

		if l <= 0 {
			// 更新所有点
			for j := 0; j <= r; j++ {
				minPointsNum[j] = 1
			}
		} else {
			if num, ok := minPointsNum[l]; ok && num > 0 {
				for j := l; j <= r; j++ {
					if jN, ok := minPointsNum[j]; ok {
						minPointsNum[j] = min1326(num+1, jN)
					} else {
						minPointsNum[j] = num + 1
					}
				}
			} else {
				for j := l + 1; j <= r; j++ {
					minPointsNum[j] = 0
				}
			}
		}
	}
	fmt.Println(minPointsNum)
	if min, ok := minPointsNum[n]; ok {
		return min
	} else {
		return -1
	}
}

func min1326(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
