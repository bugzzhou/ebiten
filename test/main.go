package main

import (
	"log"

	"test/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame(initialScene scene.Scene) *Game {
	return &Game{currentScene: initialScene}
}

type Game struct {
	currentScene scene.Scene
}

func (g *Game) Update() error {
	return g.currentScene.Update(g.changeScene)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

// changeScene 切换场景的方法
func (g *Game) changeScene(newScene scene.Scene) {
	g.currentScene = newScene
}

func main() {
	// 初始化当前场景为 SceneA
	currentScene := &scene.SceneA{}
	sceneManager := NewGame(currentScene)

	// 创建游戏实例
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Scene Switching Example")
	if err := ebiten.RunGame(sceneManager); err != nil {
		log.Fatal(err)
	}
}
