package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindDuplicatesConstantSpace(t *testing.T) {
	testsTable := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "dva",
			nums:   []int{1, 3, 4, 2, 2},
			result: 2,
		},
		{
			name:   "tri odin",
			nums:   []int{3, 1, 3, 4, 2},
			result: 3,
		},
		{
			name:   "tri mnoga",
			nums:   []int{3, 3, 3, 3, 3},
			result: 3,
		},
		{
			name:   "odin dvazhdi",
			nums:   []int{1, 1},
			result: 1,
		},
		{
			name:   "odin dva odin",
			nums:   []int{2, 1, 2},
			result: 2,
		},
		{
			name:   "один по краям",
			nums:   []int{1, 3, 4, 2, 1},
			result: 1,
		},
	}
	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			res := FindDuplicatesConstantSpace(tt.nums)
			require.Equal(t, tt.result, res)
		})
	}
}
