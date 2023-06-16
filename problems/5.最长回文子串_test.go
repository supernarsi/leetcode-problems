package problems

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"t1", "babad", "bab"},
		{"t2", "cbbd", "bb"},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := longestPalindrome(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
