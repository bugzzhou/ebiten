package main

import (
	"ebiten/scene"
	"ebiten/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	manager := scene.NewSceneManager()
	manager.SetScene(scene.NewMapScene(manager))

	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Scene Switch Example")
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
