package Tries

func LexicalOrder(n int) []int {
	lexica := make([]int, 0, n)
	current := 1
	for range n {
		lexica = append(lexica, current)
		if current*10 <= n {
			current *= 10
		} else {
			for current%10 == 9 || current >= n {
				current /= 10
			}
			current += 1
		}
	}
	return lexica
}
