package core

import "testing"

func gridLevel(rows []string) *Level {
	l := ParseLevel("test", rows)
	return &l
}

func TestResolveMove_SingleCrabMovesIntoFloor(t *testing.T) {
	level := gridLevel([]string{
		"###",
		"#.#",
		"#C#",
		"###",
	})
	crabs := clonePoints(level.Crabs)
	newCrabs, moved := ResolveMove(level, crabs, Up)
	if !moved {
		t.Fatalf("expected movement to happen")
	}
	want := Point{X: 1, Y: 1}
	if newCrabs[0] != want {
		t.Fatalf("got %+v, want %+v", newCrabs[0], want)
	}
}

func TestResolveMove_BlockedByWall(t *testing.T) {
	level := gridLevel([]string{
		"###",
		"#C#",
		"###",
	})
	crabs := clonePoints(level.Crabs)
	newCrabs, moved := ResolveMove(level, crabs, Up)
	if moved {
		t.Fatalf("expected no movement, crab is against the wall")
	}
	if newCrabs[0] != crabs[0] {
		t.Fatalf("crab position should stay unchanged when blocked")
	}
}

func TestResolveMove_ChainFollowsWhenFrontMoves(t *testing.T) {
	// B 在 A 上方一格，上方是空地：整体应该都往上移动一格。
	level := gridLevel([]string{
		"###",
		"#.#",
		"#C#", // B
		"#C#", // A
		"###",
	})
	crabs := clonePoints(level.Crabs) // crabs[0] = B(1,2), crabs[1] = A(1,3)
	newCrabs, moved := ResolveMove(level, crabs, Up)
	if !moved {
		t.Fatalf("expected chain movement")
	}
	if newCrabs[0] != (Point{X: 1, Y: 1}) {
		t.Fatalf("front crab should move to (1,1), got %+v", newCrabs[0])
	}
	if newCrabs[1] != (Point{X: 1, Y: 2}) {
		t.Fatalf("trailing crab should follow into (1,2), got %+v", newCrabs[1])
	}
}

func TestResolveMove_ChainBlockedWhenFrontHitsWall(t *testing.T) {
	// B紧贴墙壁，A 紧跟 B：整条链都不应该移动。
	level := gridLevel([]string{
		"###",
		"#C#", // B against wall
		"#C#", // A
		"###",
	})
	crabs := clonePoints(level.Crabs)
	newCrabs, moved := ResolveMove(level, crabs, Up)
	if moved {
		t.Fatalf("expected no movement, whole chain should be blocked")
	}
	if newCrabs[0] != crabs[0] || newCrabs[1] != crabs[1] {
		t.Fatalf("chain positions should stay unchanged when blocked")
	}
}

func TestResolveMove_OutOfBoundsTreatedAsWall(t *testing.T) {
	level := gridLevel([]string{
		"C.",
	})
	_, moved := ResolveMove(level, clonePoints(level.Crabs), Up)
	if moved {
		t.Fatalf("moving toward the map boundary should be blocked")
	}
}

func TestIsWin(t *testing.T) {
	level := gridLevel([]string{
		"#####",
		"#.T.#",
		"#.C.#",
		"#####",
	})
	crabs := clonePoints(level.Crabs)
	if IsWin(level, crabs) {
		t.Fatalf("should not be won at the initial state")
	}
	newCrabs, _ := ResolveMove(level, crabs, Up)
	if !IsWin(level, newCrabs) {
		t.Fatalf("should be won after moving the crab onto the target")
	}
}
