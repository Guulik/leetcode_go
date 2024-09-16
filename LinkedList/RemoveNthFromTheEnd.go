package LinkedList

func RemoveNthFromEnd(head *ListNode, pos int) *ListNode {
	if head.Next == nil {
		return nil
	}
	node := head
	sz := 0
	st := make(stack, 0, 5)
	for node != nil {
		sz++
		st = st.Push(node)
		node = node.Next
	}
	if pos == sz {
		head = head.Next
		return head
	}
	if pos == 1 {
		st, _ = st.Pop()
		st, node = st.Pop()
		node.Next = nil
		return head
	}
	for popped := 1; popped < pos; popped++ {
		st, _ = st.Pop()
	}
	st, victim := st.Pop()
	st, node = st.Pop()
	node.Next = victim.Next
	return head
}
