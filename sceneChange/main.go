package main

import (
	"log"
	scenes "t1/scenes" // 替换为实际的包路径

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	manager := scenes.NewSceneManager()
	manager.SetScene(scenes.NewScene1(manager))

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Scene Switch Example")
	if err := ebiten.RunGame(manager); err != nil {
		log.Fatal(err)
	}
}
