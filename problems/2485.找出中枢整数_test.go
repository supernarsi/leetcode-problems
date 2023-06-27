package problems

import "testing"

func TestPivotInteger(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"t1", 8, 6},
		{"t2", 1, 1},
		{"t3", 2, -1},
		{"t4", 3, -1},
		{"t5", 4, -1},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := pivotInteger(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
