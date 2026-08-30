package core

import (
	"fmt"
	"sort"
	"strings"
)

// DefaultMaxSearchStates 是求解器默认允许展开的最大状态数量，
// 用于防止在异常巨大的关卡上无限展开导致内存爆炸。
const DefaultMaxSearchStates = 500000

// stateKey 把一组螃蟹坐标规约成一个可比较的字符串。
// 螃蟹彼此不可区分，游戏状态只由"哪些格子被占据"决定，
// 因此必须先排序再编码，否则同一局面会因为切片顺序不同被误判成不同状态。
func stateKey(crabs []Point) string {
	pts := make([]Point, len(crabs))
	copy(pts, crabs)
	sort.Slice(pts, func(i, j int) bool {
		if pts[i].Y != pts[j].Y {
			return pts[i].Y < pts[j].Y
		}
		return pts[i].X < pts[j].X
	})
	var b strings.Builder
	for _, p := range pts {
		fmt.Fprintf(&b, "%d,%d|", p.X, p.Y)
	}
	return b.String()
}

type solverNode struct {
	crabs   []Point
	prevKey string
	dir     Direction
	hasPrev bool
}

// Solve 在状态空间中做无权最短路搜索（BFS），找到从关卡初始布局到
// "所有目标点都被占据"的最少步数移动序列。
//
// 之所以直接用 BFS 就能得到最优解，是因为这个游戏与经典推箱子不同：
// 这里没有"玩家"需要额外走位去推某一个箱子的角落，每一次按键（上/下/左/右）
// 对应确定性的唯一结算结果，所以每个状态的出度最多为 4，
// 状态空间是一张简单的有向图，标准 BFS 即可保证求得的路径步数最少。
func Solve(level *Level) ([]Direction, bool) {
	moves, ok, _ := SolveWithLimit(level, DefaultMaxSearchStates)
	return moves, ok
}

// SolveWithLimit 同 Solve，允许指定最多展开的状态数量上限。
//
// 返回值：
//   - moves: 找到的最短移动序列（找不到解时为 nil）
//   - solvable: 是否确定可解
//   - exhausted: 是否因为触达状态数上限而被迫提前放弃
//     （此时 solvable 一定为 false，但并不代表关卡真的无解，只是没搜完）
func SolveWithLimit(level *Level, maxStates int) (moves []Direction, solvable bool, exhausted bool) {
	start := clonePoints(level.Crabs)
	startKey := stateKey(start)

	if IsWin(level, start) {
		return nil, true, false
	}

	type queueItem struct {
		crabs []Point
		key   string
	}

	visited := map[string]solverNode{startKey: {crabs: start}}
	queue := []queueItem{{crabs: start, key: startKey}}
	dirs := [4]Direction{Up, Down, Left, Right}

	for len(queue) > 0 {
		if maxStates > 0 && len(visited) > maxStates {
			return nil, false, true
		}
		cur := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			next, moved := ResolveMove(level, cur.crabs, d)
			if !moved {
				continue
			}
			nk := stateKey(next)
			if _, seen := visited[nk]; seen {
				continue
			}
			visited[nk] = solverNode{crabs: next, prevKey: cur.key, dir: d, hasPrev: true}
			if IsWin(level, next) {
				return reconstructPath(visited, nk), true, false
			}
			queue = append(queue, queueItem{crabs: next, key: nk})
		}
	}
	return nil, false, false
}

func reconstructPath(visited map[string]solverNode, goalKey string) []Direction {
	var rev []Direction
	k := goalKey
	for {
		node := visited[k]
		if !node.hasPrev {
			break
		}
		rev = append(rev, node.dir)
		k = node.prevKey
	}
	moves := make([]Direction, len(rev))
	for i, d := range rev {
		moves[len(rev)-1-i] = d
	}
	return moves
}
