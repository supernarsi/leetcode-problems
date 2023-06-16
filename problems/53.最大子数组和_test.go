package problems

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"t2", []int{1}, 1},
		{"t3", []int{5, 4, -1, 7, 8}, 23},
		{"t4", []int{5, 4, -1, 7, 8, -1, -1, -1, 11, -2, -3, 1, 1}, 31},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxSubArray(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
