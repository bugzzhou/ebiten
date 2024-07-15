package main

import "github.com/hajimehoshi/ebiten/v2"

func changeStatus() {
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	x, y := ebiten.CursorPosition()
	gridX, gridY := x/cellSize, y/cellSize

	if mousePressed {
		if gridX < gridSize && gridY < gridSize {
			point := Point{X: gridX, Y: gridY}
			if !changedCells[point] {
				grid[gridX][gridY] = !grid[gridX][gridY]
				changedCells[point] = true
			}
		}
	} else {
		changedCells = make(map[Point]bool) // 重置记录
	}

	if mousePressed && !mousePressedLastFrame {
		if startButton.IsClicked(x, y) {
			updateGrid() // 调用改变二维数组状态的函数
		}
	}

	mousePressedLastFrame = mousePressed
}

func (b *Button) IsClicked(x, y int) bool {
	return x >= b.x && x <= b.x+b.width && y >= b.y && y <= b.y+b.height
}

func updateGrid() {
	// 这里实现改变二维数组状态的逻辑
}
