package avito_interview

import "slices"

func FindPairsSum(nums []int, target int) [][]int {
	pairs := make([][]int, 0, len(nums)/2)
	slices.Sort(nums)
	var (
		left  int
		right = len(nums) - 1
	)

	for left < right && left < len(nums)/2+1 {
		sum := nums[left] + nums[right]
		if sum == target {
			pairs = append(pairs, []int{nums[left], nums[right]})
			left++
			right--
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return pairs
}
