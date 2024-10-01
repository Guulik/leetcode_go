package avito_interview

func RemoveZeroes(nums []int) []int {
	fast := 0
	slow := 0
	for slow < len(nums) {
		if nums[slow] == 0 {
			fast = slow
			for nums[fast] == 0 && fast < len(nums) {
				fast++
			}
			if nums[fast] != 0 {
				nums[slow] = nums[fast]
				slow++
			}
		} else {
			slow++
		}
	}
	return nums
}
