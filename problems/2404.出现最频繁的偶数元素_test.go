package problems

import "testing"

func TestMostFrequentEven(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 2, 3, 3, 3, 4}}, 2},
		{"test2", args{[]int{1, 2, 2, 3, 3, 3, 4, 4}}, 2},
		{"test3", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4}}, 4},
		{"test4", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}}, 4},
		{"test5", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4}}, 4},
		{"test6", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4}}, 4},
		{"test7", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4}}, 4},
		{"test8", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4}}, 4},
		{"test9", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}}, 4},
		{"test10", args{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}}, 4},
		{"test11", args{[]int{1, 2, 2, 3, 3, 3}}, 2},
		{"test12", args{[]int{1, 2, 2, 3, 3, 3, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentEven(tt.args.nums); got != tt.want {
				t.Errorf("MostFrequentEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
