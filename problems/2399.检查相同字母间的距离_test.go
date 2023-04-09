package problems

import "testing"

func TestCheckDistances(t *testing.T) {
	type args struct {
		s        string
		distance []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"abab", []int{1, 1, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, true},
		{"t2", args{"acac", []int{1, 3, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, false},
		{"t3", args{"abba", []int{2, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, true},
		{"t4", args{"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDistances(tt.args.s, tt.args.distance); got != tt.want {
				t.Errorf("checkDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
