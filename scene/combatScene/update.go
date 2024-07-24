package combatscene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func SendCards(g *Game) {
	g.Character.Shuffle()
	g.Character.DrawCard(5)
	g.Character.Energy = 3
}

// 结束回合
func EndCards(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x1, x2, y1, y2 := GetXYRangeInt(EndButton)
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			g.Character.EndTurn()
			EnemyAct(g)
			g.RoundBegin = true
		}
	}
}

// 1、未拖拽状态 （直接通过鼠标位置的xy判断悬停）
// 2、拖拽状态（点击的一刹那会记录卡牌的index，修改全局变量index、isDrag； 通过index和isDrag 修改index张牌的状态）
// 松手的一刹那 修改index张牌为原位置，并且把index、isDrag修改回来
func ChangeStatus(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		index := getExpandIndex(len(g.Character.HandCards))
		if index != -1 {
			g.IsDragging = true
			g.DraggingIndex = index
			g.ExpandIndex = -1
		} else {
			return
		}
	}

	//拖动过程中
	if g.IsDragging {
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			if isMouseOverEnemy() && enemyIsEnough(g, g.DraggingIndex) {
				g.Character.PlayCard(g.DraggingIndex, &g.Enemy)
			}
			g.IsDragging = false
			g.DraggingIndex = -1
		}
	}
	if !g.IsDragging {
		g.ExpandIndex = getExpandIndex(len(g.Character.HandCards))
	}
}

// 判断鼠标出现在了卡牌消失的地方
func isMouseOverEnemy() bool {
	x, y := ebiten.CursorPosition()
	enemyX, enemyY := GetXY(EnemyPos)
	return x >= int(enemyX) && x <= int(enemyX)+imageWidth && y >= int(enemyY) && y <= int(enemyY)+imageHeight
}

func enemyIsEnough(g *Game, index int) bool {
	c := g.Character.HandCards[index]
	return g.Character.Energy >= c.Cost
}

// enemy行动
func EnemyAct(g *Game) {
	g.Enemy.EnemyAct(g.Round, g.Character)
	g.Round += 1
}
