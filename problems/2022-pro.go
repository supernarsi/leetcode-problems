package problems

func giveGem(gem []int, operations [][]int) int {
	for _, target := range operations {
		x := target[0]
		y := target[1]
		give := gem[x] / 2
		gem[x] /= 2
		gem[y] += give
	}
	minVal, maxVal := 10000, 0
	for _, val := range gem {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal - minVal
}
