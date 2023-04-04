package problems

import "testing"

func TestMergeStones(t *testing.T) {
	t.Log(mergeStones([]int{3, 2, 4, 1}, 2))
	t.Log(mergeStones([]int{3, 5, 1, 2, 6}, 3))
	t.Log(mergeStones([]int{3, 2, 4, 1}, 3))
	t.Log(mergeStones([]int{3, 5, 1, 2, 6}, 2))
}
