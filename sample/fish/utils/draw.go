package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawUI(screen *ebiten.Image, clockMap [100][100]int) {
	xSize := ScreenWidth / len(clockMap[0])
	ySize := ScreenHeight / len(clockMap)
	for i := range clockMap {
		for j := range clockMap[i] {
			if clockMap[i][j] == 1 || clockMap[i][j] == 2 {
				vector.DrawFilledRect(screen, float32(j*xSize), float32(i*ySize), float32(xSize), float32(ySize), color.White, false)
			} else {
				vector.DrawFilledRect(screen, float32(j*xSize), float32(i*ySize), float32(xSize), float32(ySize), color.Black, false)
			}
		}
	}
}

func drawButton(screen *ebiten.Image) {
	// 绘制按钮
	startButton := ebiten.NewImage(100, 30)
	stopButton := ebiten.NewImage(100, 30)
	startButton.Fill(color.RGBA{0, 255, 0, 255}) // 绿色按钮
	stopButton.Fill(color.RGBA{255, 0, 0, 255})  // 红色按钮

	ebitenutil.DebugPrintAt(startButton, "开始", 35, 10)
	ebitenutil.DebugPrintAt(stopButton, "停止", 35, 10)

	// 计算按钮位置
	startButtonOpts := &ebiten.DrawImageOptions{}
	stopButtonOpts := &ebiten.DrawImageOptions{}
	startButtonOpts.GeoM.Translate(50, ScreenHeight-50)
	stopButtonOpts.GeoM.Translate(200, ScreenHeight-50)

	screen.DrawImage(startButton, startButtonOpts)
	screen.DrawImage(stopButton, stopButtonOpts)

	// 检测鼠标点击
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= 50 && x <= 150 && y >= ScreenHeight-50 && y <= ScreenHeight-20 {
			// 点击了“开始”按钮
			simulationRunning = true
		}
		if x >= 200 && x <= 300 && y >= ScreenHeight-50 && y <= ScreenHeight-20 {
			// 点击了“停止”按钮
			simulationRunning = false
		}
	}
}
