package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawUI(screen *ebiten.Image) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				vector.DrawFilledRect(screen, float32(i*cellSize), float32(j*cellSize), float32(cellSize), float32(cellSize), color.White, false)
			} else {
				vector.DrawFilledRect(screen, float32(i*cellSize), float32(j*cellSize), float32(cellSize), float32(cellSize), color.Black, false)
			}
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), color.RGBA{0, 255, 255, 255}, false)
	ebitenutil.DebugPrintAt(screen, b.label, b.x+20, b.y+10)
}
