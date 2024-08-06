package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// 函数部分
type Game struct{}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func (g *Game) Update() error {
	changeClock()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawUI(screen, clockMap)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
