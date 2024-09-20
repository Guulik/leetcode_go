package LinkedList

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHasCycle(t *testing.T) {
	testsTable := []struct {
		name   string
		head   *ListNode
		pos    int
		result bool
	}{
		{
			name:   "dva",
			head:   BuildLinkedList([]int{3, 2, 0, -4}),
			pos:    1,
			result: true,
		},
	}
	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {

			if tt.pos > -1 {
				node := tt.head
				pos := 0
				var cyclePoint *ListNode
				for node != nil {
					if pos == tt.pos {
						cyclePoint = node
					}
					if node.Next == nil {
						node.Next = cyclePoint
					}
					node = node.Next
					pos++
					res := HasCycle(tt.head)
					if res == true {
						require.Equal(t, tt.result, res)
						break
					}
				}
			} else {
				res := HasCycle(tt.head)
				require.Equal(t, tt.result, res)
			}
		})
	}
}
