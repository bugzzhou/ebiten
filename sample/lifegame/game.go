package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func (g *Game) Update() error {
	changeStatus()
	if running {
		select {
		case <-ticker.C:
			updateGrid()
		default:
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawGrid(screen)

	drawButtons(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
