package BinarySearch

func SearchInRotated(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2

	for left <= right {
		switch target {
		case nums[left]:
			return left
		case nums[right]:
			return right
		case nums[middle]:
			return middle
		}
		if nums[left] > target {
			if nums[middle] > target {
				right = middle - 1
				middle = (left + right) / 2
				continue
			}
			left = middle + 1
			middle = (left + right) / 2
			continue
		}
		if nums[right] < target {
			if nums[middle] < target {
				left = middle + 1
				middle = (left + right) / 2
				continue
			}
			right = middle - 1
			middle = (left + right) / 2
			continue
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
		middle = (left + right) / 2
	}
	return -1
}
