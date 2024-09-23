package SlidingWindow

func FindLongestSbstrWoutRep(str string) int {
	if len(str) == 0 {
		return 0
	}
	maxLength := 0
	length := 0
	slow := 0
	fast := 0
	letters := make(map[uint8]bool)
	for slow < len(str) && fast < len(str) {
		//если буквы нет ищо
		if _, ok := letters[str[fast]]; !ok {
			letters[str[fast]] = true
			length = fast - slow + 1
			maxLength = max(length, maxLength)
			fast++
		} else {
			letters = make(map[uint8]bool)
			length = 0
			slow++
			fast = slow
		}
	}
	return maxLength
}
