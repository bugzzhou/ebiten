package comm

import (
	"ebiten/utils"
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)

var (
	LocalCharacter = Character{}
	LocalEnemy     = Enemy{}
)

func init() {
	cha, _, err := ebitenutil.NewImageFromFile(utils.Lieren) // 猎人的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}

	LocalCharacter = Character{
		Image:   cha,
		Hp:      99,
		Hplimit: 99,
		Energy:  3,
	}

	initCards := GetCards()
	LocalCharacter = Character{
		Image:    cha,
		Hp:       99,
		Hplimit:  99,
		Energy:   99,
		Cards:    initCards,
		DrawDeck: initCards,
	}
}

func GetLocalCharacter() *Character {
	LocalCharacter.Energy = 3
	LocalCharacter.DrawDeck = LocalCharacter.Cards
	LocalCharacter.HandCards = nil
	LocalCharacter.DiscardDeck = nil
	return &LocalCharacter
}

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap = map[int]*ebiten.Image{}

func init() {
	files, ids, err := utils.ListDir(utils.CardDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}

	for i := range files {
		idInt, err := strconv.Atoi(ids[i])
		if err != nil {
			fmt.Printf("failed to get convert, and err is: %s\n", err.Error())
			continue
		}

		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		cardImageMap[idInt] = tmpImage
	}
}
