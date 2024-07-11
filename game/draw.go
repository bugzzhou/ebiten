package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawCharAEnemy(g *Game, screen *ebiten.Image) {
	screen.Fill(color.Black)
	// Draw character and enemy first
	x1, y1 := GetXY(CharacterPos)
	chaOpt := &ebiten.DrawImageOptions{}
	chaOpt.GeoM.Translate(x1, y1)
	screen.DrawImage(g.character, chaOpt)

	x2, y2 := GetXY(EnemyPos)
	eneOpt := &ebiten.DrawImageOptions{}
	eneOpt.GeoM.Translate(x2, y2)
	screen.DrawImage(g.enemy, eneOpt)
}

func drawCard(g *Game, screen *ebiten.Image) {
	cardOpt := &ebiten.DrawImageOptions{}
	if g.isDragging || isMouseOverCard(g) {
		cardOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
		cardOpt.GeoM.Translate(g.cardX-imageWidth*0.1, g.cardY-imageHeight*0.1)
	} else {
		cardOpt.GeoM.Translate(g.cardX, g.cardY)
	}
	screen.DrawImage(g.card, cardOpt)
}

// 一般用于测试，显示信息
func drawText(g *Game, screen *ebiten.Image) {
	// 设置框子的位置和大小
	x, y, width, height := 50, 50, 400, 100

	// 画框子
	boxColor := color.RGBA{0, 0, 255, 255} // 蓝色框子
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	// 设置文本
	text := fmt.Sprintf("length are:  %d %d %d", len(g.DrawCards), len(g.HandCards), len(g.DiscardCards))

	// 设置文本的位置
	textX, textY := x+10, y+10

	// 显示文本
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

// 一般用于测试，显示信息
func drawSendButton(screen *ebiten.Image) {
	// 设置发牌按钮
	x, y, width, height := 0, 0, 50, 50

	// 画框子
	boxColor := color.RGBA{0, 255, 255, 255}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	// 设置文本
	text := "send cards"

	// 设置文本的位置
	textX, textY := x, y

	// 显示文本
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

func drawRefreshButton(screen *ebiten.Image) {
	// 设置刷新按钮
	x, y, width, height := ScreenWidth-50, 0, 50, 50

	// 画框子
	boxColor := color.RGBA{0, 255, 0, 255}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	// 设置文本
	text := "refresh"

	// 设置文本的位置
	textX, textY := x+10, y+20

	// 显示文本
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}
