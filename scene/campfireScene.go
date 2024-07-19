package scene

import (
	campfirescene "ebiten/scene/campfireScene"
	cons "ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 火堆场景
type CampfireScene struct {
	manager *SceneManager
}

func NewCampfireScene(manager *SceneManager) *Scene1 {
	return &Scene1{manager: manager}
}

func (s *CampfireScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewMapScene(s.manager))
	}
	return nil
}

func (s *CampfireScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "campfire scene--------\nyou can choose sth to do here...")
	campfirescene.DrawCampfire(screen)
}

func (s *CampfireScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}
