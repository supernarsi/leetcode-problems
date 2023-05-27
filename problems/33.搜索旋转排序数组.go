package problems

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search33(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[left] {
			// mid 右侧有序，左侧无序
			if nums[mid] < target && target <= nums[right] {
				// 如果 target 在 mid 右侧
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// mid 左侧有序，右侧无序
			if nums[left] <= target && target < nums[mid] {
				// 如果 target 在 mid 左侧
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

// @lc code=end
