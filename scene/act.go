package scene

const (
	kakaActTag = iota
)

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

func getActs(tag int) []Act {
	if tag == kakaActTag {
		return []Act{act1, act2}
	}
	return nil
}

func enemyAct(g *Game) {
	actIndex := getActIndex(kakaActTag, g.round)
	// fmt.Println(actIndex)

	if actIndex == 0 {
		g.enemy.buffs = append(g.enemy.buffs, buff1)
		return
	} else {
		b := g.enemy.buffs[0]
		attack := 6 + b.Layers*(g.round-b.StartRound)
		g.character.hp -= attack
	}

}

func getActIndex(tag, round int) int {
	if tag == kakaActTag {
		if round == 1 {
			return 0
		} else {
			return 1
		}
	}
	return 0
}
