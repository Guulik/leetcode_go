package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeTwoSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "base",
			nums1: []int{1, 2, 5},
			nums2: []int{1, 2, 3, 4, 6},
			want:  []int{1, 1, 2, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeTwoSortedArrays(tt.nums1, tt.nums2, len(tt.nums1), len(tt.nums2))
			require.Equal(t, tt.want, merged)
		})
	}
}
