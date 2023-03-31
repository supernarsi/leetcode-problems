package problems

/*
 * @lc app=leetcode.cn id=2367 lang=golang
 *
 * [2367] 算术三元组的数目
 */

// @lc code=start
func arithmeticTriplets(nums []int, diff int) (ans int) {
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j]-nums[i] == diff {
				for k := j + 1; k < len(nums); k++ {
					if nums[k]-nums[j] == diff {
						ans++
					}
				}
			}
		}
	}
	return
}

// @lc code=end
