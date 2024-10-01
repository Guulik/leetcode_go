package avito_interview

func RemoveZeroes(nums []int) []int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}

		if fast != slow {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}

		fast++
		slow++
	}
	return nums
}
