package problems

import "testing"

func TestMatrixSum(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"t1", [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}, 15},
		{"t2", [][]int{{1}}, 1},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := matrixSum(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
