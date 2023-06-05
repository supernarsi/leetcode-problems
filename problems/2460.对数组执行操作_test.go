package problems

import "testing"

func TestAapplyOperations(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"t1", []int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{"t2", []int{0, 1}, []int{1, 0}},
		{"t3", []int{1, 0}, []int{1, 0}},
		{"t4", []int{1, 1}, []int{2, 0}},
		{"t5", []int{1, 1, 0}, []int{2, 0, 0}},
		{"t6", []int{1, 1, 0, 2, 2, 1}, []int{2, 4, 1, 0, 0, 0}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := applyOperations(v.input)
			if len(got) != len(v.want) {
				t.Errorf("want %v, got %v", v.want, got)
			}
			for i, e := range got {
				if e != v.want[i] {
					t.Errorf("want %v, got %v", v.want, got)
				}
			}
		})
	}
}
