package problems

import "testing"

func TestJump(t *testing.T) {
	t.Log(jump([]int{2, 3, 1, 1, 4}) == 2)
	t.Log(jump([]int{2, 3, 0, 1, 4}) == 2)
	t.Log(jump([]int{1, 2, 3}) == 2)
	t.Log(jump([]int{1, 2, 1, 1, 1}) == 3)
	t.Log(jump([]int{1, 1, 1, 1, 1}) == 4)
}
