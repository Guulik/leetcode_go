package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name         string
		list         *ListNode
		pos          int
		expectedList *ListNode
	}{
		{
			name:         "четверку кик",
			list:         BuildLinkedList([]int{1, 2, 3, 4, 5}),
			pos:          2,
			expectedList: BuildLinkedList([]int{1, 2, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned := RemoveNthFromEnd(tt.list, tt.pos)
			require.Equal(t, ListValues(tt.expectedList), ListValues(cleaned))
		})
	}
}
