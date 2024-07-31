package comm

import (
	// "ebiten/utils"
	"ebiten/utils"
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
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

func (c *Character) DrawCard(drawNum int) {
	// 弃牌堆 - > 抽牌堆
	if len(c.DrawDeck) < drawNum {
		c.DrawDeck = append(c.DrawDeck, c.DiscardDeck...)
		c.DiscardDeck = nil
		utils.R.Shuffle(len(c.DrawDeck), func(i, j int) {
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

	c.cardAffect(index, enemy)
	c.cardDiscard(index)
}

func GetLocalCharacter() *Character {
	LocalCharacter.Energy = 3
	LocalCharacter.DrawDeck = LocalCharacter.Cards
	LocalCharacter.HandCards = nil
	LocalCharacter.DiscardDeck = nil
	return &LocalCharacter
}

func (c *Character) Shuffle() {
	rand.Shuffle(len(c.DrawDeck), func(i, j int) {
		c.DrawDeck[i], c.DrawDeck[j] = c.DrawDeck[j], c.DrawDeck[i]
	})
}

// 包内函数

/*
	※1 卡牌生效函数

1、卡牌基础属性通过配置实现，并且通过一个通用函数使用
2、卡牌的特殊效果需要专门实现。
TODO bugzzhou 卡牌的特殊效果，是否用一个函数再封装一层？
*/
func (c *Character) cardAffect(index int, enemy *Enemy) {
	card := c.HandCards[index]
	affectByCard(c, enemy, &card)

	if card.Id == 2 {
		c.DrawCard(2)
	} else if card.Id == 4 {
		c.DrawCard(4)
	}
}

func (c *Character) cardDiscard(index int) {
	c.DiscardDeck = append(c.DiscardDeck, c.HandCards[index])
	c.HandCards = append(c.HandCards[:index], c.HandCards[index+1:]...)
}

func affectByCard(c *Character, e *Enemy, card *CardInfo) {
	c.Shield += card.Shield
	c.Hp -= card.SelfAttack
	c.Energy -= card.Cost

	if card.Attack < e.Shield {
		e.Shield -= card.Attack
	} else {
		e.Hp -= (card.Attack - e.Shield)
		e.Shield = 0
	}
}
