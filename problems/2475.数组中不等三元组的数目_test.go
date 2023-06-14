package problems

import "testing"

func TestUnequalTriplets(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"t1", []int{4, 4, 2, 4, 3}, 3},
		{"t2", []int{4, 4, 4, 4, 3}, 0},
		{"t3", []int{1, 2, 3, 4, 5}, 10},
		{"t4", []int{1, 2, 3, 4, 5, 5}, 16},
		{"t5", []int{1, 2, 3, 4}, 4},
		{"t6", []int{1, 3, 1, 2, 4}, 7},
		{"t7", []int{274, 214, 8, 210, 87, 477, 467, 432, 246, 308, 10, 202, 123, 226, 98, 103, 367, 5, 188, 242, 394, 4, 103, 442, 159, 300, 87, 280, 277, 473, 137, 268, 377, 184, 415, 272, 251, 286, 477, 402, 265, 330, 424, 90, 53, 345, 126, 6, 81, 212, 197, 145, 41, 361, 198, 464, 151, 158, 220, 458, 115}, 35813},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := unequalTriplets(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
