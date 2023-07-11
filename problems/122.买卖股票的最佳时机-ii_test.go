package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"t2", []int{1, 2, 3, 4, 5}, 4},
		{"t3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxProfit(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
