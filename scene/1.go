package scene

import (
	cons "ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 该场景暂时用于表示失败后的跳转
type Scene1 struct {
	manager *SceneManager
}

func NewScene1(manager *SceneManager) *Scene1 {
	return &Scene1{manager: manager}
}

func (s *Scene1) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewMapScene(s.manager))
	}
	return nil
}

func (s *Scene1) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "failed")
}

func (s *Scene1) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}
