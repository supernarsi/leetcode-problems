package problems

import "testing"

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		inputNew []int
		want     [][2]int
	}{
		{"t1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, []int{1, 5}, [][2]int{{1, 6}, {8, 10}, {15, 18}}},
		{"t2", [][]int{{1, 2}, {2, 6}}, []int{1, 9}, [][2]int{{1, 9}}},
		{"t3", [][]int{{1, 2}, {2, 3}, {3, 6}}, []int{1, 5}, [][2]int{{1, 6}}},
		{"t4", [][]int{{1, 5}, {2, 3}, {3, 6}}, []int{1, 4}, [][2]int{{1, 6}}},
		{"t5", [][]int{{1, 5}, {2, 13}, {23, 24}}, []int{12, 13}, [][2]int{{1, 13}, {23, 24}}},
		{"t6", [][]int{{1, 5}, {2, 13}, {23, 24}}, []int{12, 15}, [][2]int{{1, 15}, {23, 24}}},
		{"t7", [][]int{{1, 5}}, []int{1, 3}, [][2]int{{1, 5}}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := insert(v.input, v.inputNew)
			if len(got) != len(v.want) {
				t.Errorf("got %v, want %v", got, v.want)
			} else {
				for i, g := range got {
					if g[0] != v.want[i][0] || g[1] != v.want[i][1] {
						t.Errorf("got %v, want %v", got, v.want)
					}
				}
			}
		})
	}
}
