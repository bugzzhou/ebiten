package scene

import (
	sceneCom "ebiten/scene/combatScene"
	cons "ebiten/scene/comm"
	"ebiten/scene/models"

	"github.com/hajimehoshi/ebiten/v2"
)

// 战斗的场景，从map跳转过来。 失败、成功、分别跳转到场景3、4
type CombatScene struct {
	manager *SceneManager
	sceneCom.Game
}

func NewCombatScene(manager *SceneManager) *CombatScene {

	allCards := sceneCom.GetCards()

	gametmp := &sceneCom.Game{
		Round: 1,

		Character: models.Character{
			Image:   cha,
			Hp:      99,
			Hplimit: 99,
			Energy:  3,
		},
		Enemy: models.Enemy{
			Image:   ene,
			Hp:      30,
			Hplimit: 30, //写大点方便多牌演示
			Action:  sceneCom.GetActs(sceneCom.KakaActTag),
		},

		Cards:     allCards,
		DrawCards: allCards,

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

	sceneCom.SendCards(g)
	sceneCom.EndCards(g)

	sceneCom.ChangeStatus(g)

	//kaka的行动判断
	sceneCom.KakaAct(g)

	ChangeScene(cs)

	return nil
}

func (cs *CombatScene) Draw(screen *ebiten.Image) {
	g := &cs.Game
	sceneCom.DrawCharAEnemy(g, screen)
	sceneCom.DrawManyCards(g, screen)
	sceneCom.DrawText(g, screen)
	sceneCom.DrawSendButton(screen)
	sceneCom.EndTurnButton(screen)

	//kaka的行为按钮
	sceneCom.KakaActButton(screen)

}

func (g *CombatScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
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
