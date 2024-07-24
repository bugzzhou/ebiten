package comm

import (
	"ebiten/utils"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)

const (
	KakaActTag = iota

	Boss1ActTag = 50

	TestEnemyActTag = 9999
)

var (
	Lieren = "./pic/lieren.jpg"
	Kaka   = "./pic/kaka.jpg"
	Rough  = "./pic/rough.jpg"

	TestEnemy = "./pic/test_enemy.jpg"
	Boss1     = "./pic/test_enemy.jpg" //TODO bugzzhou 替换boss图片
)

var (
	LocalCharacter = Character{}
	LocalEnemy     = Enemy{}
)

// 唯一标识一个卡牌
type CardInfo struct {
	Id int

	Attack     int
	Shield     int
	SelfAttack int
	Cost       int

	Image *ebiten.Image
}

func init() {
	cha, _, err := ebitenutil.NewImageFromFile(Lieren) // 猎人的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}

	LocalCharacter = Character{
		Image:   cha,
		Hp:      99,
		Hplimit: 99,
		Energy:  3,
	}

	allCards := GetCards()
	LocalCharacter = Character{
		Image:    cha,
		Hp:       99,
		Hplimit:  99,
		Energy:   99,
		Cards:    allCards,
		DrawDeck: allCards,
	}
}

func GetLocalCharacter() *Character {
	LocalCharacter.Energy = 3
	LocalCharacter.DrawDeck = GetCards()
	LocalCharacter.HandCards = nil
	LocalCharacter.DiscardDeck = nil
	return &LocalCharacter
}

// 无用函数

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap = map[string]*ebiten.Image{}

func init() {
	files, ids, err := utils.ListDir(cardDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}

	for i := range files {
		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		cardImageMap[ids[i]] = tmpImage
	}
}
