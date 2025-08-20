package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize  = 100
	gridSize  = 4
	padding   = 10
	WinWidth  = tileSize*gridSize + padding*(gridSize+1)
	WinHeight = tileSize*gridSize + padding*(gridSize+1)
)

func (g *Game) Draw(screen *ebiten.Image) {
	// 背景色
	screen.Fill(color.RGBA{250, 248, 239, 255})
	// 畫 4x4 格子
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			op := &ebiten.DrawImageOptions{}
			x := padding + col*(tileSize+padding)
			y := padding + row*(tileSize+padding)
			rect := ebiten.NewImage(tileSize, tileSize)
			rect.Fill(color.RGBA{205, 193, 180, 255})
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(rect, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WinWidth, WinHeight
}

func (g *Game) Update() error {
	return nil
}
