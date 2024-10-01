package avito_interview

func MergeTwoSortedArrays(nums1 []int, nums2 []int, m int, n int) []int {
	res := make([]int, 0, m+n)
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	first, sec := 0, 0
	for first < m && sec < n {
		if nums1[first] < nums2[sec] {
			res = append(res, nums1[first])
			first++
		} else {
			res = append(res, nums2[sec])
			sec++
		}
	}
	if first == m-1 {

		res = append(res, nums2...)
	}
	if sec == n-1 {
		res = append(res, nums1...)
	}
	return res
}
