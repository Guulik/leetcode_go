package TwoPointers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		triplets [][]int
	}{
		{
			name:     "baza",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			triplets: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "nuli",
			nums:     []int{0, 0, 0, 0},
			triplets: [][]int{{0, 0, 0}},
		}, {
			name:     "-101",
			nums:     []int{-1, 0, 1},
			triplets: [][]int{{-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sums := ThreeSum(tt.nums)
			require.Equal(t, tt.triplets, sums)
		})
	}
}
