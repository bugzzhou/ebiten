// package scene

// import (
// 	cons "ebiten/scene/const"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

// // 整个游戏的入口，用于选择地图

// type MapScene struct {
// 	manager *SceneManager
// }

// func NewMapScene(manager *SceneManager) *MapScene {
// 	return &MapScene{manager: manager}
// }

// func (s *MapScene) Update() error {
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		s.manager.SetScene(NewCombatScene(s.manager))
// 	}
// 	return nil
// }

// func (s *MapScene) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Scene map - there is a map to choose in it~~~")
// }

// func (s *MapScene) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return cons.ScreenWidth, cons.ScreenHeight
// }

package scene

import (
	"fmt"
	"image/color"

	cons "ebiten/scene/const"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Node struct {
	X, Y    int
	AdjList []int
}

type Graph struct {
	Nodes []Node
}

// 整个游戏的入口，用于选择地图

type MapScene struct {
	manager *SceneManager
	graph   *Graph
}

func NewMapScene(manager *SceneManager) *MapScene {
	return &MapScene{
		manager: manager,
		graph:   createSampleGraph(),
	}
}

func createSampleGraph() *Graph {
	return &Graph{
		Nodes: []Node{
			{X: 100, Y: 100, AdjList: []int{1, 2}}, // 0
			{X: 200, Y: 100, AdjList: []int{3, 4}}, // 1
			{X: 200, Y: 200, AdjList: []int{4}},    // 2
			{X: 300, Y: 100, AdjList: []int{5}},    // 3
			{X: 300, Y: 200, AdjList: []int{5}},    // 4
			{X: 400, Y: 150, AdjList: []int{}},     // 5
		},
	}
}

func (s *MapScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, node := range s.graph.Nodes {
			if isPointInCircle(x, y, node.X, node.Y, 30) {
				fmt.Printf("Clicked node %d\n", i)
				break
			}
		}
	}
	return nil
}

func (s *MapScene) Draw(screen *ebiten.Image) {
	// Draw nodes
	for i, node := range s.graph.Nodes {
		// ebitenutil.DrawRect(screen, float64(node.X)-5, float64(node.Y)-5, 10, 10, color.RGBA{0, 0, 255, 255})
		vector.DrawFilledCircle(screen, float32(node.X)-5, float32(node.Y)-5, 30, color.RGBA{0, 0, 255, 255}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", i), node.X-10, node.Y-10)
	}

	// Draw edges
	for _, node := range s.graph.Nodes {
		for _, adj := range node.AdjList {
			adjNode := s.graph.Nodes[adj]
			ebitenutil.DrawLine(screen, float64(node.X), float64(node.Y), float64(adjNode.X), float64(adjNode.Y), color.White)
		}
	}
}

func (s *MapScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cons.ScreenWidth, cons.ScreenHeight
}

func isPointInCircle(px, py, cx, cy, radius int) bool {
	dx := px - cx
	dy := py - cy
	return dx*dx+dy*dy <= radius*radius
}
