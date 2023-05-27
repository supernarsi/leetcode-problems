package problems

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		inputM [][]int
		inputT int
		want   bool
	}{
		{
			name: "test1",
			inputM: [][]int{
				{1, 4, 7, 11, 15},
			},
			inputT: 3,
			want:   false,
		},
		{
			name: "test2",
			inputM: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 14, 25},
			},
			inputT: 3,
			want:   false,
		},
		{
			name: "test3",
			inputM: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 14, 25},
			},
			inputT: 8,
			want:   true,
		},
		{
			name: "test4",
			inputM: [][]int{
				{-5},
			},
			inputT: -5,
			want:   true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := searchMatrix(v.inputM, v.inputT); got != v.want {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
		})
	}
}
