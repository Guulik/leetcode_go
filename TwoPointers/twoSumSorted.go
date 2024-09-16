package TwoPointers

func TwoSumSorted(nums []int, target int) []int {
	// идея достаточно проста - использхуем два поинтера
	// один вначале - другой в конце. Если получившаяся сумма больше нужного, то двигаем правый влево, если наоборот - левый вправо
	// делаем пока поинтеры не столкнутся
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
		if sum == target {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}
