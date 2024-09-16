package Arrays_and_Hashing

func ContainsDuplicate(nums []int) bool {
	uniqueNums := make(map[int]int)
	for _, n := range nums {
		if uniqueNums[n] == 0 {
			uniqueNums[n] = 1
		} else {
			return true
		}
	}
	return false
}
