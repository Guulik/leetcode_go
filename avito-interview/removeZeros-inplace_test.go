package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveZeroes(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "base",
			nums: []int{7, 3, 0, 0, 0, 2, 4, 0, 5, 19},
			want: []int{7, 3, 2, 4, 5, 19, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned := RemoveZeroes(tt.nums)
			require.Equal(t, tt.want, cleaned)
		})
	}
}
