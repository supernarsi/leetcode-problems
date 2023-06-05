package problems

import "testing"

func TestMaxRepOpt1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"t1", "ababa", 3},
		{"t2", "aaabaaa", 6},
		{"t3", "aaabbaaa", 4},
		{"t4", "aaa", 3},
		{"t5", "abcdef", 1},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := maxRepOpt1(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
