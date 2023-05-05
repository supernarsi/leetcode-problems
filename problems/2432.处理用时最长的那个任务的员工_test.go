package problems

import "testing"

func TestHardestWorker(t *testing.T) {
	type args struct {
		n    int
		logs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{10, [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}}}, 1},
		{"test2", args{10, [][]int{{1, 1}, {2, 6}, {3, 7}, {4, 8}, {5, 9}, {6, 10}}}, 2},
		{"test3", args{10, [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}, {6, 10}, {7, 12}}}, 1},
		{"test4", args{10, [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}, {6, 10}, {7, 11}, {8, 23}}}, 8},
		{"test5", args{10, [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}, {6, 10}, {7, 11}, {8, 12}, {1, 23}}}, 1},
		{"test6", args{10, [][]int{{1, 5}, {2, 6}, {3, 17}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hardestWorker(tt.args.n, tt.args.logs); got != tt.want {
				t.Errorf("hardestWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}
