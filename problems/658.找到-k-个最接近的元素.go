package problems

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) (ans []int) {
	for i := 0; i < k; i++ {
		pivot := sort.SearchInts(arr, x)
		if pivot < 0 {
			// 在最左侧，将右侧第一个加入答案
			ans = append(ans, arr[0])
			arr = arr[1:]
		} else if pivot == len(arr) {
			// 在最右侧，将左侧第一个加入答案
			ans = append(ans, arr[len(arr)-1])
			arr = arr[:len(arr)-1]
		} else {
			// 在中间
			if arr[pivot] == x {
				// 刚好命中，加入答案中
				ans = append(ans, x)
				// 将该值从原数组移除
				arr = append(arr[:pivot], arr[pivot+1:]...)
			} else {
				// 判断相邻两个元素哪个更近
				lVal, rVal := math.MinInt32, arr[pivot]
				if pivot > 0 {
					lVal = arr[pivot-1]
				}
				if x-lVal <= rVal-x {
					// 左边的更近
					ans = append(ans, lVal)
					arr = append(arr[:pivot-1], arr[pivot:]...)
				} else {
					// 右边的更近
					ans = append(ans, rVal)
					if pivot == 0 {
						arr = arr[1:]
					} else {
						arr = append(arr[0:pivot], arr[pivot+1:]...)
					}
				}
			}
		}
	}
	sort.Ints(ans)
	return ans
}

// @lc code=end
