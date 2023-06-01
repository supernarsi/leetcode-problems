package problems

import "testing"

func TestMaximumTastiness(t *testing.T) {
	tests := []struct {
		name       string
		inputPrice []int
		inputK     int
		want       int
	}{
		{"t1", []int{13, 5, 1, 8, 21, 2}, 3, 8},
		{"t2", []int{1, 3, 1}, 2, 2},
		{"t2", []int{7, 7, 7, 7}, 2, 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maximumTastiness(v.inputPrice, v.inputK); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
