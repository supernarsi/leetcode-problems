package problems

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test1",
			input: []int{0, 1, 2, 3, 4, 6},
			want:  3,
		},
		{
			name:  "test2",
			input: []int{0, 1, 4, 10, 11, 12},
			want:  4,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := hIndex(v.input); got != v.want {
				t.Errorf("got %v, want %v", got, v.want)
			}
		})
	}
}
