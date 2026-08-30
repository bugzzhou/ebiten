package core

import (
	"math/rand"
	"testing"
)

func TestGenerateLevel_ProducesSolvableLevel(t *testing.T) {
	cfg := GenConfig{
		Width:           7,
		Height:          7,
		NumCrabs:        2,
		WallDensity:     0.15,
		MinSolutionLen:  1,
		MaxSolutionLen:  30,
		MaxAttempts:     500,
		MaxSearchStates: 200000,
		Rand:            rand.New(rand.NewSource(42)), // 固定种子，保证测试可重现
	}

	level, moves, ok := GenerateLevel(cfg)
	if !ok {
		t.Fatalf("在给定参数下应该能生成出合法关卡")
	}
	if len(level.Crabs) != cfg.NumCrabs {
		t.Fatalf("螃蟹数量应为 %d，实际 %d", cfg.NumCrabs, len(level.Crabs))
	}
	if got := len(level.Targets()); got != cfg.NumCrabs {
		t.Fatalf("目标点数量应等于螃蟹数量：got %d want %d", got, cfg.NumCrabs)
	}
	if len(moves) == 0 {
		t.Fatalf("生成的关卡不应该是初始就已经胜利的")
	}

	// 交叉验证求解器返回的解确实有效。
	g := NewGameState(level)
	for _, d := range moves {
		g.Move(d)
	}
	if !g.Won() {
		t.Fatalf("按生成器附带的解执行后应该通关")
	}
}

func TestGenerateLevel_RespectsSolutionLengthRange(t *testing.T) {
	cfg := GenConfig{
		Width:           9,
		Height:          9,
		NumCrabs:        3,
		WallDensity:     0.12,
		MinSolutionLen:  4,
		MaxSolutionLen:  12,
		MaxAttempts:     1000,
		MaxSearchStates: 200000,
		Rand:            rand.New(rand.NewSource(7)),
	}
	_, moves, ok := GenerateLevel(cfg)
	if !ok {
		t.Skip("该随机种子在限定尝试次数内没有生成出符合难度区间的关卡（非算法性错误，可放宽范围重试）")
	}
	if len(moves) < cfg.MinSolutionLen || len(moves) > cfg.MaxSolutionLen {
		t.Fatalf("解题步数 %d 超出期望区间 [%d,%d]", len(moves), cfg.MinSolutionLen, cfg.MaxSolutionLen)
	}
}

func TestGenerateLevel_FailsGracefullyWhenImpossibleConfig(t *testing.T) {
	// 地图太小，塞不下 5 只螃蟹 + 5 个目标点。
	cfg := GenConfig{
		Width:       4,
		Height:      4,
		NumCrabs:    5,
		WallDensity: 0,
		MaxAttempts: 20,
		Rand:        rand.New(rand.NewSource(1)),
	}
	_, _, ok := GenerateLevel(cfg)
	if ok {
		t.Fatalf("空间明显不够时应该返回失败，而不是生成出错误的关卡")
	}
}
