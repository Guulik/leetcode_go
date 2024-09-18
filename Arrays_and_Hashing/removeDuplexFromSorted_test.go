package Arrays_and_Hashing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		expectedNums []int
	}{
		{
			name:         "Один один два",
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
		},
		{
			name:         "пять",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:         "Одинки",
			nums:         []int{1, 1, 1},
			expectedNums: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := RemoveDuplicatesSorted(tt.nums)
			require.Equal(t, len(tt.expectedNums), k)
			for i := 0; i < k; i++ {
				require.Equal(t, tt.expectedNums[i], tt.nums[i])
			}
		})
	}
}
