package problems

import "testing"

var cases = []struct {
	name string
	grid [][]byte
	want int
}{
	{
		name: "example 1 - one island",
		grid: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		want: 1,
	},
	{
		name: "example 2 - three islands",
		grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		want: 3,
	},
	{
		name: "all water",
		grid: [][]byte{
			{'0', '0', '0'},
			{'0', '0', '0'},
		},
		want: 0,
	},
	{
		name: "all land",
		grid: [][]byte{
			{'1', '1'},
			{'1', '1'},
		},
		want: 1,
	},
	{
		name: "single cell island",
		grid: [][]byte{
			{'1'},
		},
		want: 1,
	},
	{
		name: "diagonal islands not connected",
		grid: [][]byte{
			{'1', '0'},
			{'0', '1'},
		},
		want: 2,
	},
}

func Test_numIslands(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// deep copy grid since DFS mutates it
			grid := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]byte, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			got := numIslands(grid)
			if got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
