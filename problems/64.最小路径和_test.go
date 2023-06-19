package problems

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"t1", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"t2", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := minPathSum(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
