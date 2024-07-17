package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 战斗的场景，从map跳转过来。 失败、成功、分别跳转到场景3、4
type CombatScene struct {
	manager *SceneManager
	Game
}

type Character struct {
	image   *ebiten.Image
	hp      int
	hplimit int
	energy  int
}

type Enemy struct {
	image   *ebiten.Image
	hp      int
	hplimit int
	action  []Act
	buffs   []Buff
}

type Act struct {
	Id          int
	Name        string
	Description string
}

type Buff struct {
	Id          int
	Name        string
	Description string

	Layers     int
	StartRound int
	EndRound   int
}

func NewCombatScene(manager *SceneManager) *CombatScene {
	cha, _, err := ebitenutil.NewImageFromFile(lieren) // 猎人的图片
	if err != nil {
		return nil
	}
	ene, _, err := ebitenutil.NewImageFromFile(kaka) // kaka的图片
	if err != nil {
		return nil
	}

	allCards := getCards()

	gametmp := &Game{
		round: 1,

		character: Character{
			image:   cha,
			hp:      99,
			hplimit: 99,
			energy:  3,
		},
		enemy: Enemy{
			image:   ene,
			hp:      30,
			hplimit: 30, //写大点方便多牌演示
			action:  getActs(kakaActTag),
		},

		cards:     allCards,
		DrawCards: allCards,

		draggingIndex: -1,
		expandIndex:   -1,
		isDragging:    false,
	}

	return &CombatScene{
		manager: manager,
		Game:    *gametmp,
	}
}

func (cs *CombatScene) Update() error {
	g := &cs.Game

	sendCards(g)
	endCards(g)

	changeStatus(g)

	//kaka的行动判断
	kakaAct(g)

	changeScene(cs)

	return nil
}

func (cs *CombatScene) Draw(screen *ebiten.Image) {
	g := &cs.Game
	drawCharAEnemy(g, screen)
	drawManyCards(g, screen)
	drawText(g, screen)
	drawSendButton(screen)
	endTurnButton(screen)

	//kaka的行为按钮
	kakaActButton(screen)

}

func (g *CombatScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
