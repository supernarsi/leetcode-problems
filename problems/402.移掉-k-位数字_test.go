package problems

import "testing"

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		inputK int
		want   string
	}{
		{"t1", "1432219", 3, "1219"},
		{"t2", "10200", 1, "200"},
		{"t3", "10", 2, "0"},
		{"t4", "100101", 5, "0"},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := removeKdigits(v.input, v.inputK); got != v.want {
				t.Errorf("%s, want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
