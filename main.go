package main

import (
	"ebiten/scene"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	manager := scene.NewSceneManager()
	manager.SetScene(scene.NewMapScene(manager))

	ebiten.SetWindowSize(scene.ScreenWidth, scene.ScreenHeight)
	ebiten.SetWindowTitle("Scene Switch Example")
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
