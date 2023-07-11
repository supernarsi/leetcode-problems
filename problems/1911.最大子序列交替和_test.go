package problems

import "testing"

func TestMaxAlternatingSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int64
	}{
		{"t1", []int{4, 2, 5, 3}, 7},
		{"t2", []int{5, 6, 7, 8}, 8},
		{"t3", []int{6, 2, 1, 2, 4, 5}, 10},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxAlternatingSum(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.input, got)
			}
		})
	}
}
