// package scene

// import (
// 	cons "ebiten/scene/const"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

// // 整个游戏的入口，用于选择地图

// type MapScene struct {
// 	manager *SceneManager
// }

// func NewMapScene(manager *SceneManager) *MapScene {
// 	return &MapScene{manager: manager}
// }

// func (s *MapScene) Update() error {
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		s.manager.SetScene(NewCombatScene(s.manager))
// 	}
// 	return nil
// }

// func (s *MapScene) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Scene map - there is a map to choose in it~~~")
// }

// func (s *MapScene) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return cons.ScreenWidth, cons.ScreenHeight
// }

package scene

import (
	cons "ebiten/scene/comm"
	ms "ebiten/scene/mapScene"
	m "ebiten/scene/models"

	"github.com/hajimehoshi/ebiten/v2"
)

// 整个游戏的入口，用于选择地图

type MapScene struct {
	manager *SceneManager
	graph   *m.Graph
}

func NewMapScene(manager *SceneManager) *MapScene {
	return &MapScene{
		manager: manager,
		graph:   ms.CreateSampleGraph(),
	}
}

func (s *MapScene) Update() error {
	nodeType := ms.CheckNodeSite(s.graph.Nodes)
	if nodeType != -1 {
		ChooseMap(s, nodeType)
	}
	return nil
}

func (s *MapScene) Draw(screen *ebiten.Image) {
	ms.DrawMapNode(screen, s.graph.Nodes)
}

func (s *MapScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}

func ChooseMap(s *MapScene, nodeType int) {
	switch nodeType {
	case ms.CombatNode:
		s.manager.SetScene(NewCombatScene(s.manager))
	case ms.CampfireNode:
		s.manager.SetScene(NewCampfireScene(s.manager))
	default:
		s.manager.SetScene(NewScene1(s.manager)) //需要一个默认的页面， 后续增加新的页面
	}
}
