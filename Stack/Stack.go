package Stack

type stack []int

func (s stack) Push(value int) stack {
	return append(s, value)
}
func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}
