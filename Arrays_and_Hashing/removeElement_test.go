package Arrays_and_Hashing

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		val          int
		expectedNums []int
	}{
		{
			name:         "Один один два",
			nums:         []int{3, 2, 2, 3},
			val:          3,
			expectedNums: []int{2, 2},
		},
		{
			name:         "счёт",
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			expectedNums: []int{0, 1, 3, 0, 4},
		}, {
			name:         "два",
			nums:         []int{2},
			val:          3,
			expectedNums: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := RemoveElement(tt.nums, tt.val)
			require.Equal(t, len(tt.expectedNums), k)
			for i := 0; i < k; i++ {
				require.Equal(t, tt.expectedNums[i], tt.nums[i])
			}
		})
	}
}
