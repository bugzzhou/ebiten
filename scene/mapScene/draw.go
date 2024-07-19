package mapscene

import (
	"ebiten/scene/models"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawMapNode(screen *ebiten.Image, nodes []models.Node) {
	// Draw nodes
	for i, node := range nodes {
		vector.DrawFilledCircle(screen, float32(node.X)-5, float32(node.Y)-5, 30, color.RGBA{0, 0, 255, 255}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", i), node.X-10, node.Y-10)
	}

	// Draw edges
	for _, node := range nodes {
		for _, adj := range node.AdjList {
			adjNode := nodes[adj]
			vector.StrokeLine(screen, float32(node.X), float32(node.Y), float32(adjNode.X), float32(adjNode.Y), 1, color.White, false)
		}
	}
}
