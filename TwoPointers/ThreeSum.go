package TwoPointers

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {
	sums := [][]int{}
	slices.Sort(nums)
	for i, fixed := range nums {

		if i > 0 && fixed == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := fixed + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			}
			if threeSum < 0 {
				l++
			}
			if threeSum == 0 {
				sums = append(sums, []int{fixed, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return sums
}
