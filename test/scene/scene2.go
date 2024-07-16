package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// SceneB 结构体
type SceneB struct{}

// Update 方法实现，处理逻辑
func (s *SceneB) Update(changeScene func(newScene Scene)) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		changeScene(&SceneA{}) // 切换回 SceneA
	}
	return nil
}

// Draw 方法实现，绘制内容
func (s *SceneB) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})
	ebitenutil.DebugPrint(screen, "This is Scene B\nClick to switch to Scene A")
}
