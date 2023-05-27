package problems

import "testing"

func TestSearch33(t *testing.T) {
	tests := []struct {
		name         string
		inputNums    []int
		intputTarget int
		want         int
	}{
		{
			name:         "test1",
			inputNums:    []int{},
			intputTarget: 1,
			want:         -1,
		},
		{
			name:         "test2",
			inputNums:    []int{4, 5, 6, 7, 0, 1, 2},
			intputTarget: 0,
			want:         4,
		},
		{
			name:         "test3",
			inputNums:    []int{4, 5, 6, 7, 0, 1, 2},
			intputTarget: 3,
			want:         -1,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := search33(v.inputNums, v.intputTarget); got != v.want {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
		})
	}
}
