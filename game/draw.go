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
	screen.DrawImage(g.character.image, chaOpt)
	drawHp(screen, int(x1), int(y1), g.character.hp, g.character.hplimit)

	x2, y2 := GetXY(EnemyPos)
	eneOpt := &ebiten.DrawImageOptions{}
	eneOpt.GeoM.Translate(x2, y2)
	screen.DrawImage(g.enemy.image, eneOpt)
	drawHp(screen, int(x2), int(y2), g.enemy.hp, g.enemy.hplimit)

}

func drawHp(screen *ebiten.Image, x, y, hp, hplimit int) {
	y += imageWidth
	filledRatio := float64(hp) / float64(hplimit)
	filledLength := float32(filledRatio * float64(barLength))
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(barLength), float32(barHeight), color.RGBA{0, 0, 0, 255}, false) // 黑色
	vector.DrawFilledRect(screen, float32(x), float32(y), filledLength, float32(barHeight), color.RGBA{255, 0, 0, 255}, false)     // 红色
	text := fmt.Sprintf("%d/%d", hp, hplimit)
	ebitenutil.DebugPrintAt(screen, text, x, y+barHeight+10)
}

// 一般用于测试，显示信息
func drawText(g *Game, screen *ebiten.Image) {
	x, y, width, height := 50, 50, 400, 100
	boxColor := color.RGBA{0, 0, 255, 255} // 蓝色框子
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	// 设置文本
	str := `
	length are:  %d %d %d
	round is: %v
	energy is: %v`
	text := fmt.Sprintf(str, len(g.DrawCards), len(g.HandCards), len(g.DiscardCards), g.round, g.character.energy)
	textX, textY := x+10, y+10
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

func drawSendButton(screen *ebiten.Image) {
	// 设置发牌按钮
	x, y, width, height := 0, 0, 50, 50

	boxColor := color.RGBA{0, 255, 255, 255}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	text := "send cards"
	textX, textY := x, y
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

func endTurnButton(screen *ebiten.Image) {
	// 设置结束回合按钮
	x, y, width, height := ScreenWidth-50, 0, 50, 50

	boxColor := color.RGBA{0, 255, 255, 255}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	text := "end turn"
	textX, textY := x, y
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

// 1 获得 handCard （5张牌）
// 2 根据handCard的id 从imageMap中取值
// 3 根据 position 函数获取5张牌的xy位置
// 4 循环绘制
func drawManyCards(g *Game, screen *ebiten.Image) {
	handCards := g.HandCards
	handXY := getHandcardXYs(len(handCards))

	for i, v := range handCards {
		chaOpt := &ebiten.DrawImageOptions{}
		if i == g.expandIndex {
			chaOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
			chaOpt.GeoM.Translate(float64(handXY[i].X)-0.1*imageWidth, float64(handXY[i].Y)*0.9-0.1*imageHeight)
		} else if g.isDragging && i == g.draggingIndex {
			chaOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
			x, y := ebiten.CursorPosition()
			chaOpt.GeoM.Translate(float64(x)-imageWidth/2-0.1*imageWidth, float64(y)-imageHeight/2-0.1*imageHeight)
		} else {
			chaOpt.GeoM.Translate(float64(handXY[i].X), float64(handXY[i].Y))
		}

		screen.DrawImage(v.image, chaOpt)
	}
}

func kakaActButton(screen *ebiten.Image) {
	// 设置发牌按钮
	x, y, width, height := 100, 0, 100, 50

	boxColor := color.RGBA{0, 255, 255, 255}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	text := "kaka takes actions"
	ebitenutil.DebugPrintAt(screen, text, x, y)
}
