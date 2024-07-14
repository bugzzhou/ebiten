package game

import (
	_ "image/jpeg"
	_ "image/png"
	"time"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
	imageWidth   = 150
	imageHeight  = 200
)

type Game struct {
	character Character

	enemy Enemy

	expandIndex   int
	draggingIndex int
	isDragging    bool

	cards        []CardInfo
	DrawCards    []CardInfo
	HandCards    []CardInfo
	DiscardCards []CardInfo

	round int //回合数，后续用于计算buff的生效数值
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

func NewGame() (*Game, error) {
	cha, _, err := ebitenutil.NewImageFromFile(lieren) // 猎人的图片
	if err != nil {
		return nil, err
	}
	ene, _, err := ebitenutil.NewImageFromFile(kaka) // kaka的图片
	if err != nil {
		return nil, err
	}

	allCards := getCards()

	return &Game{
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
	}, nil
}

func (g *Game) Update() error {
	sendCards(g)
	endCards(g)

	changeStatus(g)

	//kaka的行动判断
	kakaAct(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCharAEnemy(g, screen)
	drawManyCards(g, screen)
	drawText(g, screen)
	drawSendButton(screen)
	endTurnButton(screen)

	//kaka的行为按钮
	kakaActButton(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
