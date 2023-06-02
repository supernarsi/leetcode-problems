package problems

import "testing"

func TestVowelStrings(t *testing.T) {
	tests := []struct {
		name       string
		inputWords []string
		inputPos   [][]int
		want       []int
	}{
		{"t1", []string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}, []int{2, 3, 0}},
		{"t2", []string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}, []int{3, 2, 1}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := vowelStrings(v.inputWords, v.inputPos)
			if len(got) != len(v.want) {
				t.Errorf("want %v, got %v", v.want, got)
			} else {
				for i, eGot := range got {
					if eGot != v.want[i] {
						t.Errorf("want %v, got %v", v.want, got)
					}
				}
			}
		})
	}
}
