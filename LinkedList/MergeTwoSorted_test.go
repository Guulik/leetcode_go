package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeSortedLists(t *testing.T) {
	tests := []struct {
		name       string
		list1      *ListNode
		list2      *ListNode
		mergedList *ListNode
	}{
		{
			name:       "первый",
			list1:      BuildLinkedList([]int{1, 2, 4}),
			list2:      BuildLinkedList([]int{1, 3, 4}),
			mergedList: BuildLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:       "pusto",
			list1:      BuildLinkedList([]int{}),
			list2:      BuildLinkedList([]int{}),
			mergedList: BuildLinkedList([]int{}),
		},
		{
			name:       "pusto i nolik",
			list1:      BuildLinkedList([]int{}),
			list2:      BuildLinkedList([]int{0}),
			mergedList: BuildLinkedList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeSortedLists(tt.list1, tt.list2)
			require.Equal(t, tt.mergedList, merged)
		})
	}
}
