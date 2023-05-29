package problems

import "testing"

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		inputArr []int
		inputK   int
		inputX   int
		want     []int
	}{
		{"test1", []int{1, 2, 3, 4}, 4, -1, []int{1, 2, 3, 4}},
		{"test2", []int{11, 22, 33, 44, 55}, 3, 23, []int{11, 22, 33}},
		{"test3", []int{11, 22, 33, 44, 55}, 2, 50, []int{44, 55}},
		{"test4", []int{11, 22, 33, 44, 55}, 2, 44, []int{33, 44}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			gotArr := findClosestElements(v.inputArr, v.inputK, v.inputX)
			if len(gotArr) != len(v.want) {
				t.Errorf("%s got %v, want %v", v.name, gotArr, v.want)
			} else {
				for i, eachVal := range gotArr {
					if eachVal != v.want[i] {
						t.Errorf("%s got %v, want %v", v.name, gotArr, v.want)
					}
				}
			}
		})
	}
}
