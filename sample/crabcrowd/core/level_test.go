package core

import "testing"

func TestParseLevel_PanicsOnInconsistentRowWidth(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected ParseLevel to panic on a jagged (non-rectangular) map")
		}
	}()
	ParseLevel("bad", []string{
		"#####",
		"#.T#", // 少了一个字符，长度和首行不一致
		"#.C.#",
		"#####",
	})
}

func TestParseLevel_CrabOnTargetSymbolX(t *testing.T) {
	lv := ParseLevel("x-test", []string{
		"###",
		"#X#",
		"###",
	})
	if len(lv.Crabs) != 1 {
		t.Fatalf("expected exactly one crab, got %d", len(lv.Crabs))
	}
	if got := lv.CellAt(lv.Crabs[0]); got != Target {
		t.Fatalf("expected the crab's starting cell to be a Target, got %v", got)
	}
	if IsWin(&lv, lv.Crabs) == false {
		t.Fatalf("a crab starting directly on the only target should already count as won")
	}
}

func TestLevel_RenderASCII(t *testing.T) {
	lv := gridLevel([]string{
		"###",
		"#T#",
		"#C#",
		"###",
	})
	got := lv.RenderASCII(lv.Crabs)
	want := "###\n#T#\n#C#\n###\n"
	if got != want {
		t.Fatalf("RenderASCII mismatch:\ngot:\n%s\nwant:\n%s", got, want)
	}

	moved, _ := ResolveMove(lv, lv.Crabs, Up)
	got2 := lv.RenderASCII(moved)
	want2 := "###\n#x#\n#.#\n###\n"
	if got2 != want2 {
		t.Fatalf("RenderASCII after move mismatch:\ngot:\n%s\nwant:\n%s", got2, want2)
	}
}
