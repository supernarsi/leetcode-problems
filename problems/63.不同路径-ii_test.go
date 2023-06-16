package problems

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"t1", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"t2", [][]int{{0, 1}, {0, 0}}, 1},
		{"t3", [][]int{{1, 0}}, 0},
		{"t4", [][]int{{0}, {0}}, 1},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
