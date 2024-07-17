package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene3 struct {
	manager *SceneManager
}

func NewScene3(manager *SceneManager) *Scene3 {
	return &Scene3{manager: manager}
}

func (s *Scene3) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewScene1(s.manager))
	}
	return nil
}

func (s *Scene3) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Scene 3")
}

func (s *Scene3) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
