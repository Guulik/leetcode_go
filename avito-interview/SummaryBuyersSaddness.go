package avito_interview

import (
	"math"
	"slices"
)

func GetSaddness(goods []int, needs []int) int {
	slices.Sort(goods)
	sad := 0

	for _, n := range needs {
		closest := binSearch(goods, n)
		for closest == -1 {
			closest = binSearch(goods, n+1)
			if closest > -1 {
				break
			}
			closest = binSearch(goods, n-1)
		}
		sad += int(math.Abs(float64(n - closest)))
	}
	return sad
}

func binSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return nums[middle]
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
	}
	return -1
}
