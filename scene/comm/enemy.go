package comm

import (
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

const (
	KakaId = iota

	TestEnemyId = 9999
)

// 修改成三个map
// map[int]map[int]Act enemy下的所有动作
// map[enemyName][int]int ， enemy第n个回合，enemy采取的动作是所有动作中第m个
// map[int]Buff， key是唯一标识，value是单个buff的具体信息
var actsMap map[int]map[int]Act
var actINdexMap map[int][]int
var buffMap map[int]Buff

func init() {
	actsMap = map[int]map[int]Act{
		KakaActTag: {
			1: Act{Id: 1, Name: "添加仪式buff", Description: "给予自己3层仪式，每回合增加3点力量"},
			2: Act{Id: 2, Name: "攻击", Description: "给予对手6+1*力量的伤害"},
		},

		TestEnemyActTag: {
			1: Act{Id: 1, Name: "攻击", Description: "给英雄单位造成6点伤害"},
			2: Act{Id: 2, Name: "防御", Description: "给自己提供5点护盾"},
		},
	}

	buffMap = map[int]Buff{
		1: Buff{
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

	if id == KakaId {
		if actIndex == 0 {
			enemy.Buffs = append(enemy.Buffs, buffMap[1])
			return
		} else {
			b := enemy.Buffs[0]
			attack := 6 + b.Layers*(round-b.StartRound)
			character.Hp -= attack
		}
	}

	if id == TestEnemyActTag {
		if actIndex == 0 {
			character.Hp -= 6
		} else {
			enemy.Shield = 0 //不一定是无用代码，先留着，保证思路
			enemy.Shield = 5
		}
	}

}

// kaka！
func GetLocalKaka() *Enemy {
	kakaImage, _, err := ebitenutil.NewImageFromFile(Kaka) // kaka的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Id:      KakaId,
		Image:   kakaImage,
		Hp:      30,
		Hplimit: 30, //写大点方便多牌演示
		Action:  getActs(KakaActTag),
	}
	return &LocalEnemy
}

func GetTestEnemy() *Enemy {
	testEnemyImage, _, err := ebitenutil.NewImageFromFile(TestEnemy) // 测试敌人的图片替换
	if err != nil {
		fmt.Printf("failed to get testEnemy pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Id:      TestEnemyId,
		Image:   testEnemyImage,
		Hp:      50,
		Hplimit: 50,
		Action:  getActs(TestEnemyActTag),
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
// 		Action:  getActs(KakaActTag),
// 	}
// 	return &LocalEnemy
// }

func getActs(tag int) map[int]Act {
	return actsMap[tag]
}

func getActIndex(tag, round int) int {
	if tag == KakaActTag {
		if round == 1 {
			return 0
		} else {
			return 1
		}
	}

	if tag == TestEnemyActTag {
		return round % 2
	}
	return 0
}
