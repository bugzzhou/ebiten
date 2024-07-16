package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// SceneA 结构体
type SceneA struct{}

// Update 方法实现，处理逻辑
func (s *SceneA) Update(changeScene func(newScene Scene)) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		changeScene(&SceneB{}) // 切换到 SceneB
	}
	return nil
}

// Draw 方法实现，绘制内容
func (s *SceneA) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "This is Scene A\nClick to switch to Scene B")
}
