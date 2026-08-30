package game

import (
	"crabcrowd/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 窗口/网格相关的常量。CellSize 决定每个格子的像素大小；
// hudTop/hudBottom 是顶部信息栏与底部操作提示栏预留的空间，
// 地图会在剩余区域内居中显示，从而适配不同大小的关卡。
const (
	CellSize  = 64
	ScreenW   = 800
	ScreenH   = 720
	hudTop    = 90
	hudBottom = 60
)

// Game 实现 ebiten.Game 接口，负责输入处理与渲染调度；
// 具体的移动/撤销/胜负判定全部委托给 core 包的纯逻辑实现。
type Game struct {
	levelIdx int
	state    *core.GameState
}

// NewGame 创建一局新游戏，默认从第 1 关开始。
func NewGame() (*Game, error) {
	g := &Game{}
	g.loadLevel(0)
	return g, nil
}

func (g *Game) loadLevel(idx int) {
	if idx < 0 {
		idx = 0
	}
	if idx >= len(core.Levels) {
		idx = len(core.Levels) - 1
	}
	g.levelIdx = idx
	g.state = core.NewGameState(&core.Levels[idx])
}

func (g *Game) Update() error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		g.state.Move(core.Up)
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		g.state.Move(core.Down)
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		g.state.Move(core.Left)
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		g.state.Move(core.Right)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		g.state.Undo()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.state.Reset()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.loadLevel(g.levelIdx + 1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.loadLevel(g.levelIdx - 1)
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenW, ScreenH
}
