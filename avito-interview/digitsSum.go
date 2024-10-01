package avito_interview

func digitsSum(nums1, nums2 []int) []int {
	first, sec := len(nums1)-1, len(nums2)-1
	res := make([]int, max(len(nums1), len(nums2))+1)

	for first >= 0 && sec >= 0 {
		res[first+1] += nums1[first] + nums2[sec]
		if res[first+1] >= 10 {
			res[first+1] = res[first+1] % 10
			res[first] += 1
		}
		first--
		sec--
	}
	if first == 0 {
		for ; sec >= 0; sec-- {
			res[sec+1] += nums2[sec]
			if res[sec+1] >= 10 {
				res[sec+1] = res[sec+1] % 10
				res[sec] += 1
			}
		}
	}
	if sec == 0 {
		for ; first >= 0; first-- {
			res[first+1] += nums2[first]
			if res[first+1] >= 10 {
				res[first+1] = res[first+1] % 10
				res[first] += 1
			}
		}
	}
	if res[0] == 0 {
		res = append(res[:0], res[1:]...)
	}
	return res
}
