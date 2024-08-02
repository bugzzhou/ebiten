package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 100
)

type Point struct {
	X, Y int
}

type Button struct {
	x, y          int
	width, height int
	label         string
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), color.RGBA{0, 255, 255, 255}, false)
	ebitenutil.DebugPrintAt(screen, b.label, b.x+20, b.y+10)
}

func (b *Button) IsClicked(x, y int) bool {
	return x >= b.x && x <= b.x+b.width && y >= b.y && y <= b.y+b.height
}

var (
	// 存放细胞的方格
	grid     = make([][]bool, gridSize)
	cellSize = screenWidth / gridSize

	// 按钮信息
	stepButton = &Button{
		x:      screenWidth - 300,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Step",
	}
	runButton = &Button{
		x:      screenWidth - 200,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Run",
	}
	stopButton = &Button{
		x:      screenWidth - 100,
		y:      screenHeight - 40,
		width:  80,
		height: 30,
		label:  "Stop",
	}

	// 其它中间变量
	mousePressedLastFrame bool
	changedCells          = make(map[Point]bool)
	running               bool
	ticker                *time.Ticker
)

func init() {
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}
}
