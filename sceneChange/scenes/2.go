package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Scene2 struct {
	manager *SceneManager
}

func NewScene2(manager *SceneManager) *Scene2 {
	return &Scene2{manager: manager}
}

func (s *Scene2) Update() error {
	update1()
	update2()
	s.clickChange()
	return nil
}

func (s *Scene2) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Scene 2")
	draw1(screen)
	draw2(screen)
	draw3(screen)
	draw4(screen)
}

func (s *Scene2) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (s *Scene2) clickChange() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		s.manager.SetScene(NewScene3(s.manager))
	}
}

func update1() {}
func update2() {}

func draw1(screen *ebiten.Image) {}
func draw2(screen *ebiten.Image) {}
func draw3(screen *ebiten.Image) {}
func draw4(screen *ebiten.Image) {}
