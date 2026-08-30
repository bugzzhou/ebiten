package core

import "testing"

func TestSolver_SolvesPresetLevels_Optimally(t *testing.T) {
	// 这些是我们手工验证过的可行解（未必最优），用来交叉校验：
	// BFS 求解器给出的最短解步数不应该比手工解更长。
	manualSolutionLen := []int{1, 2, 3, 2, 3}

	for idx := range Levels {
		lv := Levels[idx]
		moves, ok := Solve(&lv)
		if !ok {
			t.Fatalf("%s: 求解器未能找到解", lv.Name)
		}
		if len(moves) == 0 {
			t.Fatalf("%s: 不应该在还没开始移动就处于胜利状态", lv.Name)
		}
		if len(moves) > manualSolutionLen[idx] {
			t.Errorf("%s:求解器给出的解(%d步)比手工验证的解(%d步)更长，说明求解器不是最优的",
				lv.Name, len(moves), manualSolutionLen[idx])
		}

		// 交叉验证：把求解器给出的移动序列真正跑一遍，必须能通关。
		g := NewGameState(&lv)
		for _, d := range moves {
			g.Move(d)
		}
		if !g.Won() {
			t.Errorf("%s: 按求解器给出的移动序列执行后未能通关", lv.Name)
		}
	}
}

func TestSolve_AlreadyWonLevel(t *testing.T) {
	level := gridLevel([]string{
		"###",
		"#X#",
		"###",
	})
	moves, ok := Solve(level)
	if !ok {
		t.Fatalf("初始已胜利的关卡应该被判定为可解")
	}
	if len(moves) != 0 {
		t.Fatalf("初始已胜利的关卡不应该需要任何移动，got %v", moves)
	}
}

func TestSolve_UnsolvableLevel(t *testing.T) {
	// 目标点被整整一排墙彻底封死在另一个房间里，任何移动序列都无法到达。
	level := gridLevel([]string{
		"#####",
		"#C..#",
		"#####",
		"#.T.#",
		"#####",
	})
	_, ok := Solve(level)
	if ok {
		t.Fatalf("这个关卡应该是无解的（目标点所在房间完全没有出入口）")
	}
}

func TestSolveWithLimit_ExhaustedWhenCapTooLow(t *testing.T) {
	level := buildLongCorridor(210)
	// 把目标点放在离螃蟹很远的地方，同时给一个极小的状态上限，逼迫求解器提前放弃。
	level.Terrain[1][len(level.Terrain[1])-2] = Target
	_, solvable, exhausted := SolveWithLimit(level, 1)
	if solvable {
		t.Fatalf("状态上限极小时不应该声称已经确定可解")
	}
	if !exhausted {
		t.Fatalf("应该报告 exhausted=true 表示搜索被提前截断")
	}
}
