package Arrays_and_Hashing

func ProductExceptSelf(nums []int) []int {
	//идея: собрать два вспомогательных массива: в первом записать все произведения слева от числа, во втором - справа
	leftProduct := make([]int, len(nums))
	leftProduct[0] = 1
	rightProduct := make([]int, len(nums))
	rightProduct[len(nums)-1] = 1
	productExcept := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		productExcept[i] = leftProduct[i] * rightProduct[i]
	}
	return productExcept
}
