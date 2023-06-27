package problems

import "testing"

func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		name  string
		input [7]int
		want  bool
	}{
		{"t1", [7]int{1, 0, 0, 1, -1, 3, 1}, true},
		{"t2", [7]int{1, 1, 1, 1, -3, 2, -1}, false},
		{"t3", [7]int{1, 0, 0, -1, 0, 0, 1}, true},
	}

	for _, v := range tests {
		if got := checkOverlap(v.input[0], v.input[1], v.input[2], v.input[3], v.input[4], v.input[5], v.input[6]); got != v.want {
			t.Errorf("%v, want %v, got %v", v.name, v.want, got)
		}
	}
}
