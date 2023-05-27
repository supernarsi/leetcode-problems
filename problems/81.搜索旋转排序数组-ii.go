package problems

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	newNums := nums
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			newNums = nums[i:]
			newNums = append(newNums, nums[0:i]...)
			break
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if newNums[mid] == target {
			return true
		}
		if newNums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end
