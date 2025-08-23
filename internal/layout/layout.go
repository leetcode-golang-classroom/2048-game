package layout

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
)

var restartButtonRect = image.Rect(165, 250, 285, 300) // X1,Y1,X2,Y2

const (
	tileSize  = 100
	gridSize  = 4
	padding   = 10
	WinWidth  = tileSize*gridSize + padding*(gridSize+1)
	WinHeight = tileSize*gridSize + padding*(gridSize+1)
)

type GameLayout struct {
	gameInstance *game.Game
	isGameOver   bool
}

// drawCell - 透過目前值來畫出目前 cell 的格子顏色
func (g *GameLayout) drawCell(screen *ebiten.Image, value, row, col int) {
	cellXPos := padding + col*(tileSize+padding)
	cellYPos := padding + row*(tileSize+padding)
	cellColor := getTileColor(value)
	cell := ebiten.NewImage(tileSize, tileSize)
	cell.Fill(cellColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cellXPos), float64(cellYPos))
	op.ColorScale.ScaleWithColor(g.tileBackgroundColor(value))
	screen.DrawImage(cell, op)
}

// drawTileText -  透過目前值來畫出目前 cell 的文字顏色
func (g *GameLayout) drawTileText(screen *ebiten.Image, value, row, col int) {
	if value > 0 {
		// 繪製數字 (置中)
		textValue := fmt.Sprintf("%d", value)
		textXPos := (col+1)*padding + col*tileSize + (tileSize)/2
		textYPos := (row+1)*padding + row*tileSize + (tileSize)/2
		textOpts := &text.DrawOptions{}
		textOpts.ColorScale.ScaleWithColor(getTileColor(value))
		textOpts.PrimaryAlign = text.AlignCenter
		textOpts.SecondaryAlign = text.AlignCenter
		textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
		text.Draw(screen, textValue, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   getFontSize(value),
		}, textOpts)
	}
}

func (g *GameLayout) Draw(screen *ebiten.Image) {
	// 背景色
	screen.Fill(color.RGBA{250, 248, 239, 255})
	// 畫出目前局面
	g.drawBoard(screen)
	// 當 gameOver 顯示 GameOver
	if g.isGameOver {
		g.drawGameOver(screen)
	}
}

// drawBoard - 畫出目前局面
func (g *GameLayout) drawBoard(screen *ebiten.Image) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			// 取出值
			value := g.gameInstance.Data(row, col)
			// 畫格子背景
			g.drawCell(screen, value, row, col)
			// 畫文字
			g.drawTileText(screen, value, row, col)
		}
	}
}

// drawCoverOnGameOver - 畫出無法操作的灰色遮罩
func (g *GameLayout) drawCoverOnGameOver(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	vector.DrawFilledRect(
		screen,
		0, 0, // x, y
		float32(w), float32(h), // width, height
		color.RGBA{0, 0, 0, 128}, // 半透明黑色 (128 = 約 50% 透明)
		false,
	)
}

// drawGameOver 畫出 GameOver
func (g *GameLayout) drawGameOver(screen *ebiten.Image) {
	g.drawCoverOnGameOver(screen)
	// 設定顯示 Game Over 文字
	textXPos := WinHeight / 2
	textYPos := WinWidth / 2
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.RGBA{128, 200, 200, 255})
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(textXPos), float64(textYPos))
	text.Draw(screen, "Game Over", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   48.0,
	}, textOpts)
	g.drawRestartButton(screen)
}

// drawRestartButton - 畫出 restart button
func (g *GameLayout) drawRestartButton(screen *ebiten.Image) {
	// 畫 Restart 按鈕
	vector.DrawFilledRect(screen,
		float32(restartButtonRect.Min.X),
		float32(restartButtonRect.Min.Y),
		float32(restartButtonRect.Dx()),
		float32(restartButtonRect.Dy()),
		color.RGBA{200, 200, 200, 255},
		true,
	)
	textOpts := &text.DrawOptions{}
	textOpts.ColorScale.ScaleWithColor(color.Black)
	textOpts.PrimaryAlign = text.AlignCenter
	textOpts.SecondaryAlign = text.AlignCenter
	textOpts.GeoM.Translate(float64(restartButtonRect.Min.X+restartButtonRect.Dx()/2), float64(restartButtonRect.Min.Y+restartButtonRect.Dy()/2))
	text.Draw(screen, "Restart", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30,
	}, textOpts)
}

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WinWidth, WinHeight
}

func NameGameLayout(gameInstance *game.Game) *GameLayout {
	return &GameLayout{
		gameInstance: gameInstance,
	}
}
