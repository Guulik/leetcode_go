package LinkedList

func FindDuplicatesConstantSpace(nums []int) int {
	var head = &ListNode{}
	head.Val = 0
	visited := make(map[int]bool)

	var node = head
	_, ok := visited[node.Val]
	for !ok {
		visited[node.Val] = true
		nextNode := &ListNode{Val: nums[node.Val]}
		node.Next = nextNode
		node = node.Next
		_, ok = visited[node.Val]
	}
	return node.Val
}
