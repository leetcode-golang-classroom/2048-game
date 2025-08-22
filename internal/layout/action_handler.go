package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
)

// Update - 用來處理畫面偵測,與使用者互動，並且觸發狀態變更
func (g *GameLayout) Update() error {
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

	return nil
}
