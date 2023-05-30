package problems

import (
	"testing"
)

func TestNumMatchingSubseq(t *testing.T) {
	tests := []struct {
		name       string
		inputStr   string
		inputWords []string
		want       int
	}{
		{"test1", "abcde", []string{"a", "bb", "acd", "ace"}, 3},
		{"test2", "abbcde", []string{"a", "bb", "acd", "ace"}, 4},
		{"test3", "dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}, 2},
		{"test4", "dsahjpjaujauhas", []string{"ahjpjau", "ja", "sasa", "sahs"}, 3},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := numMatchingSubseq(v.inputStr, v.inputWords); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
