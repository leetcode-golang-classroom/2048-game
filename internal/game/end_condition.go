package game

// hasEmptyTile - 判斷是否還有可移動的位置
func (g *Game) hasEmptyTile() bool {
	for row := 0; row < sideSize; row++ {
		for col := 0; col < sideSize; col++ {
			if g.board[row][col] == 0 {
				return true
			}
		}
	}

	return false
}

// canMerge 判斷是否有可以合併的 tiles
func (g *Game) canMerge() bool {
	for row := 0; row < sideSize; row++ {
		for col := 0; col < sideSize; col++ {
			if col < sideSize-1 && g.board[row][col] == g.board[row][col+1] {
				return true
			}
			if row < sideSize-1 && g.board[row][col] == g.board[row+1][col] {
				return true
			}
		}
	}

	return false
}

// IsGameOver - 判斷遊戲是否無法繼續
func (g *Game) IsGameOver() bool {
	return !g.hasEmptyTile() && !g.canMerge()
}

// InitGame 初始化遊戲
func (g *Game) InitGame() {
	board := make([][]int, sideSize)
	for idx := range board {
		board[idx] = make([]int, sideSize)
	}
	g.board = board
	g.AddRandomTile(Default)
	g.AddRandomTile(Default)
}
