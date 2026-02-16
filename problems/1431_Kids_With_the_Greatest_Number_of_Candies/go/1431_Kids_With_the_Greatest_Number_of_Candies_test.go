package problems

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example 1",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
		{
			name: "Example 2",
			args: args{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			name: "Example 3",
			args: args{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
		{
			name: "Single kid",
			args: args{
				candies:      []int{5},
				extraCandies: 3,
			},
			want: []bool{true},
		},
		{
			name: "All kids can have the greatest",
			args: args{
				candies:      []int{1, 1, 1, 1},
				extraCandies: 0,
			},
			want: []bool{true, true, true, true},
		},
		{
			name: "Extra candies make everyone equal to max",
			args: args{
				candies:      []int{2, 8, 7},
				extraCandies: 6,
			},
			want: []bool{true, true, true},
		},
		{
			name: "Large numbers",
			args: args{
				candies:      []int{100, 200, 150, 175},
				extraCandies: 50,
			},
			want: []bool{false, true, true, true},
		},
		{
			name: "Zero extra candies",
			args: args{
				candies:      []int{5, 10, 3, 8},
				extraCandies: 0,
			},
			want: []bool{false, true, false, false},
		},
		{
			name: "Everyone can reach maximum with extra",
			args: args{
				candies:      []int{1, 2, 3},
				extraCandies: 2,
			},
			want: []bool{true, true, true},
		},
		{
			name: "Mixed case with zeros",
			args: args{
				candies:      []int{0, 5, 3, 2},
				extraCandies: 5,
			},
			want: []bool{true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
