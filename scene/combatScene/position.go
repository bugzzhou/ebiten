package combatscene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	centerX = ScreenWidth / 2
	centerY = ScreenHeight / 2

	roleWidth  = 150
	roleHeight = 150
)

const (
	CardPos = iota
	CharacterPos
	EnemyPos
	SendButton
	EndButton

	KakaActButtonFlag

	TestFlag = 999
)

const interval = 120 //两张手牌之间的间隔

var (
	barLength = 150
	barHeight = 10
)

type HandcardXY struct {
	X int
	Y int
}

// 用于获取图片的中心位置
// 供ebiten.DrawImageOptions.GeoM.Translate(x, y)使用
func GetXY(flag int) (x, y float64) {
	switch flag {
	case CardPos:
		return 0, 0
	case CharacterPos:
		return 0, centerY - roleHeight/2
	case EnemyPos:
		return ScreenWidth - roleWidth, centerY - roleHeight/2
	case SendButton:
		return 25, 25
	case EndButton:
		return ScreenWidth - 25, 25
	case KakaActButtonFlag:
		return 150, 25
	case TestFlag:
		return 0, 0
	}

	return 0, 0

}

func GetXYRange(flag int) (x1, x2, y1, y2 float64) {
	switch flag {
	case CardPos:
		return 0, 0, 0, 0
	case CharacterPos:
		return 0, roleWidth, centerY - roleHeight/2, centerY + roleHeight/2
	case EnemyPos:
		return ScreenWidth - roleWidth, ScreenWidth, centerY - roleHeight/2, centerY + roleHeight/2
	case SendButton:
		return 0, 50, 0, 50
	case EndButton:
		return ScreenWidth - 50, ScreenWidth, 0, 50
	case KakaActButtonFlag:
		return 100, 200, 0, 50
	case TestFlag:
		return 0, 0, 0, 0
	}

	return 0, 0, 0, 0
}

func GetXYRangeInt(flag int) (x1, x2, y1, y2 int) {
	switch flag {
	case CardPos:
		return 0, 0, 0, 0
	case CharacterPos:
		return 0, roleWidth, centerY - roleHeight/2, centerY + roleHeight/2
	case EnemyPos:
		return ScreenWidth - roleWidth, ScreenWidth, centerY - roleHeight/2, centerY + roleHeight/2
	case SendButton:
		return 0, 50, 0, 50
	case EndButton:
		return ScreenWidth - 50, ScreenWidth, 0, 50
	case KakaActButtonFlag:
		return 100, 200, 0, 50
	case TestFlag:
		return 0, 0, 0, 0
	}

	return 0, 0, 0, 0
}

func getHandcardXYs(count int) []HandcardXY {
	res := []HandcardXY{}

	tmpX := make([]int, count)

	if count > 10 {
		return res
	}

	if count <= 0 {
		return res
	} else if count%2 == 0 {
		//偶数牌平均分
		start := centerX - interval*(count/2) + interval/2
		for i := range tmpX {
			tmpX[i] = start + interval*i - imageWidth/2
		}
	} else {
		//基数牌平均分
		start := centerX - (count-1)/2*interval
		for i := range tmpX {
			// 绘制点是图片左上角，而不是正中心，所以x需要往左移动图片宽度的一半
			tmpX[i] = start + interval*i - imageWidth/2
		}
	}

	for _, v := range tmpX {
		tmpXY := HandcardXY{
			X: v,
			Y: ScreenHeight - imageHeight,
		}
		res = append(res, tmpXY)
	}

	return res
}

func getExpandIndex(count int) int {
	xys := getHandcardXYs(count)

	x, y := ebiten.CursorPosition()

	if len(xys) <= 0 {
		return -1
	}

	if y < ScreenHeight-imageHeight || y > ScreenHeight || x < xys[0].X || x > xys[len(xys)-1].X+imageWidth {
		return -1
	}

	return (x - xys[0].X) / interval

}
