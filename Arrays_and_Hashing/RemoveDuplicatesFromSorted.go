package Arrays_and_Hashing

func RemoveDuplicatesSorted(nums []int) int {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
