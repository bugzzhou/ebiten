package mapscene

import (
	m "ebiten/scene/models"
	"ebiten/utils"
)

var LastChooseRoomIndex int

var LocalMapInfo *m.Graph

func init() {
	LastChooseRoomIndex = -1

	LocalMapInfo = &m.Graph{
		Nodes: gensimplePathMap(),
	}
}

func CreateSampleGraph() *m.Graph {
	return LocalMapInfo
}

func IsPointInCircle(px, py, cx, cy, radius int) bool {
	dx := px - cx
	dy := py - cy
	return dx*dx+dy*dy <= radius*radius
}

// 仅一条路的测试路径地图
func gensimplePathMap() []m.Node {
	res := make([]m.Node, 13)

	var nodetype string
	var nextIndex int

	for i := range res {
		if i == 12 {
			nodetype = utils.BossNode

		} else if i%4 == 3 {
			nodetype = utils.CampfireNode
		} else {
			nodetype = utils.CombatNode
		}

		if i == 12 {
			nextIndex = -1
		} else {
			nextIndex = i + 1
		}

		tmpNode := m.Node{
			Id:       i,
			NodeType: nodetype,
			X:        100 + i*100,
			Y:        250,
			AdjList:  []int{nextIndex},
		}
		res[i] = tmpNode
	}
	return res
}
