package Arrays_and_Hashing

func LongestSequence(nums []int) int {
	set := make(map[int]bool)
	maxSeq := 0
	curSeq := 0
	for _, n := range nums {
		set[n] = false //фалзе, чтобы обозначить, что мы ещё не спрашивали это число, пока перебираем лист
	}
	for _, n := range nums {
		if _, isL := set[n-1]; !isL {
			curSeq = 1
			k := n
			for _, isR := set[k+1]; isR; _, isR = set[k+1] {
				curSeq += 1
				k++
			}
		}
		if curSeq > maxSeq {
			maxSeq = curSeq
		}
	}
	return maxSeq
}
