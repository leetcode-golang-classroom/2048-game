package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
)

// Update - 用來處理畫面偵測,與使用者互動，並且觸發狀態變更
func (g *GameLayout) Update() error {
	// 判斷是否遊戲結束
	if g.isGameOver {
		// 處理 restart 邏輯
		g.handleRestartGame()
		return nil
	}

	// 判斷是否 Player Win
	if g.isPlayerWin && !g.isContinue {
		g.handleContinueGame()
		return nil
	}
	// 根據輸入產生對應的更新
	g.handleInput()

	// 根據目前的盤面決定是否要顯示 You Win
	if !g.isContinue && g.gameInstance.IsPlayerWin() {
		g.isPlayerWin = true
		return nil
	}

	// 根據目前的盤面跟更新是否能夠繼續執行
	if g.gameInstance.IsGameOver() {
		g.isGameOver = true
	}
	return nil
}

// handleInput - 處理輸入產生對應的更新
func (g *GameLayout) handleInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.gameInstance.MoveUp()
		g.gameInstance.AddRandomTile(game.DirectionUp)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.gameInstance.MoveDown()
		g.gameInstance.AddRandomTile(game.DirectionDown)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.gameInstance.MoveLeft()
		g.gameInstance.AddRandomTile(game.DirectionLeft)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.gameInstance.MoveRight()
		g.gameInstance.AddRandomTile(game.DirectionRight)
	}
}

// handleRestartGame - 偵測目前 restart button
func (g *GameLayout) handleRestartGame() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if restartButtonRect.Min.X <= x && x <= restartButtonRect.Max.X &&
			restartButtonRect.Min.Y <= y && y <= restartButtonRect.Max.Y {
			g.restartGame()
		}
	}
}

// handleContinueGame - 偵測目前 restart button
func (g *GameLayout) handleContinueGame() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if restartButtonRect.Min.X <= x && x <= restartButtonRect.Max.X &&
			restartButtonRect.Min.Y <= y && y <= restartButtonRect.Max.Y {
			g.continueGame()
		}
	}
}

// restartGame - 重設目前遊戲狀態
func (g *GameLayout) restartGame() {
	g.gameInstance.InitGame()
	g.isGameOver = false
	g.isContinue = false
}

// continueGame -
func (g *GameLayout) continueGame() {
	g.isPlayerWin = false
	g.isContinue = true
}
