package LinkedList

func HasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	if head == nil {
		return false
	}
	node := head
	for node.Next != nil {
		if _, ok := visited[node]; !ok {
			visited[node] = true
		} else {
			return true
		}
		node = node.Next
	}
	return false
}
