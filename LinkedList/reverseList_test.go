package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		original *ListNode
		reversed *ListNode
	}{
		{
			name:     "раздватричетырепять",
			original: BuildLinkedList([]int{1, 2, 3, 4, 5}),
			reversed: BuildLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "raz dva",
			original: BuildLinkedList([]int{1, 2}),
			reversed: BuildLinkedList([]int{2, 1}),
		},
		{
			name:     "empty",
			original: BuildLinkedList([]int{}),
			reversed: BuildLinkedList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversedHead := ReverseLinkedList(tt.original)
			vals := ListValues(reversedHead)
			require.Equal(t, ListValues(tt.reversed), vals)
		})
	}
}
