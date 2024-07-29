package comm

import (
	"ebiten/utils"
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	initCardMap()
	initLocalCharacter()

	initActMap()
	initBuffMap()
}

func initCardMap() {
	// get cardImageMap
	cardImageMap = make(map[int]*ebiten.Image)
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

	// add cardInfo into cardImageMap
	allCardsMap = make(map[int]CardInfo)
	allCardBaseinfo, err := readCSVFile(utils.CardInfoPath)
	if err != nil {
		fmt.Printf("failed to read csv, and err is: %s\n", err.Error())
	}
	for _, v := range allCardBaseinfo {
		v.Image = cardImageMap[v.Id]
		allCardsMap[v.Id] = v
	}
	// 此时的allCardsMap 已经包含了所有卡牌信息、图片，可以直接用
}

func initLocalCharacter() {
	cha, _, err := ebitenutil.NewImageFromFile(utils.Lieren) // 猎人的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}

	initCards := getInitCards()
	LocalCharacter = Character{
		Image:    cha,
		Hp:       99,
		Hplimit:  99,
		Energy:   3,
		Cards:    initCards,
		DrawDeck: initCards,
	}
}

func getInitCards() []CardInfo {
	c1 := allCardsMap[1]
	c2 := allCardsMap[2]
	c3 := allCardsMap[3]
	c4 := allCardsMap[4]
	c5 := allCardsMap[5]

	return []CardInfo{
		c1, c1, c1, c1, c5,
		c2, c2, c4, c3, c3,
	}
}

/*
	所有与enemy相关的初始化数据

TODO jszhou 把act、buff格式化。 类似卡牌一样的实现
Tips: 配置化是为了实现具体内容的时候可以通过通用函数直接实现，而不是为了配置化而配置化，请注意！
*/
func initActMap() {
	actsMap = map[int]map[int]Act{
		utils.KakaId: {
			1: Act{Id: 1, Name: "添加仪式buff", Description: "给予自己3层仪式，每回合增加3点力量"},
			2: Act{Id: 2, Name: "攻击", Description: "给予对手6+1*力量的伤害"},
		},

		utils.Boss1Id: {
			1: Act{Id: 1, Name: "重击", Description: "20攻击"},
			2: Act{Id: 2, Name: "多段", Description: "3*6攻击"},
			3: Act{Id: 3, Name: "防御", Description: "20护盾"},
		},

		utils.TestEnemyId: {
			1: Act{Id: 1, Name: "攻击", Description: "给英雄单位造成6点伤害"},
			2: Act{Id: 2, Name: "防御", Description: "给自己提供5点护盾"},
		},
	}
}

func initBuffMap() {
	buffMap = map[int]Buff{
		1: {
			Id:          1,
			Name:        "仪式",
			Description: "每回合增加3点力量",

			Layers:     3,
			StartRound: 1,
			EndRound:   -1, //永远不会消失
		},
	}
}
