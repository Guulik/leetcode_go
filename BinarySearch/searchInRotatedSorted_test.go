package BinarySearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchInRotated(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		target        int
		expectedIndex int
	}{
		{
			name:          "baza",
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			target:        0,
			expectedIndex: 4,
		},
		{
			name:          "netu 3",
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			target:        3,
			expectedIndex: -1,
		},
		{
			name:          "odin edinitchka",
			nums:          []int{1},
			target:        1,
			expectedIndex: 0,
		},
		{
			name:          "netu nolika u edinichki((",
			nums:          []int{1},
			target:        0,
			expectedIndex: -1,
		},
		{
			name:          "vosem",
			nums:          []int{4, 5, 6, 7, 8, 1, 2, 3},
			target:        8,
			expectedIndex: 4,
		}, {
			name:          "183 test",
			nums:          []int{5, 1, 2, 3, 4},
			target:        1,
			expectedIndex: 1,
		},
		{
			name:          "176 test",
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			target:        0,
			expectedIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := SearchInRotated(tt.nums, tt.target)
			require.Equal(t, tt.expectedIndex, index)
		})
	}
}
