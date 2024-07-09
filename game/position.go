package game

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

	TestFlag = 999
)

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
	case TestFlag:
		return 0, 0, 0, 0
	}

	return 0, 0, 0, 0
}
