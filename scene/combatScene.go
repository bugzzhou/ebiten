package scene

import (
	combatscene "ebiten/scene/combatScene"
	cs "ebiten/scene/combatScene"
	cons "ebiten/scene/const"
	"ebiten/scene/models"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 战斗的场景，从map跳转过来。 失败、成功、分别跳转到场景3、4
type CombatScene struct {
	manager *SceneManager
	combatscene.Game
}

func NewCombatScene(manager *SceneManager) *CombatScene {
	cha, _, err := ebitenutil.NewImageFromFile(cs.Lieren) // 猎人的图片
	if err != nil {
		return nil
	}
	ene, _, err := ebitenutil.NewImageFromFile(cs.Kaka) // kaka的图片
	if err != nil {
		return nil
	}

	allCards := cs.GetCards()

	gametmp := &cs.Game{
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
			Action:  combatscene.GetActs(combatscene.KakaActTag),
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

	combatscene.SendCards(g)
	combatscene.EndCards(g)

	combatscene.ChangeStatus(g)

	//kaka的行动判断
	combatscene.KakaAct(g)

	ChangeScene(cs)

	return nil
}

func (cs *CombatScene) Draw(screen *ebiten.Image) {
	g := &cs.Game
	combatscene.DrawCharAEnemy(g, screen)
	combatscene.DrawManyCards(g, screen)
	combatscene.DrawText(g, screen)
	combatscene.DrawSendButton(screen)
	combatscene.EndTurnButton(screen)

	//kaka的行为按钮
	combatscene.KakaActButton(screen)

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
