package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindPairsSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "base",
			nums:   []int{2, 4, 5, 3},
			target: 7,
			want:   [][]int{{2, 5}, {4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pairs := FindPairsSum(tt.nums, tt.target)
			require.Equal(t, tt.want, pairs)
		})
	}
}
