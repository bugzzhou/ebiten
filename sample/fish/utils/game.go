package utils

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480

	CellSize = 4
)

var (
	simulationRunning bool
	ocean             Ocean
)

type Game struct{}

func NewGame() *Game {
	ocean.Initialize()
	return &Game{}
}

func (g *Game) Update() error {
	// fmt.Printf("%v %v\n", simulationRunning)
	if simulationRunning {
		ocean.UpdateFishMovement()
		ocean.UpdateSharkMovement()
		time.Sleep(100 * time.Millisecond) // 每秒更新一次
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	drawUI(screen, ocean.Grid)

	drawButton(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
