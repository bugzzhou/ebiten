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
	ScreenWidth  = 800
	ScreenHeight = 600
	imageWidth   = 150
	imageHeight  = 200
)

type Game struct {
	character *ebiten.Image
	enemy     *ebiten.Image

	cards        []*ebiten.Image
	DrawCards    []*ebiten.Image
	HandCards    []*ebiten.Image
	DiscardCards []*ebiten.Image

	counter1 int
	counter2 int
	counter3 int
}

func NewGame() (*Game, error) {
	cha, _, err := ebitenutil.NewImageFromFile(lieren)
	if err != nil {
		return nil, err
	}
	ene, _, err := ebitenutil.NewImageFromFile(kaka)
	if err != nil {
		return nil, err
	}

	allCards := getCards()

	return &Game{
		cards:     allCards,
		DrawCards: allCards,
		character: cha,
		enemy:     ene,
	}, nil
}

func (g *Game) Update() error {
	// 检测鼠标点击
	checkAddClick(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawCharAEnemy(g, screen)

	drawText(g, screen)

	drawSendButton(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
