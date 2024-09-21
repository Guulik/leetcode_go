package BinarySearch

func SearchInRotated(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2

	for left <= right {
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if nums[middle] < target || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		middle = (left + right) / 2
	}
	return -1
}
