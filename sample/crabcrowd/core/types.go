package core

// Point 表示地图上的一个坐标（X 为列，Y 为行）。
type Point struct {
	X, Y int
}

// Add 返回 p 与 o 相加后的新坐标。
func (p Point) Add(o Point) Point {
	return Point{X: p.X + o.X, Y: p.Y + o.Y}
}

// Direction 表示一次全局移动指令的方向。
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Delta 返回该方向对应的坐标增量。
func (d Direction) Delta() Point {
	switch d {
	case Up:
		return Point{X: 0, Y: -1}
	case Down:
		return Point{X: 0, Y: 1}
	case Left:
		return Point{X: -1, Y: 0}
	case Right:
		return Point{X: 1, Y: 0}
	default:
		return Point{X: 0, Y: 0}
	}
}

// Cell 表示地图上的静态地形，不包含螃蟹（螃蟹位置单独维护）。
type Cell int

const (
	Floor  Cell = iota // 空地
	Wall               // 墙，不可通行
	Target             // 目标点（本质是特殊的空地）
)
