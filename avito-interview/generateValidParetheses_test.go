package avito_interview

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGenerateValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "tri",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		}, {
			name: "adin",
			n:    1,
			want: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentas := GenerateValidParentheses(tt.n)
			fmt.Println(parentas)
			for i := range len(parentas) {
				require.Equal(t, true, strings.EqualFold(tt.want[i], parentas[i]))
			}
		})
	}
}
