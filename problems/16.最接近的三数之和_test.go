package problems

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		inputN []int
		inputT int
		want   int
	}{
		{"t1", []int{-1, 2, 1, -4}, 1, 2},
		{"t2", []int{0, 0, 0}, 1, 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := threeSumClosest(v.inputN, v.inputT); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
