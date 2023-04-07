package problems

import "testing"

func TestNumMovesStonesII(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"", []int{7, 4, 9}, []int{1, 2}},
		{"", []int{6, 5, 4, 3, 10}, []int{2, 3}},
		{"", []int{100, 101, 104, 102, 103}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStonesII(tt.args); !equal(got, tt.want) {
				t.Errorf("numMovesStonesII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
