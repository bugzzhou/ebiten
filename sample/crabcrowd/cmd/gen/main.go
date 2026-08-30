// Command gen 用于批量随机生成一批"保证可解"的螃蟹关卡，
// 方便在没有现成关卡库可下载的情况下，自己"刷"出一批可用的关卡。
//
// 用法示例：
//
//	go run ./cmd/gen -count=5 -width=9 -height=9 -crabs=3 -minlen=6 -maxlen=20
package main

import (
	"crabcrowd/core"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	count := flag.Int("count", 5, "生成关卡的数量")
	width := flag.Int("width", 9, "地图宽度（含四周墙壁）")
	height := flag.Int("height", 9, "地图高度（含四周墙壁）")
	crabs := flag.Int("crabs", 3, "螃蟹（=目标点）数量")
	wallDensity := flag.Float64("wall", 0.15, "内部格子变成墙的概率 [0,1)")
	minLen := flag.Int("minlen", 3, "期望的最少解题步数")
	maxLen := flag.Int("maxlen", 25, "期望的最多解题步数")
	attempts := flag.Int("attempts", 500, "每一关最多随机尝试多少次")
	seed := flag.Int64("seed", 0, "随机种子；0 表示每次运行都不同")
	flag.Parse()

	var rng *rand.Rand
	if *seed != 0 {
		rng = rand.New(rand.NewSource(*seed))
	}

	produced := 0
	for produced < *count {
		cfg := core.GenConfig{
			Width:           *width,
			Height:          *height,
			NumCrabs:        *crabs,
			WallDensity:     *wallDensity,
			MinSolutionLen:  *minLen,
			MaxSolutionLen:  *maxLen,
			MaxAttempts:     *attempts,
			MaxSearchStates: 300000,
			Rand:            rng,
		}
		level, moves, ok := core.GenerateLevel(cfg)
		if !ok {
			fmt.Fprintf(os.Stderr, "第 %d 关生成失败（参数太苛刻或运气不好），已跳过\n", produced+1)
			produced++
			continue
		}

		produced++
		fmt.Printf("===== 关卡 %d/%d（最短%d 步可通关） =====\n", produced, *count, len(moves))
		fmt.Print(level.RenderASCII(level.Crabs))
		fmt.Printf("参考解: %s\n\n", formatMoves(moves))
	}
}

func formatMoves(moves []core.Direction) string {
	names := make([]string, len(moves))
	for i, d := range moves {
		switch d {
		case core.Up:
			names[i] = "W"
		case core.Down:
			names[i] = "S"
		case core.Left:
			names[i] = "A"
		case core.Right:
			names[i] = "D"
		}
	}
	return strings.Join(names, " ")
}
