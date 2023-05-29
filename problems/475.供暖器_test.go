package problems

import "testing"

func TestFindRadius(t *testing.T) {
	tests := []struct {
		name      string
		inHouses  []int
		inHeaters []int
		want      int
	}{
		{"test1", []int{1, 2, 3, 4}, []int{2}, 2},
		{"test2", []int{1, 12, 23, 44}, []int{2, 5, 9}, 35},
		{"test3", []int{1, 2, 3, 4, 5, 6}, []int{2, 6}, 2},
		{"test4", []int{1}, []int{12}, 11},
		{"test5", []int{1, 100}, []int{30}, 70},
		{"test6", []int{1, 2, 3, 4, 20, 38, 192}, []int{4, 33, 90}, 102},
		{"test7", []int{1, 5}, []int{2}, 3},
		{"test8", []int{10, 25}, []int{1}, 24},
		{"test9", []int{10, 25}, []int{11, 20, 22}, 3},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := findRadius(v.inHouses, v.inHeaters); got != v.want {
				t.Errorf("%s want %v, got %v", v.name, v.want, got)
			}
		})
	}
}
