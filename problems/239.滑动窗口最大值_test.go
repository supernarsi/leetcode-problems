package problems

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	test := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "test1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "test2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "test3",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "test4",
			nums: []int{9, 11},
			k:    2,
			want: []int{11},
		},
		{
			name: "test5",
			nums: []int{4, -2},
			k:    2,
			want: []int{4},
		},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := maxSlidingWindow(v.nums, v.k)
			if len(got) != len(v.want) {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
			for i := 0; i < len(got); i++ {
				if got[i] != v.want[i] {
					t.Errorf("got: %v, want: %v", got, v.want)
				}
			}
		})
	}

}
