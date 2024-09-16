package LinkedList

func ReverseLinkedList(head *ListNode) *ListNode {
	if head.Next != nil {
		head = ReverseLinkedList(head.Next)
	} else {
		return head
	}
	return head
}
