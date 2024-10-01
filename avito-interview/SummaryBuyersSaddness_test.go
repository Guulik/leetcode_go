package avito_interview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSaddness(t *testing.T) {
	tests := []struct {
		name  string
		goods []int
		needs []int
		want  int
	}{
		{
			name:  "bas",
			goods: []int{8, 3, 5},
			needs: []int{5, 6},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sad := GetSaddness(tt.goods, tt.needs)
			require.Equal(t, tt.want, sad)
		})
	}
}
