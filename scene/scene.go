package scene

import (
	cons "ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

//上面是一个interface的接口，后续每个场景都需要实现上述接口

//下面是一个管理scene的类，主要适用于切换场景

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
		return sm.currentScene.Layout(cons.ScreenWidth, cons.ScreenHeight)
	}
	return cons.ScreenWidth, cons.ScreenHeight
}

func (sm *SceneManager) SetScene(scene Scene) {
	sm.currentScene = scene
}
