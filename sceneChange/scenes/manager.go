package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

type SceneManager struct {
	currentScene Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

func (sm *SceneManager) Update() error {
	if sm.currentScene != nil {
		return sm.currentScene.Update()
	}
	return nil
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	if sm.currentScene != nil {
		sm.currentScene.Draw(screen)
	}
}

func (sm *SceneManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	if sm.currentScene != nil {
		return sm.currentScene.Layout(outsideWidth, outsideHeight)
	}
	return outsideWidth, outsideHeight
}

func (sm *SceneManager) SetScene(scene Scene) {
	sm.currentScene = scene
}
