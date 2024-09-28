package TwoPointers

import "slices"

// идея в том, чтобы прийти к TwoSum Arrays Sorted. А там уже легко решается двумя поинтерами
// крч сначала сортируем массив, потом в большом цикле прогоняем X и для него ищем twoSum = -X
// правда я не оч помню и представляю, как не допустить дубликатов для триплетов. это я подсмотрю

func ThreeSum(nums []int) [][]int {
	sums := make([][]int, 0, 3)
	slices.Sort(nums)
	for i := range len(nums) - 1 {
		if nums[i] > 0 {
			return sums
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum > 0 {
				right--
			}
			if sum < 0 {
				left++
			}
			if sum == 0 {
				sums = append(sums, []int{nums[i], nums[left], nums[right]})
				left++
				if nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return sums
}
