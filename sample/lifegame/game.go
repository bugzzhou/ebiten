package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 100
)

var (
	grid       = make([][]bool, gridSize)
	cellSize   = screenWidth / gridSize
	stepButton = &Button1{
		x:      screenWidth - 300,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Step",
	}
	runButton = &Button1{
		x:      screenWidth - 200,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Run",
	}
	stopButton = &Button1{
		x:      screenWidth - 100,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Stop",
	}
	mousePressedLastFrame bool
	changedCells          = make(map[Point]bool)
	running               bool
	ticker                *time.Ticker
)

type Point struct {
	X, Y int
}

type Button1 struct {
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
	drawUI(screen)

	stepButton.Draw(screen)
	runButton.Draw(screen)
	stopButton.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
