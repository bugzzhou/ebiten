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
	Shield  int
	Energy  int

	Cards       []CardInfo
	DrawDeck    []CardInfo
	HandCards   []CardInfo
	DiscardDeck []CardInfo
}

func (c *Character) Shuffle() {
	rand.Shuffle(len(c.DrawDeck), func(i, j int) {
		c.DrawDeck[i], c.DrawDeck[j] = c.DrawDeck[j], c.DrawDeck[i]
	})

}

func (c *Character) DrawCard(drawNum int) {
	// 弃牌堆 - > 抽牌堆
	if len(c.DrawDeck) < drawNum {
		c.DrawDeck = append(c.DrawDeck, c.DiscardDeck...)
		c.DiscardDeck = nil
		R.Shuffle(len(c.DrawDeck), func(i, j int) {
			c.DrawDeck[i], c.DrawDeck[j] = c.DrawDeck[j], c.DrawDeck[i]
		})
	}

	// 抽牌堆 - > 手牌
	if len(c.DrawDeck) <= drawNum {
		c.HandCards = append(c.HandCards, c.DrawDeck...)
		c.DrawDeck = nil
	} else {
		c.HandCards = append(c.HandCards, c.DrawDeck[:drawNum]...)
		c.DrawDeck = c.DrawDeck[drawNum:]
	}
}

func (c *Character) EndTurn() {
	c.DiscardDeck = append(c.DiscardDeck, c.HandCards...)
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
		Id:         1,
		Attack:     5,
		Shield:     0,
		SelfAttack: 0,
		Cost:       1,
		Image:      att5,
	}

	get2, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "2.jpg"))
	c2 := CardInfo{
		Id:    2,
		Cost:  1,
		Image: get2,
	}

	att20, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "3.jpg"))
	c3 := CardInfo{
		Id:         3,
		Attack:     20,
		SelfAttack: 2,
		Cost:       2,
		Image:      att20,
	}

	get4, _, _ := ebitenutil.NewImageFromFile(filepath.Join(cardDir, "4.jpg"))
	c4 := CardInfo{
		Id:    4,
		Cost:  4,
		Image: get4,
	}

	return []CardInfo{
		c1, c1, c1, c1, c1,
		c2, c2, c4, c3, c3,
	}
}

func (c *Character) CardAffect(index int, enemy *Enemy) {
	card := c.HandCards[index]
	affectByCard(c, enemy, &card)

	if card.Id == 2 {
		c.DrawCard(2)
	} else if card.Id == 4 {
		c.DrawCard(4)
	}

	// if card.Id == 1 {
	// 	enemy.Hp -= 5
	// 	c.Energy -= 1
	// } else if card.Id == 2 {

	// 	c.Energy -= 1
	// } else if card.Id == 3 {
	// 	enemy.Hp -= 20
	// 	c.Hp -= 2
	// 	c.Energy -= 1
	// } else if card.Id == 4 {

	// 	c.Energy -= 2
	// }
}

func (c *Character) CardDiscard(index int) {
	c.DiscardDeck = append(c.DiscardDeck, c.HandCards[index])
	c.HandCards = append(c.HandCards[:index], c.HandCards[index+1:]...)
}

func affectByCard(c *Character, e *Enemy, card *CardInfo) {
	c.Shield += card.Shield
	fmt.Printf("%v, %v, %v\n", card.Attack, e.Shield, e.Hp)
	c.Hp -= card.SelfAttack
	c.Energy -= card.Cost

	if card.Attack < e.Shield {
		e.Shield -= card.Attack
	} else {
		e.Hp -= (card.Attack - e.Shield)
		e.Shield = 0
	}
	fmt.Printf("%v, %v, %v\n", card.Attack, e.Shield, e.Hp)
}
