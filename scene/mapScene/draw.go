package mapscene

import (
	"ebiten/scene/models"
	"ebiten/utils"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// var (
// 	campfireIcon *ebiten.Image
// 	combatIcon   *ebiten.Image
// 	StartIcon    *ebiten.Image
// 	EndIcon      *ebiten.Image
// )

var mapIconMap = map[string]*ebiten.Image{}

var mapIconDir = "./pic/mapScene"

func init() {
	files, ids, err := utils.ListDir(mapIconDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}

	for i := range files {
		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		mapIconMap[ids[i]] = tmpImage
	}
	// fmt.Printf("mapIconMap is: %v\n", mapIconMap)
}

func DrawMapNode(screen *ebiten.Image, nodes []models.Node) {
	// Draw nodes
	for i, node := range nodes {
		drawNodeTypePic(screen, node.NodeType, float64(node.X), float64(node.Y))

		text := fmt.Sprintf("%d", i)
		textWidth := len(text) * 7 // 假设每个字符的宽度为7像素，可以根据实际字体调整
		textHeight := 12           // 假设文本高度为12像素，可以根据实际字体调整
		bgX := float32(node.X - 10)
		bgY := float32(node.Y - 10)
		bgWidth := float32(textWidth)
		bgHeight := float32(textHeight)
		bgColor := color.RGBA{0, 0, 0, 255}
		if node.Explored {
			bgColor = color.RGBA{255, 0, 0, 255}
		}

		vector.DrawFilledRect(screen, bgX, bgY, bgWidth, bgHeight, bgColor, false)

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

func drawNodeTypePic(screen *ebiten.Image, nodeType string, x1, y1 float64) {
	chaOpt := &ebiten.DrawImageOptions{}
	chaOpt.GeoM.Translate(x1, y1)
	// fmt.Printf("nodeType, chOpt is: %v, %v\n", nodeType, chaOpt)
	screen.DrawImage(mapIconMap[nodeType], chaOpt)
}
