package Tries

func LexicalOrder(n int) []int {
	lexica := make([]int, 0, n)

	for series := 1; len(lexica) < n; series++ {
		numbre := series
		for numbre <= n && getFirstDigit(numbre) == series && len(lexica) < n {
			lexica = append(lexica, numbre)
			for numbre*10 <= n {
				numbre *= 10
				lexica = append(lexica, numbre)
			}
			numbre += 1
		}
	}
	return lexica
}

func getFirstDigit(n int) int {
	num := n
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	divider := 1
	for range count - 1 {
		divider *= 10
	}
	return n / divider
}
