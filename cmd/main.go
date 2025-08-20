package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leetcode-golang-classroom/2048-game/internal"
)

func main() {
	ebiten.SetWindowSize(internal.WinWidth, internal.WinHeight)
	ebiten.SetWindowTitle("2048 - Day 7 測試")
	game := internal.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
