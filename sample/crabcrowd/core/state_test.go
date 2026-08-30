package core

import "testing"

func TestGameState_MoveStepsAndWin(t *testing.T) {
	level := gridLevel([]string{
		"#####",
		"#.T.#",
		"#.C.#",
		"#####",
	})
	g := NewGameState(level)

	if g.Won() {
		t.Fatalf("should not start in a won state")
	}
	if !g.Move(Up) {
		t.Fatalf("expected the move to take effect")
	}
	if g.Steps != 1 {
		t.Fatalf("expected Steps=1, got %d", g.Steps)
	}
	if !g.Won() {
		t.Fatalf("expected to be won after moving onto the target")
	}
}

func TestGameState_NoOpMoveDoesNotCountOrPush(t *testing.T) {
	level := gridLevel([]string{
		"###",
		"#C#",
		"###",
	})
	g := NewGameState(level)
	if g.Move(Up) {
		t.Fatalf("move against a wall should report no effect")
	}
	if g.Steps != 0 {
		t.Fatalf("blocked move must not count as a step, got %d", g.Steps)
	}
	if g.UndoDepth() != 0 {
		t.Fatalf("blocked move must not push an undo snapshot")
	}
}

func TestGameState_UndoRestoresPreviousPosition(t *testing.T) {
	level := gridLevel([]string{
		"#####",
		"#...#",
		"#.C.#",
		"#####",
	})
	g := NewGameState(level)
	before := clonePoints(g.Crabs)

	g.Move(Up)
	if g.Steps != 1 {
		t.Fatalf("expected Steps=1 after one move")
	}
	if !g.Undo() {
		t.Fatalf("expected undo to succeed")
	}
	if g.Steps != 0 {
		t.Fatalf("expected Steps=0 after undo, got %d", g.Steps)
	}
	for i := range before {
		if g.Crabs[i] != before[i] {
			t.Fatalf("undo did not restore original position: got %+v want %+v", g.Crabs[i], before[i])
		}
	}
	if g.Undo() {
		t.Fatalf("undo stack should be empty now")
	}
}

func TestGameState_Reset(t *testing.T) {
	level := gridLevel([]string{
		"#####",
		"#...#",
		"#.C.#",
		"#####",
	})
	g := NewGameState(level)
	g.Move(Up)
	g.Move(Up)
	g.Reset()
	if g.Steps != 0 {
		t.Fatalf("expected Steps=0 after reset, got %d", g.Steps)
	}
	if g.UndoDepth() != 0 {
		t.Fatalf("expected empty undo stack after reset")
	}
	for i, p := range g.Crabs {
		if p != level.Crabs[i] {
			t.Fatalf("reset did not restore initial crab positions")
		}
	}
}

func TestGameState_UndoStackCappedAt100(t *testing.T) {
	level := buildLongCorridor(210)
	g := NewGameState(level)
	for i := 0; i < 150; i++ {
		g.Move(Right)
	}
	if g.UndoDepth() != MaxUndo {
		t.Fatalf("expected undo depth capped at %d, got %d", MaxUndo, g.UndoDepth())
	}
	if g.Steps != 150 {
		t.Fatalf("expected Steps=150, got %d", g.Steps)
	}
}

func repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	out := make([]byte, 0, len(s)*n)
	for i := 0; i < n; i++ {
		out = append(out, s...)
	}
	return string(out)
}

// buildLongCorridor 构造一个 1xwidth 的横向长通道（两端及上下都是墙），
// 螃蟹起始在最左侧空地，用于测试长距离连续移动与撤销栈上限。
func buildLongCorridor(width int) *Level {
	l := ParseLevel("corridor", []string{
		repeat("#", width),
		"#" + repeat(".", width-2) + "#",
		repeat("#", width),
	})
	l.Crabs = []Point{{X: 1, Y: 1}}
	return &l
}
