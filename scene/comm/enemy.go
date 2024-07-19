package comm

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Enemy struct {
	Image   *ebiten.Image
	Hp      int
	Hplimit int
	Action  []Act
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

var (
	act1 = Act{
		Id:          1,
		Name:        "添加仪式buff",
		Description: "给予自己3层仪式，每回合增加3点力量",
	}
	act2 = Act{
		Id:          2,
		Name:        "攻击",
		Description: "给予对手6+1*力量的伤害",
	}
)

var (
	buff1 = Buff{
		Id:          1,
		Name:        "仪式",
		Description: "每回合增加3点力量",

		Layers:     3,
		StartRound: 1,
		EndRound:   -1, //永远不会消失
	}
)

func GetActs(tag int) []Act {
	if tag == KakaActTag {
		return []Act{act1, act2}
	}
	return nil
}

func (enemy *Enemy) EnemyAct(round int, character *Character) {
	actIndex := getActIndex(KakaActTag, round)
	// fmt.Println(actIndex)

	if actIndex == 0 {
		enemy.Buffs = append(enemy.Buffs, buff1)
		return
	} else {
		b := enemy.Buffs[0]
		attack := 6 + b.Layers*(round-b.StartRound)
		character.Hp -= attack
	}

}

func getActIndex(tag, round int) int {
	if tag == KakaActTag {
		if round == 1 {
			return 0
		} else {
			return 1
		}
	}
	return 0
}

func GetLocalKaka() *Enemy {
	ene, _, err := ebitenutil.NewImageFromFile(Kaka) // kaka的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}
	LocalEnemy = Enemy{
		Image:   ene,
		Hp:      30,
		Hplimit: 30, //写大点方便多牌演示
		Action:  GetActs(KakaActTag),
	}
	return &LocalEnemy
}
