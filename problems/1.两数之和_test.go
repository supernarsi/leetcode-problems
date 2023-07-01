package problems

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{"t1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"t2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"t3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := twoSum(v.input, v.target)
			if len(got) != len(v.want) || got[0] != v.want[0] || got[1] != v.want[1] {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
