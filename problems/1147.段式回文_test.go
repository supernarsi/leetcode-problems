package problems

import "testing"

func TestLongestDecomposition(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"", "ghiabcdefhelloadamhelloabcdefghi", 7},
		{"", "merchant", 1},
		{"", "antaprezatepzapreanta", 11},
		{"", "aaa", 3},
		{"", "elvtoelvto", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition(tt.text); got != tt.want {
				t.Errorf("longestDecomposition() = %v, want %v", got, tt.want)
			}
		})
	}
}
