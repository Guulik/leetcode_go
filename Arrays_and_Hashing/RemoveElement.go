package Arrays_and_Hashing

func RemoveElement(nums []int, val int) int {
	current := 0
	for scout := 0; scout < len(nums); scout++ {
		if nums[scout] != val {
			nums[current] = nums[scout]
			current++
		}
	}
	return current
}
