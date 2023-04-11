package problems

import "testing"

func TestIsRobotBounded(t *testing.T) {
	test := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "test1",
			in:   "GGLLGG",
			want: true,
		},
		{
			name: "test2",
			in:   "GG",
			want: false,
		},
		{
			name: "test3",
			in:   "GL",
			want: true,
		},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := isRobotBounded(v.in)
			if got != v.want {
				t.Errorf("got: %v, want: %v", got, v.want)
			}
		})
	}
}
