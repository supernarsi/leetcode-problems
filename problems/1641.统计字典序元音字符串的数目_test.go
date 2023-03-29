package problems

import "testing"

func TestCountVowelStrings(t *testing.T) {
	test := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test1",
			n:    1,
			want: 5,
		},
		{
			name: "test2",
			n:    2,
			want: 15,
		},
		{
			name: "test3",
			n:    33,
			want: 66045,
		},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := countVowelStrings(v.n)
			if got != v.want {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
		})
	}
}
