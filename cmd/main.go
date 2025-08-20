package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/2048-game/internal/layout"
)

func main() {
	ebiten.SetWindowSize(layout.WinWidth, layout.WinHeight)
	ebiten.SetWindowTitle("2048 - Day 7 測試")
	game := layout.NameGameLayout()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
