package problems

import "testing"

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		name  string
		input [4][]int
		want  int
	}{
		{"test1", [4][]int{{1, 2}, {-2, -1}, {-1, 2}, {0, 2}}, 2},
		{"test2", [4][]int{{1, 2, 3}, {-2, -1, -3}, {-1, 1, 2}, {0, -1, 2}}, 13},
		{"test3", [4][]int{{0}, {0}, {0}, {0}}, 1},
		{"test4", [4][]int{{0}, {0}, {0}, {1}}, 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := fourSumCount(v.input[0], v.input[1], v.input[2], v.input[3]); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
