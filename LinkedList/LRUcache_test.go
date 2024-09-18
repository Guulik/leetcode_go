package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name              string
		capacity          int
		funcs             []string
		args              [][]int
		expectedGetOutput []int
	}{
		{
			name:              "база",
			capacity:          2,
			funcs:             []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args:              [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expectedGetOutput: []int{1, -1, -1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			getpos := 0
			for i, f := range tt.funcs {
				switch f {
				case "put":
					cache.Put(tt.args[i][0], tt.args[i][1])
				case "get":
					val := cache.Get(tt.args[i][0])
					require.Equal(t, tt.expectedGetOutput[getpos], val)
					getpos++
				}
			}
		})
	}

}
