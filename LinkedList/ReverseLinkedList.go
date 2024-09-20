package LinkedList

func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
