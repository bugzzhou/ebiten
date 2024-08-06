package campfirescene

import (
	"ebiten/utils"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var campFirePic *ebiten.Image

func init() {
	var err error
	campFirePic, _, err = ebitenutil.NewImageFromFile(utils.CampfirePic) // 火堆里的图片-现在只有睡觉
	if err != nil {
		fmt.Printf("failed to get campfirePic, and err is: %s\n", err.Error())
	}
}

func DrawCampfire(screen *ebiten.Image) {
	screen.Fill(color.Black)
	x1, y1 := GetXY(utils.CampFileFlag)
	chaOpt := &ebiten.DrawImageOptions{}
	chaOpt.GeoM.Translate(x1, y1)
	screen.DrawImage(campFirePic, chaOpt)
}
