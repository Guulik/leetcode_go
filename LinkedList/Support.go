package LinkedList

func BuildLinkedList(vals []int) *ListNode {
	head := &ListNode{}
	node := head
	if len(vals) == 0 {
		return nil
	}
	for i, val := range vals {
		node.Val = val
		if i < len(vals)-1 {
			newNode := &ListNode{}
			node.Next = newNode
			node = node.Next
		}
	}
	return head
}

func ListValues(head *ListNode) []int {
	node := head
	vals := make([]int, 0, 5)

	for node != nil {
		vals = append(vals, node.Val)
		node = node.Next
	}
	return vals
}
