package problems

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name        string
		inputNums   []int
		inputTarget int
		want        bool
	}{
		{
			name:        "test1",
			inputNums:   []int{5, 6, 7, 0, 2, 4},
			inputTarget: 3,
			want:        false,
		},
		{
			name:        "test2",
			inputNums:   []int{5, 6, 7, 1, 2, 4},
			inputTarget: 2,
			want:        true,
		},
		{
			name:        "test3",
			inputNums:   []int{2, 5, 6, 0, 0, 1, 2},
			inputTarget: 3,
			want:        false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := search(v.inputNums, v.inputTarget); got != v.want {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
		})
	}
}
