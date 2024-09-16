package LinkedList

type stack []*ListNode

func (s stack) Push(value *ListNode) stack {
	return append(s, value)
}
func (s stack) Pop() (stack, *ListNode) {
	l := len(s)
	if l == 0 {
		return s, nil
	}
	return s[:l-1], s[l-1]
}
