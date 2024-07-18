package combatscene

import m "ebiten/scene/models"

const (
	KakaActTag = iota
)

var (
	act1 = m.Act{
		Id:          1,
		Name:        "添加仪式buff",
		Description: "给予自己3层仪式，每回合增加3点力量",
	}
	act2 = m.Act{
		Id:          2,
		Name:        "攻击",
		Description: "给予对手6+1*力量的伤害",
	}
)

var (
	buff1 = m.Buff{
		Id:          1,
		Name:        "仪式",
		Description: "每回合增加3点力量",

		Layers:     3,
		StartRound: 1,
		EndRound:   -1, //永远不会消失
	}
)

func GetActs(tag int) []m.Act {
	if tag == KakaActTag {
		return []m.Act{act1, act2}
	}
	return nil
}

func enemyAct(g *Game) {
	actIndex := getActIndex(KakaActTag, g.Round)
	// fmt.Println(actIndex)

	if actIndex == 0 {
		g.Enemy.Buffs = append(g.Enemy.Buffs, buff1)
		return
	} else {
		b := g.Enemy.Buffs[0]
		attack := 6 + b.Layers*(g.Round-b.StartRound)
		g.Character.Hp -= attack
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
