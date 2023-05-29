package problems

import "testing"

func TestAverageValue(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "test1", input: []int{1, 3, 6, 10, 12, 15}, want: 9},
		{name: "test2", input: []int{1, 2, 6, 10, 12, 18}, want: 12},
		{name: "test3", input: []int{1, 2, 4, 7, 10}, want: 0},
		{name: "test4", input: []int{1, 2, 4, 6, 10}, want: 6},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := averageValue(v.input); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
