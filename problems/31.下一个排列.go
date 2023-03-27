package problems

import "fmt"

func nextPermutation(nums []int) {
	size := len(nums)
	if size == 1 {
		return
	}
	i, j := size-2, size-1
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i >= 0 {
		for j > i {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
	l, r := i+1, size-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	fmt.Println(nums)
	return
}
