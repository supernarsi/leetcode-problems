package problems

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][2]int
	}{
		{"t1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][2]int{{1, 6}, {8, 10}, {15, 18}}},
		{"t2", [][]int{{1, 2}, {2, 6}}, [][2]int{{1, 6}}},
		{"t3", [][]int{{1, 2}, {2, 3}, {3, 6}}, [][2]int{{1, 6}}},
		{"t4", [][]int{{1, 5}, {2, 3}, {3, 6}}, [][2]int{{1, 6}}},
		{"t5", [][]int{{1, 5}, {2, 13}, {23, 24}}, [][2]int{{1, 13}, {23, 24}}},
		{"t6", [][]int{{1, 5}}, [][2]int{{1, 5}}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			got := merge(v.input)
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
