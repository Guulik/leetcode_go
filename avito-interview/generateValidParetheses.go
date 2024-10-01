package avito_interview

import "strings"

func GenerateValidParentheses(n int) []string {
	st := make(stack, 0, n*2)
	res := make([]string, 0, n)

	var backtrack func(int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			res = append(res, strings.Join(st, ""))
			return
		}
		if openN < n {
			st = st.Push("(")
			backtrack(openN+1, closedN)
			st = st.Pop()
		}
		if closedN < openN {
			st = st.Push(")")
			backtrack(openN, closedN+1)
			st = st.Pop()
		}
	}
	backtrack(0, 0)
	return res
}

type stack []string

func (s stack) Push(val string) stack {
	return append(s, val)
}

func (s stack) Pop() stack {
	l := len(s)
	if l == 0 {
		return s
	}
	return s[:l-1]
}
