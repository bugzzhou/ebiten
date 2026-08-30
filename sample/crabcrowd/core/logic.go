package core

// ResolveMove 计算一次全局移动的结算结果。
//
// 规则（顶推链式）：沿移动方向看，每只螃蟹是否能移动取决于它前方那一格——
//   - 前方是墙（或越界）：不能移动。
//   - 前方是另一只螃蟹：取决于那只螃蟹本次是否能移动（递归判定，形成连锁）。
//   - 前方是空地/目标点且没有螃蟹：可以移动。
//
// 返回：移动结算后的新坐标集合（与输入 crabs 一一对应，顺序不变），
// 以及本次是否至少有一只螃蟹发生了实际位移。
func ResolveMove(level *Level, crabs []Point, dir Direction) ([]Point, bool) {
	delta := dir.Delta()
	n := len(crabs)

	posIndex := make(map[Point]int, n)
	for i, p := range crabs {
		posIndex[p] = i
	}

	visited := make([]bool, n)
	canMove := make([]bool, n)

	var resolve func(i int) bool
	resolve = func(i int) bool {
		if visited[i] {
			return canMove[i]
		}
		visited[i] = true // 提前标记，防御性避免异常环形依赖导致死循环

		next := crabs[i].Add(delta)
		if level.IsWall(next) {
			canMove[i] = false
			return false
		}
		if j, ok := posIndex[next]; ok {
			canMove[i] = resolve(j)
			return canMove[i]
		}
		canMove[i] = true
		return true
	}

	newCrabs := make([]Point, n)
	moved := false
	for i := range crabs {
		if resolve(i) {
			newCrabs[i] = crabs[i].Add(delta)
			moved = true
		} else {
			newCrabs[i] = crabs[i]
		}
	}
	return newCrabs, moved
}

// IsWin 判断当前螃蟹坐标集合是否使地图上所有目标点都被占据。
func IsWin(level *Level, crabs []Point) bool {
	occupied := make(map[Point]bool, len(crabs))
	for _, p := range crabs {
		occupied[p] = true
	}
	for _, t := range level.Targets() {
		if !occupied[t] {
			return false
		}
	}
	return true
}
