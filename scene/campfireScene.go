package scene

import (
	campfirescene "ebiten/scene/campfireScene"
	cons "ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 火堆场景
type CampfireScene struct {
	manager *SceneManager
}

func NewCampfireScene(manager *SceneManager) *CampfireScene {
	return &CampfireScene{manager: manager}
}

func (s *CampfireScene) Update() error {
	ok := campfirescene.Recover(cons.GetLocalCharacter())
	if ok {
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
