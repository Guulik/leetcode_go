package LinkedList

func MergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if isAllEmpty(lists) {
		return nil
	}

	var head = &ListNode{Val: 0, Next: nil}
	node := head
	for {
		var minIndex = -1
		minimum := 2147483647
		for i, list := range lists {
			if list != nil && list.Val < minimum {
				minimum = list.Val
				minIndex = i
			}
		}
		// minIndex remain -1 only if no one list is not nil
		if minIndex == -1 {
			break
		}

		node.Next = lists[minIndex]
		lists[minIndex] = lists[minIndex].Next
		node = node.Next
	}
	return head.Next
}

func isAllEmpty(lists []*ListNode) bool {
	count := 0
	for _, l := range lists {
		if l == nil {
			count++
		}
	}
	return count == len(lists)
}
