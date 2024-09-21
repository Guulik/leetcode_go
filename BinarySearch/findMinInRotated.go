package BinarySearch

func FindMinRotatedSorted(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if left == right {
			return nums[right]
		}
		if nums[left] < nums[right] {
			return nums[left]
		}
		if nums[middle] < nums[left] && nums[middle] < nums[right] {
			right = middle
		}
		if nums[middle] > nums[right] {
			left = middle + 1
		}
	}
	return -1
}
