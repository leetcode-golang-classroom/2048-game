package layout

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
)

const (
	tileSize  = 100
	gridSize  = 4
	padding   = 10
	WinWidth  = tileSize*gridSize + padding*(gridSize+1)
	WinHeight = tileSize*gridSize + padding*(gridSize+1)
)

type GameLayout struct {
	gameInstance *game.Game
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
	// 畫 4x4 格子
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

func (g *GameLayout) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WinWidth, WinHeight
}

func NameGameLayout(gameInstance *game.Game) *GameLayout {
	return &GameLayout{
		gameInstance: gameInstance,
	}
}
