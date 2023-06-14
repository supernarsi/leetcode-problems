package problems

import "testing"

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"t1", [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{"t2", [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := equalPairs(v.input); got != v.want {
				t.Errorf("%v, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
