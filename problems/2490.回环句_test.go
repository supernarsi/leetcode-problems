package problems

import "testing"

func TestIsCircularSentence(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"t1", "leet", false},
		{"t2", "a", true},
		{"t3", "eetcode", true},
		{"t4", "Leetcode is cool", false},
		{"t5", "leetcode eis scool", true},
		{"t6", "leetcode eats soul", true},
		{"t7", "happy Leetcode", false},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := isCircularSentence(v.input); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
