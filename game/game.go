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
	ScreenWidth  = 1000
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

	cards        []*ebiten.Image
	DrawCards    []*ebiten.Image
	HandCards    []*ebiten.Image
	DiscardCards []*ebiten.Image

	showCard bool

	testCount int
	testHp    int
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

	allCards := getCards()

	return &Game{
		cards:            allCards,
		DrawCards:        allCards,
		character:        cha,
		characterHp:      99,
		characterHpLimit: 99,
		enemy:            ene,
		enemyHp:          30,
		enemyHpLimit:     30,

		draggingIndex: -1,
		expandIndex:   -1,
		isDragging:    false,
	}, nil
}

func (g *Game) Update() error {
	sendCards(g)

	changeStatus(g)

	changeHpRand(g)

	// checkCardDrag(g)

	// checkRefreshButtonClick(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCharAEnemy(g, screen)
	drawManyCards(g, screen)

	drawText(g, screen)
	drawSendButton(screen)

	DrawHealthBar(g, screen, 600, 400, 100)
	// drawRefreshButton(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// // 新卡牌再次刷出来
// func (g *Game) RefreshCard() {
// 	g.cardX = g.cardOriginalX
// 	g.cardY = g.cardOriginalY
// 	g.isDragging = false
// 	g.showCard = true
// }
