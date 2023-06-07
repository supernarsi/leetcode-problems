package problems

import "testing"

func TestMiceAndCheese(t *testing.T) {
	tests := []struct {
		name   string
		inputA []int
		inputB []int
		inputK int
		want   int
	}{
		{"t1", []int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2, 15},
		{"t2", []int{1, 1}, []int{1, 1}, 2, 2},
		{"t3", []int{1, 3, 2, 1}, []int{1, 1, 9, 9}, 3, 15},
		{"t4", []int{1, 3, 2, 1}, []int{1, 1, 9, 9}, 1, 22},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := miceAndCheese(v.inputA, v.inputB, v.inputK); got != v.want {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
