package BinarySearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMinRotatedSorted(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		min  int
	}{
		{
			name: "baza",
			nums: []int{3, 4, 5, 1, 2},
			min:  1,
		},
		{
			name: "nulik",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			min:  0,
		},
		{
			name: "decimal",
			nums: []int{11, 13, 15, 17},
			min:  11,
		},
		{
			name: "tri raz dva",
			nums: []int{3, 1, 2},
			min:  1,
		},
		{
			name: "sam pridumal",
			nums: []int{7, 0, 1, 2, 3},
			min:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minimim := FindMinRotatedSorted(tt.nums)
			require.Equal(t, tt.min, minimim)
		})
	}
}
