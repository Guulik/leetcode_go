package LinkedList

func MergeSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var resultHead *ListNode
	//нам нужно какое-то начальное значение для головешки.
	if list1.Val <= list2.Val {
		resultHead = list1
		//если первое значение взяли, значит повторно его использовать при переборе нельзя. Сдвигаем голову на 1 вперед
		list1 = list1.Next
	} else {
		resultHead = list2
		//если первое значение взяли, значит повторно его использовать при переборе нельзя. Сдвигаем голову на 1 вперед
		list2 = list2.Next
	}
	resNode := resultHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			resNode.Next = list1
			//двигаемся вперед по 1 листу
			list1 = list1.Next
		} else {
			resNode.Next = list2
			//двигаемся вперед по 2 листу
			list2 = list2.Next
		}
		resNode = resNode.Next
	}

	// put the remaining list
	if list1 != nil {
		resNode.Next = list1
	}
	if list2 != nil {
		resNode.Next = list2
	}
	return resultHead
}
