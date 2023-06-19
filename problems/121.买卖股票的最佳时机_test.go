package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"t2", []int{7, 6, 5, 4, 3, 1}, 0},
		{"t3", []int{4, 5, 10, 2, 3}, 6},
		{"t4", []int{4, 6, 7, 1, 5, 2, 4}, 4},
		{"t5", []int{3, 5, 7, 9, 1, 8, 12, 10, 20}, 19},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxProfit(v.input); got != v.want {
				t.Errorf("%v, want %v, got %v", v.name, v.input, got)
			}
		})
	}
}
