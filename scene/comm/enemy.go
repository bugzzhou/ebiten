package comm

import (
	"ebiten/utils"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	Id      int
	Image   *ebiten.Image
	Hp      int
	Hplimit int
	Shield  int
	Action  map[int]Act
	Buffs   []Buff
}

type Act struct {
	Id          int
	Name        string
	Description string
}

type Buff struct {
	Id          int
	Name        string
	Description string

	Layers     int
	StartRound int
	EndRound   int
}

var IsBossRoom bool

// 修改成三个map
// map[int]map[int]Act enemy下的所有动作
// map[enemyName][int]int ， enemy第n个回合，enemy采取的动作是所有动作中第m个
// map[int]Buff， key是唯一标识，value是单个buff的具体信息
var actsMap map[int]map[int]Act

// var actINdexMap map[int][][]int
var buffMap map[int]Buff

func init() {
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

// 主要函数
func (enemy *Enemy) EnemyAct(round int, character *Character) {
	id := enemy.Id

	actIndex := getActIndex(id, round)

	if id == utils.KakaId {
		if actIndex == 0 {
			enemy.Buffs = append(enemy.Buffs, buffMap[1])
			return
		} else {
			b := enemy.Buffs[0]
			attack := 6 + b.Layers*(round-b.StartRound)
			character.Hp -= attack
		}
	}

	if id == utils.TestEnemyId {
		if actIndex == 0 {
			character.Hp -= 6
		} else {
			enemy.Shield = 0 //不一定是无用代码，先留着，保证思路
			enemy.Shield = 5
		}
	}

}

func GetEnemy() *Enemy {
	randn := utils.R.Intn(2)
	var enemy *Enemy

	if IsBossRoom {
		return GetBoss1()
	}

	switch randn {
	case 0:
		enemy = GetLocalKaka()
	case 1:
		enemy = GetTestEnemy()
	}

	return enemy

}

// kaka！
func GetLocalKaka() *Enemy {
	kakaImage, _, err := ebitenutil.NewImageFromFile(utils.Kaka) // kaka的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Id:      utils.KakaId,
		Image:   kakaImage,
		Hp:      30,
		Hplimit: 30, //写大点方便多牌演示
		Action:  getActs(utils.KakaId),
	}
	return &LocalEnemy
}

func GetTestEnemy() *Enemy {
	testEnemyImage, _, err := ebitenutil.NewImageFromFile(utils.TestEnemy) // 测试敌人的图片替换
	if err != nil {
		fmt.Printf("failed to get testEnemy pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Id:      utils.TestEnemyId,
		Image:   testEnemyImage,
		Hp:      50,
		Hplimit: 50,
		Action:  getActs(utils.TestEnemyId),
	}
	return &LocalEnemy
}

func GetBoss1() *Enemy {
	bossEnemyImage, _, err := ebitenutil.NewImageFromFile(utils.Boss1)
	if err != nil {
		fmt.Printf("failed to get testEnemy pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Id:      utils.Boss1Id,
		Image:   bossEnemyImage,
		Hp:      200,
		Hplimit: 200,
		Action:  getActs(utils.Boss1Id),
	}
	return &LocalEnemy
}

// // 面团
// func GetDough() *Enemy {
// 	roughImage, _, err := ebitenutil.NewImageFromFile(Rough) // TODO bugzzhou 面团的图片替换
// 	if err != nil {
// 		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
// 	}
// 	LocalEnemy = Enemy{
// 		Image:   roughImage,
// 		Hp:      50,
// 		Hplimit: 50,
// 		Action:  getActs(utils.KakaId),
// 	}
// 	return &LocalEnemy
// }

func getActs(tag int) map[int]Act {
	return actsMap[tag]
}

func getActIndex(tag, round int) int {
	if tag == utils.KakaId {
		if round == 1 {
			return 0
		} else {
			return 1
		}
	}

	if tag == utils.TestEnemyId {
		return round % 2
	}

	// 重击、多段、防御 顺序执行
	if tag == utils.Boss1Id {
		return round % 3
	}

	return 0
}
