package LinkedList

func MergeSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resultHead := &ListNode{}

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val >= list2.Val {
		resultHead = list2
		list2 = list2.Next
	} else {
		resultHead = list1
		list1 = list1.Next
	}
	current := resultHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return resultHead
}
