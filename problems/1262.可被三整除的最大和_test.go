package problems

import "testing"

func TestMaxSumDivThree(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{3, 6, 5, 1, 8}, 18},
		{"t2", []int{4}, 0},
		{"t3", []int{1, 2, 3, 4, 4}, 12},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxSumDivThree(v.input); got != v.want {
				t.Errorf("%v, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
