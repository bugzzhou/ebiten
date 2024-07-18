package scene

import (
	cons "ebiten/scene/const"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 整个游戏的入口，用于选择地图

type MapScene struct {
	manager *SceneManager
}

func NewMapScene(manager *SceneManager) *MapScene {
	return &MapScene{manager: manager}
}

func (s *MapScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.manager.SetScene(NewCombatScene(s.manager))
	}
	return nil
}

func (s *MapScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Scene map - there is a map to choose in it~~~")
}

func (s *MapScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}
