package mapscene

import m "ebiten/scene/models"

const (
	StartNode    = "start"
	CombatNode   = "combat"
	CampfireNode = "fire"
	RandomNode   = "random"
	EndNode      = "end"
	BossNode     = "boss"
)

var LastChooseRoomIndex int

var LocalMapInfo *m.Graph

// var sampleMap = []m.Node{
// 	{Id: 0, NodeType: CombatNode, X: 100, Y: 250, AdjList: []int{1, 2}}, // 0
// 	{Id: 1, NodeType: CombatNode, X: 200, Y: 150, AdjList: []int{3, 4}}, // 1
// 	{Id: 2, NodeType: CombatNode, X: 200, Y: 350, AdjList: []int{5, 6}}, // 2
// 	{Id: 3, NodeType: CombatNode, X: 300, Y: 100, AdjList: []int{7}},    // 3
// 	{Id: 4, NodeType: CombatNode, X: 300, Y: 200, AdjList: []int{7}},    // 4
// 	{Id: 5, NodeType: CombatNode, X: 300, Y: 300, AdjList: []int{8}},    // 5
// 	{Id: 6, NodeType: CombatNode, X: 300, Y: 400, AdjList: []int{8}},    // 6
// 	{Id: 7, NodeType: CampfireNode, X: 400, Y: 150, AdjList: []int{9}},  // 7
// 	{Id: 8, NodeType: CampfireNode, X: 400, Y: 350, AdjList: []int{9}},  // 8
// 	{Id: 9, NodeType: EndNode, X: 500, Y: 250, AdjList: []int{}},        // 9
// }

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

//
//
//
//
//
//
//
// 仅一条路的测试路径地图
func gensimplePathMap() []m.Node {
	res := make([]m.Node, 13)

	var nodetype string
	var nextIndex int

	for i := range res {
		if i == 12 {
			nodetype = BossNode

		} else if i%4 == 3 {
			nodetype = CampfireNode
		} else {
			nodetype = CombatNode
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
