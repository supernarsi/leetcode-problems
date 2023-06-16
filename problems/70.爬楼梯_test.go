package problems

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"t1", 3, 3},
		{"t2", 2, 2},
		{"t3", 1, 1},
		{"t4", 4, 5},
		{"t5", 10, 89},
		{"t6", 45, 1836311903},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := climbStairs(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
