package problems

import "testing"

func TestSmallestSufficientTeam(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want []int
	}{
		{"", [][]string{{"java", "nodejs", "reactjs"}, {"nodejs", "reactjs"}, {"java", "nodejs"}, {"nodejs", "reactjs"}, {"csharp", "nodejs"}, {"nodejs", "reactjs"}, {"csharp", "nodejs"}, {"java", "nodejs"}, {"csharp", "nodejs"}, {"csharp", "nodejs"}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSufficientTeam(tt.args[0], tt.args[1:]); !equal(got, tt.want) {
				t.Errorf("smallestSufficientTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
