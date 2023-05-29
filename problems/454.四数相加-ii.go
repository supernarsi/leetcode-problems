package problems

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countAB := make(map[int]int)
	ans := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			countAB[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			ans += countAB[-c-d]
		}
	}
	return ans
}

// @lc code=end
