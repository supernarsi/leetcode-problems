package problems

import "testing"

func TestArithmeticTriplets(t *testing.T) {
	t.Log(arithmeticTriplets([]int{2, 4, 6, 8, 10}, 2))
	t.Log(arithmeticTriplets([]int{9, 4, 7, 2, 10}, 3))
	t.Log(arithmeticTriplets([]int{1, 1, 1, 1}, 0))
}
