package problems

import "testing"

func TestSolveNQueens(t *testing.T) {
	test := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "test1",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "test2",
			n:    1,
			want: [][]string{
				{"Q"},
			},
		},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := solveNQueens(v.n)
			if len(got) != len(v.want) {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(v.want[i]) {
					t.Errorf("got: %v, want: %v", got, v.want)
				}
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != v.want[i][j] {
						t.Errorf("got: %v, want: %v", got, v.want)
					}
				}
			}
		})
	}

}
