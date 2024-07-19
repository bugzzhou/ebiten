package comm

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	cardDir = "./pic/cards"
)

// combatScene 的结构体
type Character struct {
	Image   *ebiten.Image
	Hp      int
	Hplimit int
	Energy  int

	Cards        []CardInfo
	DrawCards    []CardInfo
	HandCards    []CardInfo
	DiscardCards []CardInfo
}

func (c *Character) Shuffle() {
	rand.Shuffle(len(c.DrawCards), func(i, j int) {
		c.DrawCards[i], c.DrawCards[j] = c.DrawCards[j], c.DrawCards[i]
	})

}

func (c *Character) DrawCard(drawNum int) {
	// 弃牌堆 - > 抽牌堆
	if len(c.DrawCards) < drawNum {
		c.DrawCards = append(c.DrawCards, c.DiscardCards...)
		c.DiscardCards = nil
		R.Shuffle(len(c.DrawCards), func(i, j int) {
			c.DrawCards[i], c.DrawCards[j] = c.DrawCards[j], c.DrawCards[i]
		})
	}

	// 抽牌堆 - > 手牌
	if len(c.DrawCards) <= drawNum {
		c.HandCards = append(c.HandCards, c.DrawCards...)
		c.DrawCards = nil
	} else {
		c.HandCards = append(c.HandCards, c.DrawCards[:drawNum]...)
		c.DrawCards = c.DrawCards[drawNum:]
	}
}

func (c *Character) EndTurn() {
	c.DiscardCards = append(c.DiscardCards, c.HandCards...)
	c.HandCards = nil
}

func (c *Character) PlayCard(index int, enemy *Enemy) {
	if len(c.HandCards) < index {
		fmt.Printf("failed to play card, out of length of the handCard\n")
		return
	}

	c.CardAffect(index, enemy)
	c.CardDiscard(index)
}

func GetCards() []CardInfo {
	att5, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "1.jpg"))
	c1 := CardInfo{
		Id:    1,
		Image: att5,
	}

	get2, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "2.jpg"))
	c2 := CardInfo{
		Id:    2,
		Image: get2,
	}

	att20, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "3.jpg"))
	c3 := CardInfo{
		Id:    3,
		Image: att20,
	}

	get4, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "4.jpg"))
	c4 := CardInfo{
		Id:    4,
		Image: get4,
	}

	return []CardInfo{
		c1, c1, c1, c1, c1,
		c2, c2, c4, c3, c3,
	}
}

func (c *Character) CardAffect(index int, enemy *Enemy) {
	card := c.HandCards[index]
	if card.Id == 1 {
		enemy.Hp -= 5
		c.Energy -= 1
	} else if card.Id == 2 {
		c.DrawCard(2)
		c.Energy -= 1
	} else if card.Id == 3 {
		enemy.Hp -= 20
		c.Hp -= 2
		c.Energy -= 1
	} else if card.Id == 4 {
		c.DrawCard(4)
		c.Energy -= 2
	}
}

func (c *Character) CardDiscard(index int) {
	c.DiscardCards = append(c.DiscardCards, c.HandCards[index])
	c.HandCards = append(c.HandCards[:index], c.HandCards[index+1:]...)
}
