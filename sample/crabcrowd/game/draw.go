package game

import (
	"crabcrowd/core"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	colorBackground = color.RGBA{24, 26, 34, 255}
	colorWall       = color.RGBA{70, 62, 58, 255}
	colorFloor      = color.RGBA{214, 214, 220, 255}
	colorTargetRing = color.RGBA{240, 190, 70, 255}
	colorCrabBody   = color.RGBA{224, 108, 58, 255}
	colorCrabOnGoal = color.RGBA{92, 190, 120, 255}
	colorCrabEye    = color.RGBA{255, 255, 255, 255}
	colorCrabPupil  = color.RGBA{30, 30, 30, 255}
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colorBackground)

	lv := g.state.Level
	offsetX, offsetY := g.gridOrigin(lv)

	for y := 0; y < lv.Height(); y++ {
		for x := 0; x < lv.Width(); x++ {
			p := core.Point{X: x, Y: y}
			cx := offsetX + float32(x*CellSize)
			cy := offsetY + float32(y*CellSize)
			switch lv.CellAt(p) {
			case core.Wall:
				vector.DrawFilledRect(screen, cx+1, cy+1, CellSize-2, CellSize-2, colorWall, false)
			case core.Floor:
				vector.DrawFilledRect(screen, cx+1, cy+1, CellSize-2, CellSize-2, colorFloor, false)
			case core.Target:
				vector.DrawFilledRect(screen, cx+1, cy+1, CellSize-2, CellSize-2, colorFloor, false)
				vector.DrawFilledCircle(screen, cx+CellSize/2, cy+CellSize/2, CellSize/5, colorTargetRing, true)
			}
		}
	}

	for _, p := range g.state.Crabs {
		cx := offsetX + float32(p.X*CellSize) + CellSize/2
		cy := offsetY + float32(p.Y*CellSize) + CellSize/2
		body := colorCrabBody
		if lv.CellAt(p) == core.Target {
			body = colorCrabOnGoal
		}
		radius := float32(CellSize) * 0.36
		vector.DrawFilledCircle(screen, cx, cy, radius, body, true)

		eyeOffsetX := radius * 0.45
		eyeOffsetY := radius * 0.5
		eyeR := radius * 0.22
		vector.DrawFilledCircle(screen, cx-eyeOffsetX, cy-eyeOffsetY, eyeR, colorCrabEye, true)
		vector.DrawFilledCircle(screen, cx+eyeOffsetX, cy-eyeOffsetY, eyeR, colorCrabEye, true)
		vector.DrawFilledCircle(screen, cx-eyeOffsetX, cy-eyeOffsetY, eyeR*0.5, colorCrabPupil, true)
		vector.DrawFilledCircle(screen, cx+eyeOffsetX, cy-eyeOffsetY, eyeR*0.5, colorCrabPupil, true)
	}

	g.drawHUD(screen)
}

// gridOrigin 计算地图在屏幕上的绘制起点，使其在 HUD 区域之下居中显示。
func (g *Game) gridOrigin(lv *core.Level) (float32, float32) {
	gridW := float32(lv.Width() * CellSize)
	gridH := float32(lv.Height() * CellSize)
	x := (float32(ScreenW) - gridW) / 2
	y := hudTop + (float32(ScreenH-hudTop-hudBottom)-gridH)/2
	return x, y
}

func (g *Game) drawHUD(screen *ebiten.Image) {
	lv := g.state.Level

	title := fmt.Sprintf("%s   (%d/%d)", lv.Name, g.levelIdx+1, len(core.Levels))
	ebitenutil.DebugPrintAt(screen, title, 20, 16)

	status := fmt.Sprintf("步数: %d   可撤销: %d", g.state.Steps, g.state.UndoDepth())
	ebitenutil.DebugPrintAt(screen, status, 20, 34)

	help := "移动 WASD | 撤销 U | 重置 R | 上一关 P | 下一关 N"
	ebitenutil.DebugPrintAt(screen, help, 20, ScreenH-40)

	if g.state.Won() {
		ebitenutil.DebugPrintAt(screen, "恭喜过关！按 N 进入下一关，或按 R 重玩本关", 20, ScreenH-20)
	}
}
