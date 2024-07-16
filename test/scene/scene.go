package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene 接口，所有场景都需要实现
type Scene interface {
	Update(changeScene func(newScene Scene)) error
	Draw(screen *ebiten.Image)
}
