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
	ended := make([]bool, len(lists))
	for {
		var minList *ListNode
		minimum := 2147483647
		for i, list := range lists {
			if list == nil {
				ended[i] = true
				continue
			}
			if list.Val < minimum {
				minimum = list.Val
				minList = list
			}
		}
		node.Next = &ListNode{minList.Val, nil}
		minList = minList.Next
		node = node.Next

		//check if all lists is ended
		allEnded := true
		for _, e := range ended {
			if e == false {
				allEnded = false
				break
			}
		}
		if allEnded {
			break
		}
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
