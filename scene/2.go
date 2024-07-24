package scene

import (
	cons "ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 该场景暂时用于表示成功后的跳转
type Scene2 struct {
	manager *SceneManager
}

func NewScene2(manager *SceneManager) *Scene2 {
	return &Scene2{manager: manager}
}

func (s *Scene2) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewMapScene(s.manager))
	}
	return nil
}

func (s *Scene2) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "succeed")
}

func (s *Scene2) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}
