package Arrays_and_Hashing

func GroupAnagram(words []string) [][]string {
	resMap := make(map[[26]int][]string)

	for _, word := range words {
		count := make([]int, 26, 26)
		for _, char := range word {
			count[char-'a'] += 1
		}
		fixedArrayCount := [26]int{}
		copy(fixedArrayCount[:], count)
		resMap[fixedArrayCount] = append(resMap[fixedArrayCount], word)
	}
	return mapValues(resMap)
}

func mapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
