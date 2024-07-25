package combatscene

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawCharAEnemy(g *Game, screen *ebiten.Image) {
	screen.Fill(color.Black)
	// Draw character and enemy first
	x1, y1 := GetXY(CharacterPos)
	chaOpt := &ebiten.DrawImageOptions{}
	chaOpt.GeoM.Translate(x1, y1)
	screen.DrawImage(g.Character.Image, chaOpt)
	drawHpAndShield(screen, int(x1), int(y1), g.Character.Hp, g.Character.Hplimit, 0)

	x2, y2 := GetXY(EnemyPos)
	eneOpt := &ebiten.DrawImageOptions{}
	eneOpt.GeoM.Translate(x2, y2)
	screen.DrawImage(g.Enemy.Image, eneOpt)
	drawHpAndShield(screen, int(x2), int(y2), g.Enemy.Hp, g.Enemy.Hplimit, g.Enemy.Shield)

}

// 一般用于测试，显示信息
func DrawText(g *Game, screen *ebiten.Image) {
	x, y, width, height := 50, 50, 400, 100
	boxColor := color.RGBA{0, 0, 255, 255} // 蓝色框子
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

	// 设置文本
	str := `
	length are:  %d %d %d
	round is: %v
	energy is: %v`
	text := fmt.Sprintf(str, len(g.Character.DrawDeck), len(g.Character.HandCards), len(g.Character.DiscardDeck), g.Round, g.Character.Energy)
	textX, textY := x+10, y+10
	ebitenutil.DebugPrintAt(screen, text, textX, textY)
}

func drawHpAndShield(screen *ebiten.Image, x, y, hp, hplimit, shield int) {
	y += imageWidth

	// 绘制血条背景
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(barLength), float32(barHeight), color.RGBA{0, 0, 0, 255}, false) // 黑色

	// 计算血条的当前填充长度
	filledRatio := float64(hp) / float64(hplimit)
	filledLength := float32(filledRatio * float64(barLength))
	// 绘制当前血量的红色部分
	vector.DrawFilledRect(screen, float32(x), float32(y), filledLength, float32(barHeight), color.RGBA{255, 0, 0, 255}, false) // 红色

	// 显示血条的数值
	text := fmt.Sprintf("%d/%d", hp, hplimit)
	ebitenutil.DebugPrintAt(screen, text, x, y+barHeight+10)

	// 绘制护盾的蓝色方框和文本
	if shield > 0 {
		shieldColor := color.RGBA{0, 0, 255, 255} // 蓝色
		shieldSize := barHeight                   // 假设护盾方框的大小与血条高度相同
		shieldX := x + (barLength-shieldSize)/2   // 护盾方框的x位置在血条的正下方中间
		shieldY := y + barHeight                  // 护盾方框的y位置在血条的正下方

		// 绘制护盾方框
		vector.DrawFilledRect(screen, float32(shieldX), float32(shieldY), float32(shieldSize), float32(shieldSize), shieldColor, false)

		// 显示护盾的数值
		shieldText := fmt.Sprintf("%d", shield)
		// shieldTextWidth := ebiten.TextWidth(ebiten.DefaultFont, shieldText)
		shieldTextX := shieldX
		shieldTextY := shieldY + shieldSize/2 + 1 // 稍微偏移一点以适应文本
		ebitenutil.DebugPrintAt(screen, shieldText, shieldTextX, shieldTextY)
	}
}

func EndTurnButton(screen *ebiten.Image) {
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
func DrawManyCards(g *Game, screen *ebiten.Image) {
	handCards := g.Character.HandCards
	handXY := getHandcardXYs(len(handCards))

	for i, v := range handCards {
		chaOpt := &ebiten.DrawImageOptions{}
		if i == g.ExpandIndex {
			chaOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
			chaOpt.GeoM.Translate(float64(handXY[i].X)-0.1*imageWidth, float64(handXY[i].Y)*0.9-0.1*imageHeight)
		} else if g.IsDragging && i == g.DraggingIndex {
			chaOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
			x, y := ebiten.CursorPosition()
			chaOpt.GeoM.Translate(float64(x)-imageWidth/2-0.1*imageWidth, float64(y)-imageHeight/2-0.1*imageHeight)
		} else {
			chaOpt.GeoM.Translate(float64(handXY[i].X), float64(handXY[i].Y))
		}
		screen.DrawImage(v.Image, chaOpt)
	}
}
