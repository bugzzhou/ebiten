package mapscene

import (
	"ebiten/scene/models"
	"ebiten/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CheckNodeSite(nodes []models.Node) (string, int) {
	nodeType := ""
	nodeIndex := -1
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, node := range nodes {
			if IsPointInCircle(x, y, node.X, node.Y, 30) && roomIndexIsAbleToChoose(nodes, i) {
				nodeType = node.NodeType
				nodeIndex = i
				break
			}
		}
	}
	return nodeType, nodeIndex
}

func roomIndexIsAbleToChoose(rooms []models.Node, index int) bool {
	if LastChooseRoomIndex == -1 {
		return index == 0
	}
	if len(rooms) < LastChooseRoomIndex {
		return false
	}
	return utils.IsInSlice(rooms[LastChooseRoomIndex].AdjList, index)
}
