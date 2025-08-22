package game

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

type RandomType int

const (
	Default RandomType = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

// Init - 初始化
func (g *Game) Init(data [][]int, randomPosFunc randomPositoner, randomFunc randomGenerator) {
	// setup random functions
	if randomFunc != nil {
		g.randomPositonerFunc = randomPosFunc
	}
	if randomPosFunc != nil {
		g.randomFunc = randomFunc
	}
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

var collectStrategyMap map[RandomType]func(g *Game) [][2]int = map[RandomType]func(g *Game) [][2]int{
	Default: func(g *Game) [][2]int {
		emptyTiles := make([][2]int, 0, sideSize*sideSize)
		for r := 0; r < sideSize; r++ {
			for c := 0; c < sideSize; c++ {
				if g.board[r][c] == 0 {
					emptyTiles = append(emptyTiles, [2]int{r, c})
				}
			}
		}
		return emptyTiles
	},
	DirectionDown: func(g *Game) [][2]int {
		emptyTiles := make([][2]int, 0, sideSize*sideSize)
		for r := 0; r < sideSize; r++ {
			for c := 0; c < sideSize; c++ {
				if g.board[r][c] == 0 {
					if len(emptyTiles) >= 4 {
						break
					}
					emptyTiles = append(emptyTiles, [2]int{r, c})
				}
			}
		}
		return emptyTiles
	},
	DirectionUp: func(g *Game) [][2]int {
		emptyTiles := make([][2]int, 0, sideSize*sideSize)
		for r := sideSize - 1; r > 0; r-- {
			for c := 0; c < sideSize; c++ {
				if g.board[r][c] == 0 {
					if len(emptyTiles) >= 4 {
						break
					}
					emptyTiles = append(emptyTiles, [2]int{r, c})
				}
			}
		}
		return emptyTiles
	},
	DirectionLeft: func(g *Game) [][2]int {
		emptyTiles := make([][2]int, 0, sideSize*sideSize)
		for c := sideSize - 1; c > 0; c-- {
			for r := 0; r < sideSize; r++ {
				if g.board[r][c] == 0 {
					if len(emptyTiles) >= 4 {
						break
					}
					emptyTiles = append(emptyTiles, [2]int{r, c})
				}
			}
		}
		return emptyTiles
	},
	DirectionRight: func(g *Game) [][2]int {
		emptyTiles := make([][2]int, 0, sideSize*sideSize)
		for c := 0; c < sideSize; c++ {
			for r := 0; r < sideSize; r++ {
				if g.board[r][c] == 0 {
					if len(emptyTiles) >= 4 {
						break
					}
					emptyTiles = append(emptyTiles, [2]int{r, c})
				}
			}
		}
		return emptyTiles
	},
}

// AddRandomTile - 新增隨機的 2 或是 4 到一個空的 tile 內
func (g *Game) AddRandomTile(randomType RandomType) {
	// 蒐集所有空的 tile
	emptyTiles := collectStrategyMap[randomType](g)

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

// MoveLeft - 整個 board 同時左移
func (g *Game) MoveLeft() {
	for r := 0; r < sideSize; r++ {
		g.board[r] = g.slideAndMergeLeft(g.board[r][:])
	}
}

// transpose - 把整個  board 作轉置
func (g *Game) transpose() [][]int {
	// board[r][c] = board[c][r]
	transposedBoard := make([][]int, sideSize)
	for r := range sideSize {
		transposedBoard[r] = make([]int, sideSize)
		for c := range sideSize {
			transposedBoard[r][c] = g.board[c][r]
		}
	}
	return transposedBoard
}

// reverseRow - 把整個 Row 反過來
func (g *Game) reverseRow(row []int) []int {
	reversedRow := make([]int, len(row))
	for idx := range len(row) {
		reversedRow[idx] = row[len(row)-idx-1]
	}
	return reversedRow
}

// MoveRight - 整個 board 同時往右
func (g *Game) MoveRight() {
	// 先把整個  board 作 reverse
	for i := range g.board {
		g.board[i] = g.reverseRow(g.board[i])
	}
	// 把整個 board 往左移動
	g.MoveLeft()
	// 再整個  board 作 reverse 回來
	for i := range g.board {
		g.board[i] = g.reverseRow(g.board[i])
	}
}

// MoveUp - 整個 board 同時往上
func (g *Game) MoveUp() {
	// 先把 board 作轉置
	g.board = g.transpose()
	// 再把 board 同時往左
	g.MoveLeft()
	// 再把 board 作轉置
	g.board = g.transpose()
}

// MoveDown - 把整個 board 往下移動
func (g *Game) MoveDown() {
	// 先把整個 board 轉置
	g.board = g.transpose()
	// 再把整個 board 往右滑
	g.MoveRight()
	// 再把整個 board 轉置
	g.board = g.transpose()
}

func NewGame() *Game {
	board := make([][]int, sideSize)
	for idx := range board {
		board[idx] = make([]int, sideSize)
	}
	return &Game{
		board,
		defaultRandomPositioner,
		defaultRandomFunc,
	}
}

func (g *Game) Data(row int, col int) int {
	return g.board[row][col]
}
