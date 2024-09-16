package main

// блять я чета заебался
func SearchInRotated(nums []int, target int) int {
	lefty := 0
	righty := len(nums) - 1
	current := lefty + righty/2

	for lefty < righty {
		if target == nums[current] {
			return current
		}
		if target < nums[current] {
			if nums[lefty] < target {
				righty = current - 1
			}
			if nums[righty] > target {
				lefty = current + 1
			}
		}
		if target > nums[current] {
			if nums[righty] < target && nums[lefty] < target {
				righty = current - 1
			}
			if nums[righty] > target && nums[lefty] > target {
				lefty = current + 1
			}
		}
	}
	return 228
}
