package Stack

func DailyTemperatures(temperatures []int) []int {
	daysToWarm := make([]int, len(temperatures))
	var nextDays = make(stack, 0, len(temperatures))
	for i := 0; i < len(temperatures)-1; i++ {
		nextDays = fillStackFromEnd(temperatures, nextDays, i)
		daysCount := 0
		dayTemp := -1
		lenNextDays := len(nextDays)
		for k := 0; k < lenNextDays; k++ {
			nextDays, dayTemp = nextDays.Pop()
			daysCount++
			if dayTemp > temperatures[i] {
				daysToWarm[i] = daysCount
				break
			}
		}
	}
	return daysToWarm
}

func fillStackFromEnd(source []int, stackToFill stack, startingIndex int) stack {
	lastIndex := len(source) - 1
	for i := lastIndex; i > startingIndex; i-- {
		stackToFill = stackToFill.Push(source[i])
	}
	return stackToFill
}
