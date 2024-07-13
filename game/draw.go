package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

// 1 获得 handCard （5张牌）
// 2 根据handCard的id 从imageMap中取值
// 3 根据 position 函数获取5张牌的xy位置
// 4 循环绘制
func drawManyCards(g *Game, screen *ebiten.Image) {
	// x := 0
	// y := float64(ScreenHeight - imageHeight)
	// index := 1
	// flag := false // 表示是否有卡牌被悬停（处理两牌叠加状态）
	// x, y := ebiten.CursorPosition()

	handCards := g.HandCards
	handXY := getHandcardXYs(len(handCards))

	// expandIndex := getExpandIndex(len(handCards))

	changeStatus(g)

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

		screen.DrawImage(v, chaOpt)
	}
}

// 1、未拖拽状态 （直接通过鼠标位置的xy判断悬停）
// 2、拖拽状态（点击的一刹那会记录卡牌的index，修改全局变量index、isDrag； 通过index和isDrag 修改index张牌的状态）
//
//	松手的一刹那 修改index张牌为原位置，并且把index、isDrag修改回来
func changeStatus(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		index := getExpandIndex(len(g.HandCards))
		if index != -1 {
			g.isDragging = true
			g.draggingIndex = index
			g.expandIndex = -1
		} else {
			return
		}

	}

	//拖动过程中
	if g.isDragging {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			if isMouseOverEnemy() {
				g.DiscardCards = append(g.DiscardCards, g.HandCards[g.draggingIndex])
				g.HandCards = append(g.HandCards[:g.draggingIndex], g.HandCards[g.draggingIndex+1:]...)

				// g.showCard = false // 牌消失
			}
			g.isDragging = false
			g.draggingIndex = -1
		}
	}

	if !g.isDragging {
		// 未点击，即悬停的判断
		g.expandIndex = getExpandIndex(len(g.HandCards))
	}

	// g.testCount += 1
	// if g.testCount%10 == 0 {
	// 	fmt.Printf("%v   %v   %v  \n", g.isDragging, g.expandIndex, g.draggingIndex)
	// }

}

//无用函数

// func drawCard(g *Game, screen *ebiten.Image) {
// 	cardOpt := &ebiten.DrawImageOptions{}
// 	if g.isDragging || isMouseOverCard(g) {
// 		cardOpt.GeoM.Scale(1.2, 1.2) // 放大卡牌
// 		cardOpt.GeoM.Translate(g.cardX-imageWidth*0.1, g.cardY-imageHeight*0.1)
// 	} else {
// 		cardOpt.GeoM.Translate(g.cardX, g.cardY)
// 	}
// 	screen.DrawImage(g.card, cardOpt)
// }
// func drawRefreshButton(screen *ebiten.Image) {
// 	// 设置刷新按钮
// 	x, y, width, height := ScreenWidth-50, 0, 50, 50

// 	// 画框子
// 	boxColor := color.RGBA{0, 255, 0, 255}
// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), boxColor, false)

// 	// 设置文本
// 	text := "refresh"

// 	// 设置文本的位置
// 	textX, textY := x+10, y+20

// 	// 显示文本
// 	ebitenutil.DebugPrintAt(screen, text, textX, textY)
// }
