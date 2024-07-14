package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 点击发牌按钮，抽排
func sendCards(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x1, x2, y1, y2 := GetXYRangeInt(SendButton)
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			g.Shuffle()
			g.DrawCard(5)
		}
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

func changeHpRand(g *Game) {
	if g.characterHp <= 0 {
		g.characterHp = g.characterHpLimit
	} else {
		if R.Intn(3) == 0 {
			g.characterHp -= 1
		}
	}

	if g.enemyHp <= 0 {
		g.enemyHp = g.enemyHpLimit
	} else {
		if R.Intn(9) == 0 {
			g.enemyHp -= 1
		}
	}

	if g.testHp <= 0 {
		g.testHp = 100
	} else {
		g.testHp -= 1
	}
}

// 核心拖拽函数
/*
1、鼠标进来，判断鼠标位置，判断鼠标动作，判断选择了哪张图片，获取图片的index ※
2、
*/
// func checkCardDrag(g *Game) {
// 	length := len(g.HandCards)

// 	index := getExpandIndex(length)
// 	if index == -1 {
// 		return
// 	}
// 	//开始拖动图片的判断
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		g.isDragging = true
// 	}

// 	//拖动过程中
// 	if g.isDragging {
// 		x, y := ebiten.CursorPosition()
// 		g.cardX = float64(x) - imageWidth/2
// 		g.cardY = float64(y) - imageHeight/2

// 		//松开鼠标的那一刹那
// 		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
// 			g.isDragging = false
// 			if isMouseOverEnemy() {
// 				g.showCard = false // 牌消失
// 			} else {
// 				g.cardX = g.cardOriginalX
// 				g.cardY = g.cardOriginalY
// 			}
// 		}
// 	}
// }

// 判断鼠标出现在了卡牌消失的地方
func isMouseOverEnemy() bool {
	x, y := ebiten.CursorPosition()
	enemyX, enemyY := GetXY(EnemyPos)
	return x >= int(enemyX) && x <= int(enemyX)+imageWidth && y >= int(enemyY) && y <= int(enemyY)+imageHeight
}

// // 点击“发牌”的按钮
// func checkRefreshButtonClick(g *Game) {
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		x, y := ebiten.CursorPosition()
// 		if x >= ScreenWidth-50 && y <= 50 {
// 			g.RefreshCard()
// 		}
// 	}
// }
