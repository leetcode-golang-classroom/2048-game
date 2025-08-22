package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/2048-game/internal/game"
	"github.com/leetcode-golang-classroom/2048-game/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.WinWidth, layout.WinHeight)
	ebiten.SetWindowTitle("2048 - Day 9 測試")
	gameInstance := game.NewGame()
	gameInstance.AddRandomTile(game.Default)
	gameInstance.AddRandomTile(game.Default)
	gameLayout := layout.NameGameLayout(gameInstance)
	if err := ebiten.RunGame(gameLayout); err != nil {
		log.Fatal(err)
	}
}
