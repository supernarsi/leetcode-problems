package problems

import "testing"

func TestMaximumEvenSplit(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  []int64
	}{
		{"t1", 0, []int64{}},
		{"t2", 1, []int64{}},
		{"t3", 2, []int64{2}},
		{"t4", 4, []int64{4}},
		{"t5", 6, []int64{2, 4}},
		{"t6", 8, []int64{2, 6}},
		{"t7", 12, []int64{2, 4, 6}},
		{"t8", 18, []int64{2, 4, 12}},
		{"t9", 20, []int64{2, 4, 6, 8}},
		{"t10", 28, []int64{2, 4, 6, 16}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := maximumEvenSplit(v.input)
			if len(got) != len(v.want) {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
			for i, item := range got {
				if item != v.want[i] {
					t.Errorf("%s, want %v, got %v", v.name, v.want, got)
				}
			}
		})
	}
}
