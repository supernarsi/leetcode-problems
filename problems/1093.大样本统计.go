package problems

/*
 * @lc app=leetcode.cn id=1093 lang=golang
 *
 * [1093] 大样本统计
 */

// @lc code=start
func sampleStats(count []int) []float64 {
	n := len(count)
	total := 0
	for i := 0; i < n; i++ {
		total += count[i]
	}
	mean := 0.0
	median := 0.0
	minimum := 256
	maxnum := 0
	mode := 0

	left := (total + 1) / 2
	right := (total + 2) / 2
	cnt := 0
	maxfreq := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(count[i]) * int(i)
		if count[i] > maxfreq {
			maxfreq = count[i]
			mode = i
		}
		if count[i] > 0 {
			if minimum == 256 {
				minimum = i
			}
			maxnum = i
		}
		if cnt < right && cnt+count[i] >= right {
			median += float64(i)
		}
		if cnt < left && cnt+count[i] >= left {
			median += float64(i)
		}
		cnt += count[i]
	}
	mean = float64(sum) / float64(total)
	median = median / 2.0
	return []float64{float64(minimum), float64(maxnum), mean, median, float64(mode)}
}

// @lc code=end
