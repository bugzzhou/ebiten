package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 100
)

var (
	grid        = make([][]bool, gridSize)
	cellSize    = screenWidth / gridSize
	startButton = &Button{
		x:      screenWidth - 100,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Start",
	}
	mousePressedLastFrame bool
	changedCells          = make(map[Point]bool)
)

type Point struct {
	X, Y int
}

type Button struct {
	x, y          int
	width, height int
	label         string
}

// 函数部分
type Game struct{}

func init() {
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}
}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func (g *Game) Update() error {
	changeStatus()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawUI(screen)

	startButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
