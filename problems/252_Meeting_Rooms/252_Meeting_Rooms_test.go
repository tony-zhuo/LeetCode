package problems

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no overlap",
			args: args{
				intervals: [][]int{{1, 3}, {6, 10}, {11, 15}},
			},
			want: true,
		},
		{
			name: "overlap",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: true,
		},
		{
			name: "direct overlap",
			args: args{
				intervals: [][]int{{1, 5}, {3, 7}},
			},
			want: false,
		},
		{
			name: "same start time",
			args: args{
				intervals: [][]int{{1, 5}, {1, 3}},
			},
			want: false,
		},
		{
			name: "same end time",
			args: args{
				intervals: [][]int{{1, 5}, {3, 5}},
			},
			want: false,
		},
		{
			name: "empty intervals",
			args: args{
				intervals: [][]int{},
			},
			want: true,
		},
		{
			name: "single meeting",
			args: args{
				intervals: [][]int{{1, 3}},
			},
			want: true,
		},
		{
			name: "adjacent meetings",
			args: args{
				intervals: [][]int{{1, 3}, {3, 5}},
			},
			want: true,
		},
		{
			name: "multiple overlaps",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: false,
		},
		{
			name: "complex case no overlap",
			args: args{
				intervals: [][]int{{9, 10}, {4, 9}, {4, 17}},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				intervals: [][]int{{5, 8}, {9, 15}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
