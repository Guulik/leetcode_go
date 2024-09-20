package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeKSortedLists(t *testing.T) {
	tests := []struct {
		name   string
		lists  []*ListNode
		result *ListNode
	}{
		{
			name:   "basa",
			lists:  buildMultipleLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}),
			result: BuildLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name:   "dead",
			lists:  buildMultipleLists([][]int{}),
			result: BuildLinkedList([]int{}),
		},
		{
			name:   "dead INSIDE",
			lists:  buildMultipleLists([][]int{{}}),
			result: BuildLinkedList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := MergeKSortedLists(tt.lists)
			assert.Equal(t, ListValues(tt.result), ListValues(merged))
		})
	}
}

func buildMultipleLists(listVals [][]int) []*ListNode {
	lists := make([]*ListNode, 0, len(listVals))
	for _, vals := range listVals {
		list := BuildLinkedList(vals)
		lists = append(lists, list)
	}
	return lists
}
