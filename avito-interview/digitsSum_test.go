package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_digitsSum(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "base",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  []int{5, 7, 9},
		},
		{
			name:  "tischa",
			nums1: []int{5, 4, 4},
			nums2: []int{4, 5, 6},
			want:  []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			digits := digitsSum(tt.nums1, tt.nums2)
			require.Equal(t, tt.want, digits)
		})
	}
}
