package problems

import "testing"

func TestMctFromLeafValues(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{6, 2, 4}, 32},
		{"t2", []int{4, 11}, 44},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := mctFromLeafValues(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
