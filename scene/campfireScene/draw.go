package campfirescene

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	CampfirePic = "./pic/kaka.jpg"
)

var campFirePic *ebiten.Image

func init() {
	var err error
	campFirePic, _, err = ebitenutil.NewImageFromFile(CampfirePic) // kaka的图片
	if err != nil {
		fmt.Printf("failed to get campfirePic, and err is: %s\n", err.Error())
	}
}

func DrawCampfire(screen *ebiten.Image) {
	screen.Fill(color.Black)
	x1, y1 := GetXY(CampFileFlag)
	chaOpt := &ebiten.DrawImageOptions{}
	chaOpt.GeoM.Translate(x1, y1)
	screen.DrawImage(campFirePic, chaOpt)
}
