package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawGrid(screen *ebiten.Image) {
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

func drawButtons(screen *ebiten.Image) {
	stepButton.Draw(screen)
	runButton.Draw(screen)
	stopButton.Draw(screen)
}
