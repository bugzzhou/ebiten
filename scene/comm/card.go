package comm

import "github.com/hajimehoshi/ebiten/v2"

// 唯一标识一个卡牌
type CardInfo struct {
	Id int

	Attack     int
	Shield     int
	SelfAttack int
	Cost       int

	Image *ebiten.Image
}
