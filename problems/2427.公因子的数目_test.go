package problems

import "testing"

func TestCommonFactors(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1}, 1},
		{"2", args{2, 2}, 2},
		{"3", args{2, 3}, 1},
		{"4", args{2, 4}, 2},
		{"5", args{2, 5}, 1},
		{"6", args{2, 6}, 2},
		{"7", args{2, 7}, 1},
		{"8", args{2, 8}, 2},
		{"9", args{2, 9}, 1},
		{"10", args{2, 10}, 2},
		{"11", args{2, 11}, 1},
		{"12", args{3, 12}, 2},
		{"13", args{2, 13}, 1},
		{"14", args{4, 14}, 2},
		{"15", args{5, 15}, 2},
		{"16", args{4, 16}, 3},
		{"17", args{2, 17}, 1},
		{"18", args{6, 18}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonFactors(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("commonFactors() = %v, want %v", got, tt.want)
			}
		})
	}

}
