package LinkedList

func ReorderList(head *ListNode) {
	node := head
	lenRest := 0
	st := make(stack, 0, 3)
	for node != nil {
		lenRest++
		st = st.Push(node)
		node = node.Next
	}
	for lenRest > 2 {
		lenRest--
		tail, preTail := &ListNode{}, &ListNode{}
		st, tail = st.Pop()
		st, preTail = st.Pop()
		st = st.Push(preTail)
		tail.Next = head.Next
		head.Next = tail
		preTail.Next = nil
		head = head.Next.Next
		lenRest--
	}
}
