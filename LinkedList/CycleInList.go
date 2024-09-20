package LinkedList

func HasCycle(head *ListNode) bool {
	pointed := make(map[*ListNode]bool)
	node := head

	for node != nil {
		if beenHere := pointed[node]; !beenHere {
			pointed[node] = true
		} else {
			return true
		}
		node = node.Next
	}
	return false
}
