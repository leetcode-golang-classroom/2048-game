package internal

// sideSize - 預設 sideSize
const sideSize = 4

// Game - 紀錄當下遊戲處理狀態
//
//	board [][]int - 紀錄盤面狀態
type Game struct {
	board [][]int
}

// Init - 初始化
func (g *Game) Init(data [][]int) {
	// 建立棋盤
	g.board = make([][]int, sideSize)
	for index := range g.board {
		g.board[index] = make([]int, sideSize)
	}
	// checkout input value
	if len(data) != sideSize || len(data[0]) != sideSize {
		return
	}
	// setup data
	for r := range sideSize {
		for c := range sideSize {
			if data[r][c] != 0 {
				g.board[r][c] = data[r][c]
			}
		}
	}
}

func NewGame() *Game {
	return &Game{}
}
