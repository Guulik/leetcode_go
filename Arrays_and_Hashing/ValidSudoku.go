package Arrays_and_Hashing

func IsValidSudoku(board [][]byte) bool {
	uniq := make(map[byte]bool)
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			p := board[row][column]
			if p != '.' {
				if uniq[p] == true {
					return false
				}
				uniq[p] = true
			}
		}
		uniq = make(map[byte]bool)
	}
	uniq = make(map[byte]bool)
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			p := board[row][column]
			if p != '.' {
				if uniq[p] == true {
					return false
				}
				uniq[p] = true
			}
		}
		uniq = make(map[byte]bool)
	}
	for cubeRow := 0; cubeRow < 3; cubeRow++ {
		for cubeColumn := 0; cubeColumn < 3; cubeColumn++ {
			for row := 0; row < 3; row++ {
				for column := 0; column < 3; column++ {
					p := board[row+3*cubeRow][column+3*cubeColumn]
					if p != '.' {
						if uniq[p] == true {
							return false
						}
						uniq[p] = true
					}
				}
			}
			uniq = make(map[byte]bool)
		}
	}
	return true
}
