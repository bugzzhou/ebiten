package campfirescene

import (
	"ebiten/scene/comm"
	"ebiten/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var AmountOfRecoveryOfHp = 25

func Recover(c *comm.Character) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x1, x2, y1, y2 := GetXYRangeInt(utils.CampFileFlag)
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			c.Hp = utils.Min(c.Hp+AmountOfRecoveryOfHp, c.Hplimit)
			return true
		}
	}
	return false
}
