package combatscene

import (
	"fmt"
	"math/rand"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// dir        = "./game/pic"
	Lieren = "./pic/lieren.jpg"
	Kaka   = "./pic/kaka.jpg"
)

var (
	cardDir = "./pic/cards"
)

// 唯一标识一个卡牌
type CardInfo struct {
	id    int
	image *ebiten.Image
}

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap = map[string]*ebiten.Image{}

func init() {
	files, ids, err := listDir(cardDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}

	for i := range files {
		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		cardImageMap[ids[i]] = tmpImage
	}
}

func (g *Game) Shuffle() {
	rand.Shuffle(len(g.DrawCards), func(i, j int) {
		g.DrawCards[i], g.DrawCards[j] = g.DrawCards[j], g.DrawCards[i]
	})

}

func (g *Game) DrawCard(drawNum int) {
	// 弃牌堆 - > 抽牌堆
	if len(g.DrawCards) < drawNum {
		g.DrawCards = append(g.DrawCards, g.DiscardCards...)
		g.DiscardCards = nil
		R.Shuffle(len(g.DrawCards), func(i, j int) {
			g.DrawCards[i], g.DrawCards[j] = g.DrawCards[j], g.DrawCards[i]
		})
	}

	// 抽牌堆 - > 手牌
	if len(g.DrawCards) <= drawNum {
		g.HandCards = append(g.HandCards, g.DrawCards...)
		g.DrawCards = nil
	} else {
		g.HandCards = append(g.HandCards, g.DrawCards[:drawNum]...)
		g.DrawCards = g.DrawCards[drawNum:]
	}
}

func (g *Game) EndTurn() {
	g.DiscardCards = append(g.DiscardCards, g.HandCards...)
	g.HandCards = nil
}

func (g *Game) PlayCard(index int) {
	if len(g.HandCards) < index {
		fmt.Printf("failed to play card, out of length of the handCard\n")
		return
	}

	CardAffect(g, index)
	CardDiscard(g, index)
}

func GetCards() []CardInfo {
	att5, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "1.jpg"))
	c1 := CardInfo{
		id:    1,
		image: att5,
	}

	get2, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "2.jpg"))
	c2 := CardInfo{
		id:    2,
		image: get2,
	}

	att20, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "3.jpg"))
	c3 := CardInfo{
		id:    3,
		image: att20,
	}

	get4, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "4.jpg"))
	c4 := CardInfo{
		id:    4,
		image: get4,
	}

	return []CardInfo{
		c1, c1, c1, c1, c1,
		c2, c2, c4, c3, c3,
	}
}

func CardAffect(g *Game, index int) {
	c := g.HandCards[index]
	if c.id == 1 {
		g.Enemy.Hp -= 5
		g.Character.Energy -= 1
	} else if c.id == 2 {
		g.DrawCard(2)
		g.Character.Energy -= 1
	} else if c.id == 3 {
		g.Enemy.Hp -= 20
		g.Character.Hp -= 2
		g.Character.Energy -= 1
	} else if c.id == 4 {
		g.DrawCard(4)
		g.Character.Energy -= 2
	}
}

func CardDiscard(g *Game, index int) {
	g.DiscardCards = append(g.DiscardCards, g.HandCards[index])
	g.HandCards = append(g.HandCards[:index], g.HandCards[index+1:]...)
}
