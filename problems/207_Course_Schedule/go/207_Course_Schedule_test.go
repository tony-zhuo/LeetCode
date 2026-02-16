package problems

import "testing"

func TestCanFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("CanFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
