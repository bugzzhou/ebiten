package comm

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)

var (
	LocalCharacter = Character{}
	LocalEnemy     = Enemy{}
)

func GetLocalCharacter() *Character {
	LocalCharacter.Energy = 3
	LocalCharacter.DrawDeck = LocalCharacter.Cards
	LocalCharacter.HandCards = nil
	LocalCharacter.DiscardDeck = nil
	return &LocalCharacter
}

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap map[int]*ebiten.Image
