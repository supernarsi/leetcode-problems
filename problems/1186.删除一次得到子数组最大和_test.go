package problems

import "testing"

func TestMaximumSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{1, -2, 0, 3}, 4},
		{"t2", []int{1, -2, -2, 3}, 3},
		{"t3", []int{-1, -1, -1, -1}, -1},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maximumSum(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.input, got)
			}
		})
	}
}
