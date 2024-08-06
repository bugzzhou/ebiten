package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawUI(screen *ebiten.Image, clockMap [][]int) {
	xSize := ScreenWidth / len(clockMap[0])
	ySize := ScreenHeight / len(clockMap)
	for i := range clockMap {
		for j := range clockMap[i] {
			if clockMap[i][j] == 1 {
				vector.DrawFilledRect(screen, float32(j*xSize), float32(i*ySize), float32(xSize), float32(ySize), color.White, false)
			} else {
				vector.DrawFilledRect(screen, float32(j*xSize), float32(i*ySize), float32(xSize), float32(ySize), color.Black, false)
			}
		}
	}
}
