package Stack

import "strconv"

func Evaluate(tokens []string) int {
	st := make(stack, 0, 10)
	res := 0
	if len(tokens) == 1 {
		n, _ := strconv.Atoi(tokens[0])
		return n
	}
	for _, s := range tokens {
		t, err := strconv.Atoi(s)
		if err == nil {
			st = st.Push(t)
		} else {
			var a, b int
			switch s[0] {
			case '+':
				st, a = st.Pop()
				st, b = st.Pop()
				res = b + a
				st = st.Push(res)
			case '-':
				st, a = st.Pop()
				st, b = st.Pop()
				res = b - a
				st = st.Push(res)
			case '*':
				st, a = st.Pop()
				st, b = st.Pop()
				res = b * a
				st = st.Push(res)
			case '/':
				st, a = st.Pop()
				st, b = st.Pop()
				res = b / a
				st = st.Push(res)
			}
		}
	}
	return res
}
