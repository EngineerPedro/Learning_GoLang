package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	for r := 0; r < n; r++ {
		dupsC, dupsR := [9]bool{}, [9]bool{}
		for c := 0; c < n; c++ {
			if board[r][c] != '.' {
				if dupsC[board[r][c]-'0'-1] {
					return false
				}
				dupsC[board[r][c]-'0'-1] = true
			}
			if board[c][r] != '.' {
				if dupsR[board[c][r]-'0'-1] {
					return false
				}
				dupsR[board[c][r]-'0'-1] = true
			}
		}
	}
	for j := 0; j < n; j += 3 {
		for i := 0; i < n; i += 3 {
			dups := [9]bool{}
			for r := j; r < j+3; r++ {
				for c := i; c < i+3; c++ {
					if board[c][r] == '.' {
						continue
					}
					if dups[board[c][r]-'0'-1] {
						return false
					}
					dups[board[c][r]-'0'-1] = true
				}
			}
		}
	}
	return true
}
