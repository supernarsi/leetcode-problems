package problems

/*
 * @lc app=leetcode.cn id=2347 lang=golang
 *
 * [2347] 最好的扑克手牌
 */

// @lc code=start
func bestHand(ranks []int, suits []byte) string {
	isColor := true
	color := suits[0]
	for _, c := range suits {
		if c != color {
			isColor = false
			break
		}
	}
	if isColor {
		return "Flush"
	}
	isPair := false
	dic := map[int]int{}
	for _, num := range ranks {
		dic[num]++
		if dic[num] == 3 {
			return "Three of a Kind"
		}
		if dic[num] == 2 {
			isPair = true
		}
	}

	if isPair {
		return "Pair"
	}
	return "High Card"

}

// @lc code=end
