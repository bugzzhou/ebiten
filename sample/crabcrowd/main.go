package main

import (
	"crabcrowd/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		panic(err)
	}
	ebiten.SetWindowSize(game.ScreenW, game.ScreenH)
	ebiten.SetWindowTitle("螃蟹大挪移")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
