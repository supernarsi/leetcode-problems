package problems

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	test := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "test1",
			s:    "42",
			want: 42,
		},
		{
			name: "test2",
			s:    "   -42",
			want: -42,
		},
		{
			name: "test3",
			s:    "4193 with words",
			want: 4193,
		},
		{
			name: "test4",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "test5",
			s:    "-91283472332",
			want: -2147483648,
		},
		{
			name: "test6",
			s:    "2147483648",
			want: 2147483647,
		},
		{
			name: "test7",
			s:    "2147483646",
			want: 2147483646,
		},
		{
			name: "test8",
			s:    "2147483647",
			want: 2147483647,
		},
		{
			name: "test9",
			s:    "-2147483648",
			want: -2147483648,
		},
		{
			name: "test10",
			s:    "-2147483649",
			want: -2147483648,
		},
		{
			name: "test11",
			s:    "9223372036854775808",
			want: 2147483647,
		},
		{
			name: "test12",
			s:    "9223372036854775807",
			want: 2147483647,
		},
		{
			name: "test13",
			s:    "-9223372036854775808",
			want: -2147483648,
		},
		{
			name: "test14",
			s:    "-9223372036854775809",
			want: -2147483648,
		},
		{
			name: "test15",
			s:    "9223372036854775809",
			want: 2147483647,
		},
		{
			name: "test16",
			s:    "9223372036854775806",
			want: 2147483647,
		},
		{
			name: "test17",
			s:    "-9223372036854775807",
			want: -2147483648,
		},
		{
			name: "test18",
			s:    "-9223372036854775806",
			want: -2147483648,
		},
		{
			name: "test19",
			s:    "9223372036854775806",
			want: 2147483647,
		},
		{
			name: "test20",
			s:    "9223372036854775805",
			want: 2147483647,
		},
		{
			name: "test21",
			s:    "9223372036854775804",
			want: 2147483647,
		},
		{
			name: "test22",
			s:    "9223372036854775803",
			want: 2147483647,
		},
		{
			name: "test23",
			s:    "9223372036854775802",
			want: 2147483647,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			if got != tt.want {
				t.Errorf("MyAtoi(%s) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
