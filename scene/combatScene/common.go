package combatscene

import (
	"ebiten/scene/comm"
	_ "image/jpeg"
	_ "image/png"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
	imageWidth   = 150
	imageHeight  = 200
)

// TODO bugzzhou
// 其实可以直接融合到CombatScene中去，暂时保留，后续优化结构时，统一修改
type Game struct {
	Character *comm.Character

	Enemy comm.Enemy

	RoundBegin bool

	ExpandIndex   int
	DraggingIndex int
	IsDragging    bool

	Round int //回合数，后续用于计算buff的生效数值
}
