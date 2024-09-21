package Tries

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		order []int
	}{
		{
			name:  "pyatnadtsat",
			n:     13,
			order: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name:  "davdva",
			n:     22,
			order: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexica := LexicalOrder(tt.n)
			require.Equal(t, tt.order, lexica)
		})
	}
}
