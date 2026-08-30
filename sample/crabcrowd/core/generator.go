package core

import "math/rand"

// GenConfig 描述随机生成关卡时的参数。
type GenConfig struct {
	Width, Height int // 地图宽高（含四周墙壁）
	NumCrabs      int // 螃蟹（同时也是目标点）数量

	WallDensity float64 // 内部格子变成墙的概率，取值 [0,1)

	MinSolutionLen int // 期望的最少解题步数（用于过滤"太简单"的关卡），0 表示不限制
	MaxSolutionLen int // 期望的最多解题步数（用于过滤"太复杂/搜索太慢"的关卡），0 表示不限制

	MaxAttempts     int // 最多尝试生成多少次随机布局
	MaxSearchStates int // 传给求解器的状态数上限，防止单次尝试卡死

	Rand *rand.Rand // 随机数源；为 nil 时内部会自动创建一个
}

// GenerateLevel 按照"随机生成候选布局 + 求解器验证"的方式产出一个保证可解的关卡。
//
// 这是工程上最简单、也最稳妥的关卡生成思路：
// 直接生成关卡不保证可解，但只要跑一遍求解器确认"确实存在到达胜利状态的路径"，
// 就能100% 保证这一关不是死局；再结合最短解步数，还能顺便控制关卡难度。
//
// 返回值：生成的关卡、该关卡的一个已验证最短解，以及是否成功（尝试次数耗尽仍未找到合适关卡时为 false）。
func GenerateLevel(cfg GenConfig) (*Level, []Direction, bool) {
	rng := cfg.Rand
	if rng == nil {
		rng = rand.New(rand.NewSource(defaultSeed()))
	}
	maxAttempts := cfg.MaxAttempts
	if maxAttempts <= 0 {
		maxAttempts = 200
	}
	maxStates := cfg.MaxSearchStates
	if maxStates <= 0 {
		maxStates = DefaultMaxSearchStates
	}

	for attempt := 0; attempt < maxAttempts; attempt++ {
		level := randomLevel(cfg, rng)
		if level == nil {
			continue
		}

		moves, solvable, exhausted := SolveWithLimit(level, maxStates)
		if !solvable || exhausted {
			continue
		}
		if len(moves) == 0 {
			continue // 初始就已经胜利，太没意思，跳过
		}
		if cfg.MinSolutionLen > 0 && len(moves) < cfg.MinSolutionLen {
			continue
		}
		if cfg.MaxSolutionLen > 0 && len(moves) > cfg.MaxSolutionLen {
			continue
		}
		return level, moves, true
	}
	return nil, nil, false
}

// randomLevel 生成一份候选关卡（不保证可解，需要外层用求解器验证）。
func randomLevel(cfg GenConfig, rng *rand.Rand) *Level {
	w, h := cfg.Width, cfg.Height
	if w < 3 || h < 3 || cfg.NumCrabs <= 0 {
		return nil
	}

	terrain := make([][]Cell, h)
	for y := 0; y < h; y++ {
		terrain[y] = make([]Cell, w)
		for x := 0; x < w; x++ {
			if x == 0 || y == 0 || x == w-1 || y == h-1 {
				terrain[y][x] = Wall
				continue
			}
			if rng.Float64() < cfg.WallDensity {
				terrain[y][x] = Wall
			} else {
				terrain[y][x] = Floor
			}
		}
	}

	var floorCells []Point
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if terrain[y][x] == Floor {
				floorCells = append(floorCells, Point{X: x, Y: y})
			}
		}
	}

	need := cfg.NumCrabs * 2
	if len(floorCells) < need {
		return nil
	}

	rng.Shuffle(len(floorCells), func(i, j int) {
		floorCells[i], floorCells[j] = floorCells[j], floorCells[i]
	})

	crabs := make([]Point, cfg.NumCrabs)
	copy(crabs, floorCells[:cfg.NumCrabs])
	targets := floorCells[cfg.NumCrabs : cfg.NumCrabs*2]
	for _, t := range targets {
		terrain[t.Y][t.X] = Target
	}

	return &Level{Name: "随机关卡", Terrain: terrain, Crabs: crabs}
}
