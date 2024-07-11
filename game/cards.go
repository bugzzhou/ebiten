package game

import (
	"math/rand"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	dir        = "./game/pic"
	lieren     = "./game/pic/lieren.jpg"
	kaka       = "./game/pic/kaka.jpg"
	cardSample = "./game/pic/1.jpg"
)

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

func getCards() []*ebiten.Image {
	attCard, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "1.jpg"))

	useless1, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "2.jpg"))

	useless2, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "3.jpg"))

	return []*ebiten.Image{
		attCard, attCard, attCard, attCard, attCard,
		useless1, useless1, useless2, useless2, useless2,
	}
}
