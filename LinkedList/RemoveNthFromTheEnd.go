package LinkedList

func RemoveNthFromEnd(head *ListNode, pos int) *ListNode {
	node := head
	posFromStart := 0
	for node != nil {
		node = node.Next
		posFromStart++
	}
	posFromStart -= pos
	if posFromStart == 0 {
		head = head.Next
		return head
	}
	node = head
	currPos := 0
	for node != nil {
		if currPos == posFromStart-1 {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
		currPos++
	}
	return head
}
