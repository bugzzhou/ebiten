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
	character        *ebiten.Image
	characterHp      int
	characterHpLimit int
	enemy            *ebiten.Image
	enemyHp          int
	enemyHpLimit     int

	expandIndex   int
	draggingIndex int
	isDragging    bool

	cards        []CardInfo
	DrawCards    []CardInfo
	HandCards    []CardInfo
	DiscardCards []CardInfo
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
		character:        cha,
		characterHp:      99,
		characterHpLimit: 99,
		enemy:            ene,
		enemyHp:          300, //写大点方便多牌演示
		enemyHpLimit:     300,

		cards:     allCards,
		DrawCards: allCards,

		draggingIndex: -1,
		expandIndex:   -1,
		isDragging:    false,
	}, nil
}

func (g *Game) Update() error {
	sendCards(g)

	changeStatus(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCharAEnemy(g, screen)
	drawManyCards(g, screen)
	drawText(g, screen)
	drawSendButton(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
