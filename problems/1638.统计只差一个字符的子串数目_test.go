package problems

import "testing"

func TestCountSubstrings(t *testing.T) {
	test := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "test1",
			s:    "aba",
			t:    "baba",
			want: 6,
		},
		{
			name: "test2",
			s:    "ab",
			t:    "bb",
			want: 3,
		},
		{
			name: "test3",
			s:    "a",
			t:    "a",
			want: 0,
		},
		{
			name: "test4",
			s:    "abe",
			t:    "bbc",
			want: 10,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.s, tt.t); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
