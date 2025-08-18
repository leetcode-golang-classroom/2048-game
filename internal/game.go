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

// slideAndMergeLeft - 向左滑動並且合併
func (g *Game) slideAndMergeLeft(row []int) []int {
	// 1 去掉空格
	filtered := make([]int, 0, len(row))
	for _, v := range row {
		if v != 0 {
			filtered = append(filtered, v)
		}
	}

	// 假設沒有空格
	if len(filtered) == 0 {
		return row
	}

	// 2 合併相鄰相同數字
	for i := 0; i < len(filtered)-1; i++ {
		if filtered[i] == filtered[i+1] {
			filtered[i] *= 2
			filtered[i+1] = 0
			i++ // 跳過剛合併的數字
		}
	}

	// 3 再次去掉空格
	result := make([]int, 0, len(row))
	for _, v := range filtered {
		if v != 0 {
			result = append(result, v)
		}
	}

	// 4 補充剩下的空格為 0
	for len(result) < len(row) {
		result = append(result, 0)
	}

	return result
}

// moveLeft - 整個 board 同時左移
func (g *Game) moveLeft() {
	for r := 0; r < sideSize; r++ {
		g.board[r] = g.slideAndMergeLeft(g.board[r][:])
	}
}

func NewGame() *Game {
	return &Game{
		nil,
		defaultRandomPositioner,
		defaultRandomFunc,
	}
}
