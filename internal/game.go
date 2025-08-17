package internal

// sideSize - 預設 sideSize
const sideSize = 4

// randomPositioner - 根據給訂的 TotalSize 隨機產生一個位置
type randomPositoner func(TotalSize int) int

// randomGenerator - 隨機給個 0 - 1 之間的機率數
type randomGenerator func() float64

// Game - 紀錄當下遊戲處理狀態
//
//	board [][]int - 紀錄盤面狀態
type Game struct {
	board               [][]int
	randomPositonerFunc randomPositoner
	randomFunc          randomGenerator
}

// Init - 初始化
func (g *Game) Init(data [][]int, randomPosFunc randomPositoner, randomFunc randomGenerator) {
	// setup random functions
	g.randomPositonerFunc = randomPosFunc
	g.randomFunc = randomFunc
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

// addRandomTile - 新增隨機的 2 或是 4 到一個空的 tile 內
func (g *Game) addRandomTile() {
	// 蒐集所有空的 tile
	emptyTiles := make([][2]int, 0, sideSize*sideSize)
	for r := 0; r < sideSize; r++ {
		for c := 0; c < sideSize; c++ {
			if g.board[r][c] == 0 {
				emptyTiles = append(emptyTiles, [2]int{r, c})
			}
		}
	}
	// 如果所有格子都滿了
	if len(emptyTiles) == 0 {
		return
	}
	// 選出要填入的位置
	position := emptyTiles[g.randomPositonerFunc(len(emptyTiles))]
	// 90% 機率是 2 , 10%  機率則為 4
	value := 2
	if g.randomFunc() < 0.1 {
		value = 4
	}
	g.board[position[0]][position[1]] = value
}

func NewGame() *Game {
	return &Game{
		nil,
		defaultRandomPositioner,
		defaultRandomFunc,
	}
}
