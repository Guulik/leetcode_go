package Stack

func ValidParentheses(s string) bool {
	st := make(Stack, 0)
	for _, parenta := range s {
		if parenta == '(' || parenta == '[' || parenta == '{' {
			st = st.Push(parenta)
		}

		if parenta == ')' || parenta == ']' || parenta == '}' {
			var opener int32
			st, opener = st.Pop()
			if parenta == ')' && opener != '(' {
				return false
			}
			if parenta == ']' && opener != '[' {
				return false
			}
			if parenta == '}' && opener != '{' {
				return false
			}
		}
	}
	if len(st) == 0 {
		return true
	} else {
		return false
	}
}

type Stack []int32

func (s Stack) Push(value int32) Stack {
	return append(s, value)
}
func (s Stack) Pop() (Stack, int32) {
	l := len(s)
	if l == 0 {
		return s, -1
	}
	return s[:l-1], s[l-1]
}
