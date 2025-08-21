package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
	"github.com/leetcode-golang-classroom/2048-game/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.WinWidth, layout.WinHeight)
	ebiten.SetWindowTitle("2048 - Day 8 測試")
	gameInstance := game.NewGame()
	gameInstance.Init([][]int{
		{2, 4, 8, 16},
		{32, 64, 128, 256},
		{512, 1024, 2048, 4096},
		{0, 0, 0, 8192},
	}, nil, nil)
	gameLayout := layout.NameGameLayout(gameInstance)
	if err := ebiten.RunGame(gameLayout); err != nil {
		log.Fatal(err)
	}
}
