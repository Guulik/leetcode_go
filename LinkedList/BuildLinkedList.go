package LinkedList

// support
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
