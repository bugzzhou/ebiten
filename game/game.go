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
	character     *ebiten.Image
	enemy         *ebiten.Image
	card          *ebiten.Image
	cardX         float64
	cardY         float64
	cardOriginalX float64
	cardOriginalY float64
	isDragging    bool
	cards         []*ebiten.Image
	DrawCards     []*ebiten.Image
	HandCards     []*ebiten.Image
	DiscardCards  []*ebiten.Image
	showCard      bool
}

func NewGame() (*Game, error) {
	cha, _, err := ebitenutil.NewImageFromFile(lieren) // 替换为你的角色图像文件路径
	if err != nil {
		return nil, err
	}
	ene, _, err := ebitenutil.NewImageFromFile(kaka) // 替换为你的敌人图像文件路径
	if err != nil {
		return nil, err
	}
	card, _, err := ebitenutil.NewImageFromFile(cardSample) // 替换为你的卡牌图像文件路径
	if err != nil {
		return nil, err
	}
	allCards := getCards()

	return &Game{
		cards:         allCards,
		DrawCards:     allCards,
		character:     cha,
		enemy:         ene,
		card:          card,
		cardX:         ScreenWidth/2 - imageWidth/2,
		cardY:         ScreenHeight - imageHeight - 20,
		cardOriginalX: ScreenWidth/2 - imageWidth/2,
		cardOriginalY: ScreenHeight - imageHeight - 20,
		showCard:      true,
	}, nil
}

func (g *Game) Update() error {
	drawCards(g)
	checkCardDrag(g)
	checkRefreshButtonClick(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCharAEnemy(g, screen)
	if g.showCard {
		drawCard(g, screen)
	}
	drawText(g, screen)
	drawSendButton(screen)
	drawRefreshButton(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// 新卡牌再次刷出来
func (g *Game) RefreshCard() {
	g.cardX = g.cardOriginalX
	g.cardY = g.cardOriginalY
	g.isDragging = false
	g.showCard = true
}
