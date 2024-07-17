package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene1 struct {
	manager *SceneManager
}

func NewScene1(manager *SceneManager) *Scene1 {
	return &Scene1{manager: manager}
}

func (s *Scene1) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewScene2(s.manager))
	}
	return nil
}

func (s *Scene1) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Scene 1")
}

func (s *Scene1) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
