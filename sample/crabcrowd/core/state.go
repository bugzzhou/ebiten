package core

// MaxUndo 限制撤销栈的最大深度。
const MaxUndo = 100

// GameState 维护某一关的运行时状态：当前螃蟹位置、撤销栈、步数。
type GameState struct {
	Level     *Level
	Crabs     []Point
	Steps     int
	undoStack [][]Point
}

// NewGameState 基于关卡的初始配置创建一个新的运行状态。
func NewGameState(level *Level) *GameState {
	return &GameState{
		Level: level,
		Crabs: clonePoints(level.Crabs),
	}
}

func clonePoints(pts []Point) []Point {
	out := make([]Point, len(pts))
	copy(out, pts)
	return out
}

// Move 执行一次全局移动。
// 仅当至少一只螃蟹发生了实际位移时才会计入步数并压入撤销栈；
// 若本次按键因全部被阻挡而没有任何螃蟹移动，则视为无效操作，不计步、不可撤销。
// 返回值表示本次移动是否产生了实际效果。
func (g *GameState) Move(dir Direction) bool {
	newCrabs, moved := ResolveMove(g.Level, g.Crabs, dir)
	if !moved {
		return false
	}
	g.pushUndo(g.Crabs)
	g.Crabs = newCrabs
	g.Steps++
	return true
}

func (g *GameState) pushUndo(snapshot []Point) {
	g.undoStack = append(g.undoStack, clonePoints(snapshot))
	if len(g.undoStack) > MaxUndo {
		g.undoStack = g.undoStack[1:]
	}
}

// Undo 撤销上一步。栈为空时返回 false表示无法继续撤销。
func (g *GameState) Undo() bool {
	if len(g.undoStack) == 0 {
		return false
	}
	last := g.undoStack[len(g.undoStack)-1]
	g.undoStack = g.undoStack[:len(g.undoStack)-1]
	g.Crabs = last
	if g.Steps > 0 {
		g.Steps--
	}
	return true
}

// Reset 将当前状态重置为关卡初始配置，并清空撤销栈与步数。
func (g *GameState) Reset() {
	g.Crabs = clonePoints(g.Level.Crabs)
	g.undoStack = nil
	g.Steps = 0
}

// Won 判断当前状态是否已经胜利。
func (g *GameState) Won() bool {
	return IsWin(g.Level, g.Crabs)
}

// UndoDepth 返回当前撤销栈中的记录数量（主要用于测试/展示）。
func (g *GameState) UndoDepth() int {
	return len(g.undoStack)
}
