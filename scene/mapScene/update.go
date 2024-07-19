package mapscene

import (
	"ebiten/scene/models"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CheckNodeSite(nodes []models.Node) int {
	nodeType := -1
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, node := range nodes {
			if IsPointInCircle(x, y, node.X, node.Y, 30) {
				fmt.Printf("Clicked node %d\n", i)
				nodeType = node.NodeType
				break
			}
		}
	}
	return nodeType
}
