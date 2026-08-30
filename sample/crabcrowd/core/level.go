package core

import (
	"fmt"
	"strings"
)

// Level 表示一个关卡：静态地形 + 螃蟹初始位置。
// Terrain 的索引方式为 Terrain[y][x]。
type Level struct {
	Name    string
	Terrain [][]Cell
	Crabs   []Point
}

// Height 返回地图行数。
func (l *Level) Height() int {
	return len(l.Terrain)
}

// Width 返回地图列数。
func (l *Level) Width() int {
	if len(l.Terrain) == 0 {
		return 0
	}
	return len(l.Terrain[0])
}

// InBounds 判断坐标是否在地图范围内。
func (l *Level) InBounds(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < l.Height() && p.X < l.Width()
}

// CellAt 返回指定坐标的静态地形；超出边界视为墙。
func (l *Level) CellAt(p Point) Cell {
	if !l.InBounds(p) {
		return Wall
	}
	return l.Terrain[p.Y][p.X]
}

// IsWall 判断指定坐标是否为墙（或越界）。
func (l *Level) IsWall(p Point) bool {
	return l.CellAt(p) == Wall
}

// Targets 返回地图上所有目标点坐标。
func (l *Level) Targets() []Point {
	var targets []Point
	for y, row := range l.Terrain {
		for x, c := range row {
			if c == Target {
				targets = append(targets, Point{X: x, Y: y})
			}
		}
	}
	return targets
}

// ParseLevel 通过字符画（每个字符串代表一行）构建关卡。
//
// 符号约定：
//
//	'#' 墙
//	'.' 空地
//	'T' 目标点（初始为空）
//	'C' 螃蟹（初始站在空地上）
//	'X' 螃蟹（初始站在目标点上）
//
// 要求所有行长度一致（否则 panic，尽早暴露关卡数据录入错误），
// 且目标点数量与螃蟹数量相等（该约束由调用方保证/在测试中校验）。
func ParseLevel(name string, rows []string) Level {
	terrain := make([][]Cell, len(rows))
	var crabs []Point
	var width int
	for y, row := range rows {
		runes := []rune(row)
		if y == 0 {
			width = len(runes)
		} else if len(runes) != width {
			panic(fmt.Sprintf("ParseLevel(%q):第 %d 行长度为 %d，与首行长度 %d 不一致，关卡地图必须是矩形", name, y, len(runes), width))
		}
		line := make([]Cell, len(runes))
		for x, ch := range runes {
			switch ch {
			case '#':
				line[x] = Wall
			case 'T':
				line[x] = Target
			case 'C':
				line[x] = Floor
				crabs = append(crabs, Point{X: x, Y: y})
			case 'X':
				line[x] = Target
				crabs = append(crabs, Point{X: x, Y: y})
			default: // '.' 以及其他未知字符统一视为空地
				line[x] = Floor
			}
		}
		terrain[y] = line
	}
	return Level{Name: name, Terrain: terrain, Crabs: crabs}
}

// RenderASCII 把关卡地形与当前螃蟹坐标渲染成便于阅读的字符画，
// 主要用于控制台调试/CLI 工具输出，符号约定与 ParseLevel 保持一致
// （另外用小写 'x' 表示"螃蟹已经站在目标点上"，与解析用的大写 'X' 区分，避免歧义）。
func (l *Level) RenderASCII(crabs []Point) string {
	occupied := make(map[Point]bool, len(crabs))
	for _, p := range crabs {
		occupied[p] = true
	}

	var b strings.Builder
	for y := 0; y < l.Height(); y++ {
		for x := 0; x < l.Width(); x++ {
			p := Point{X: x, Y: y}
			hasCrab := occupied[p]
			switch {
			case l.CellAt(p) == Wall:
				b.WriteByte('#')
			case hasCrab && l.CellAt(p) == Target:
				b.WriteByte('x')
			case hasCrab:
				b.WriteByte('C')
			case l.CellAt(p) == Target:
				b.WriteByte('T')
			default:
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}
