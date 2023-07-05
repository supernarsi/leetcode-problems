package problems

import "testing"

func TestKItemsWithMaximumSum(t *testing.T) {
	tests := []struct {
		name   string
		input1 int
		input2 int
		input3 int
		inputK int
		want   int
	}{
		{"t1", 3, 2, 0, 2, 2},
		{"t2", 3, 2, 0, 4, 3},
		{"t3", 1, 3, 5, 6, -1},
		{"t4", 1, 0, 3, 2, 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := kItemsWithMaximumSum(v.input1, v.input2, v.input3, v.inputK); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
