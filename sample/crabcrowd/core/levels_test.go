package core

import "testing"

// TestPresetLevels_TargetsMatchCrabCount 保证每个预设关卡目标点数量与螃蟹数量一致。
func TestPresetLevels_TargetsMatchCrabCount(t *testing.T) {
	for _, lv := range Levels {
		lv := lv
		if got, want := len(lv.Targets()), len(lv.Crabs); got != want {
			t.Errorf("%s: 目标点数量(%d) != 螃蟹数量(%d)", lv.Name, got, want)
		}
	}
}

// TestPresetLevels_Solvable 针对每个预设关卡给出一套已验证可行的移动序列，
// 断言最终一定能进入胜利状态（同时验证胜利判定不会提前误判）。
func TestPresetLevels_Solvable(t *testing.T) {
	solutions := [][]Direction{
		{Up},         // 第1关
		{Up, Up},     // 第2关
		{Up, Up, Up}, // 第3关·顶推链
		{Up, Up},     // 第4关 · 内部隔墙
		{Up, Up, Up}, // 第5关 · 里外兼修
	}

	if len(solutions) != len(Levels) {
		t.Fatalf("测试用例数量(%d)与预设关卡数量(%d)不一致，请同步更新", len(solutions), len(Levels))
	}

	for idx, lv := range Levels {
		lv := lv
		g := NewGameState(&lv)
		if g.Won() {
			t.Fatalf("%s: 初始状态不应该直接胜利", lv.Name)
		}
		for _, dir := range solutions[idx] {
			g.Move(dir)
		}
		if !g.Won() {
			t.Errorf("%s: 按预定序列 %v 执行后未能通关，最终螃蟹位置=%v", lv.Name, solutions[idx], g.Crabs)
		}
	}
}
