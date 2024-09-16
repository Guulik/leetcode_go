package Arrays_and_Hashing

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complementary := target - num
		if j, ok := m[num]; ok {
			return []int{i, j}
		}
		m[complementary] = i
	}
	return nil
}
