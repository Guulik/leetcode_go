package LinkedList

func ListValues(head *ListNode) []int {
	node := head
	vals := make([]int, 0, 5)

	for node != nil {
		vals = append(vals, node.Val)
		node = node.Next
	}
	return vals
}
