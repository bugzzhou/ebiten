package comm

import (
	"ebiten/scene/models"
	"fmt"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)

var (
	Lieren = "./pic/lieren.jpg"
	Kaka   = "./pic/kaka.jpg"
)

func getGameInfo() {
	cha, _, err := ebitenutil.NewImageFromFile(Lieren) // 猎人的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error)
	}
	ene, _, err := ebitenutil.NewImageFromFile(Kaka) // kaka的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error)
	}

	Character := models.Character{
		Image:   cha,
		Hp:      99,
		Hplimit: 99,
		Energy:  3,
	}

	Enemy := models.Enemy{
		Image:   ene,
		Hp:      30,
		Hplimit: 30, //写大点方便多牌演示
		Action:  sceneCom.GetActs(sceneCom.KakaActTag),
	}

	gametmp := &sceneCom.Game{
		Round: 1,

		Cards:     allCards,
		DrawCards: allCards,

		DraggingIndex: -1,
		ExpandIndex:   -1,
		IsDragging:    false,
	}
}
