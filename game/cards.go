package game

import (
	"fmt"
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

var (
	cardDir = "./game/pic/cards"
)

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap = map[string]*ebiten.Image{}

func init() {
	files, ids, err := listDir(cardDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}
	fmt.Println(files)
	fmt.Println(ids)

	for i := range files {
		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		cardImageMap[ids[i]] = tmpImage
	}

	// for

	// cha, _, err := ebitenutil.NewImageFromFile(lieren) // 替换为你的角色图像文件路径

	// ene, _, err := ebitenutil.NewImageFromFile(kaka) // 替换为你的敌人图像文件路径

	// card, _, err := ebitenutil.NewImageFromFile(cardSample) // 替换为你的卡牌图像文件路径

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

func getCards() []*ebiten.Image {
	attCard, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "1.jpg"))

	useless1, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "2.jpg"))

	useless2, _, _ := ebitenutil.NewImageFromFile(filepath.Join(dir, "3.jpg"))

	return []*ebiten.Image{
		attCard, attCard, attCard, attCard, attCard,
		useless1, useless1, useless2, useless2, useless2,
	}
}
