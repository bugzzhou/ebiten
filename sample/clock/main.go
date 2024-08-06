package main

import (
	"log"

	"clock/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetTPS(2)

	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Clock")
	game, err := utils.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
