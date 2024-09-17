package Stack

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	st := make(stack, 0, len(temperatures))

	for i, t := range temperatures {
		var index int
		for len(st) != 0 && t > temperatures[st.Top()] {
			st, index = st.Pop()
			result[index] = i - index
		}
		st = st.Push(i)
	}
	return result
}
