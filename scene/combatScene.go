package scene

import (
	sceneCom "ebiten/scene/combatScene"
	"ebiten/scene/comm"

	"github.com/hajimehoshi/ebiten/v2"
)

// 战斗的场景，从map跳转过来。 失败、成功、分别跳转到场景3、4
type CombatScene struct {
	manager *SceneManager
	sceneCom.Game
}

func NewCombatScene(manager *SceneManager) *CombatScene {
	localCharacter := comm.GetLocalCharacter()
	enemy := comm.GetEnemy()
	gametmp := &sceneCom.Game{
		Round:         1,
		Character:     localCharacter,
		Enemy:         *enemy,
		RoundBegin:    true,
		DraggingIndex: -1,
		ExpandIndex:   -1,
		IsDragging:    false,
	}
	return &CombatScene{
		manager: manager,
		Game:    *gametmp,
	}
}

func (cs *CombatScene) Update() error {
	g := &cs.Game
	if g.RoundBegin {
		g.RoundBegin = false
		sceneCom.SendCards(g)
	}

	sceneCom.EndCards(g)
	sceneCom.ChangeStatus(g)
	ChangeScene(cs)
	return nil
}

func (cs *CombatScene) Draw(screen *ebiten.Image) {
	g := &cs.Game
	sceneCom.DrawCharAEnemy(g, screen)
	sceneCom.DrawManyCards(g, screen)
	sceneCom.DrawText(g, screen)
	sceneCom.EndTurnButton(screen)
}

func (g *CombatScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return comm.ScreenWidth, comm.ScreenHeight
}

// 胜利或者失败，跳转不同的场景
func ChangeScene(cs *CombatScene) {
	if cs.Character.Hp <= 0 {
		cs.manager.SetScene(NewScene1(cs.manager))
	}

	if cs.Enemy.Hp <= 0 {
		cs.manager.SetScene(NewScene2(cs.manager))
	}

}
