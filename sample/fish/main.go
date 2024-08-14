package main

import (
	"fish/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := utils.NewGame()
	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Ocean Simulation")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
