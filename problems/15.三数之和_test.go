package problems

import (
	"strconv"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{"t1", []int{-1, 0, 1, 2, -1, 4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{"t2", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"t3", []int{0, 1, 1}, [][]int{}},
		{"t4", []int{-4, -3, -2, -1, 0, 1, 2, 2, 4, 3, 5}, [][]int{{-4, -1, 5}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-3, -2, 5}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
	}

	toString := func(input []int) string {
		ans := ""
		for _, v := range input {
			ans += strconv.Itoa(v)
		}
		return ans
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := threeSum(v.input)
			if len(got) != len(v.want) {
				t.Errorf("%s, want %v, got %v", v.name, len(v.want), len(got))
			}
			gotExist := map[string]bool{}
			for _, item := range got {
				gotExist[toString(item)] = true
			}
			for _, item := range v.want {
				if !gotExist[toString(item)] {
					t.Errorf("%s, want %v, got %v", v.name, v.want, got)
				}
			}
		})
	}
}
