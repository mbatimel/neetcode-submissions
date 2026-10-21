func isValidSudoku(board [][]byte) bool {

    for i := 0; i < 9; i++ {
        rowSeen := make(map[byte]bool)
        colSeen := make(map[byte]bool)
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                if rowSeen[board[i][j]] {
                    return false
                }
                rowSeen[board[i][j]] = true
            }
            if board[j][i] != '.' {
                if colSeen[board[j][i]] {
                    return false
                }
                colSeen[board[j][i]] = true
            }
        }
    }


    for blockRow := 0; blockRow < 3; blockRow++ {
        for blockCol := 0; blockCol < 3; blockCol++ {
            seen := make(map[byte]bool)
            for i := 0; i < 3; i++ {
                for j := 0; j < 3; j++ {
                    r := blockRow*3 + i
                    c := blockCol*3 + j
                    if board[r][c] != '.' {
                        if seen[board[r][c]] {
                            return false
                        }
                        seen[board[r][c]] = true
                    }
                }
            }
        }
    }

    return true
}