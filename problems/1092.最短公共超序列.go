package problems

/*
 * @lc app=leetcode.cn id=1092 lang=golang
 *
 * [1092] 最短公共超序列
 */

// @lc code=start
func shortestCommonSupersequence(s, t string) string {
	n, m := len(s), len(t)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j < m; j++ {
		f[0][j] = j
	}
	for i := 1; i < n; i++ {
		f[i][0] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = min1092(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}

	ans := []byte{}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if s[i] == t[j] {
			ans = append(ans, s[i])
			i--
			j--
		} else if f[i+1][j+1] == f[i][j+1]+1 {
			ans = append(ans, s[i])
			i--
		} else {
			ans = append(ans, t[j])
			j--
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return s[:i+1] + t[:j+1] + string(ans)
}

func min1092(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end
