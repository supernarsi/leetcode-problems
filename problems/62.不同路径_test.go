package problems

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name   string
		inputM int
		inputN int
		want   int
	}{
		{"t1", 3, 2, 3},
		{"t2", 7, 3, 28},
		{"t3", 3, 3, 6},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := uniquePaths(v.inputM, v.inputN); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
