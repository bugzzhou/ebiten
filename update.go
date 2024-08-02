package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

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
		if stepButton.IsClicked(x, y) {
			updateGrid() // 单次迭代
		} else if runButton.IsClicked(x, y) {
			running = true
			ticker = time.NewTicker(100 * time.Millisecond) // 每秒钟迭代一次
		} else if stopButton.IsClicked(x, y) {
			running = false
			if ticker != nil {
				ticker.Stop()
			}
		}
	}

	mousePressedLastFrame = mousePressed
}

var (
	topPotential    int // 上边界可能从死变为生的方格数量
	bottomPotential int // 下边界
	leftPotential   int // 左边界
	rightPotential  int // 右边界
)

func updateGrid() {
	newGrid := make([][]bool, gridSize)
	for i := range newGrid {
		newGrid[i] = make([]bool, gridSize)
	}

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			num := countNeighbors(i, j)
			if !grid[i][j] && num == 3 {
				newGrid[i][j] = true
			} else if grid[i][j] {
				if num == 2 || num == 3 {
					newGrid[i][j] = true
				} else {
					newGrid[i][j] = false
				}
			}
		}
	}

	for i := 0; i < gridSize; i++ {
		if countNeighbors(-1, i) == 3 {
			topPotential++
		}
		if countNeighbors(gridSize, i) == 3 {
			bottomPotential++
		}
		if countNeighbors(i, -1) == 3 {
			leftPotential++
		}
		if countNeighbors(i, gridSize) == 3 {
			rightPotential++
		}
	}

	fmt.Printf("Top Potential: %d\n", topPotential)
	fmt.Printf("Bottom Potential: %d\n", bottomPotential)
	fmt.Printf("Left Potential: %d\n", leftPotential)
	fmt.Printf("Right Potential: %d\n", rightPotential)

	grid = newGrid
}

func countNeighbors(row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			tRow := row + i
			tCol := col + j
			if tRow == row && tCol == col {
				continue
			}
			if tRow >= 0 && tRow < gridSize && tCol >= 0 && tCol < gridSize {
				if grid[tRow][tCol] {
					count++
				}
			}
		}
	}
	return count
}

// func updateGrid() {
// 	newGrid := make([][]bool, gridSize)
// 	for i := range newGrid {
// 		newGrid[i] = make([]bool, gridSize)
// 	}

// 	for i := 0; i < gridSize; i++ {
// 		for j := 0; j < gridSize; j++ {
// 			num := countNeighbors(i, j)
// 			if !grid[i][j] && num == 3 {
// 				newGrid[i][j] = true
// 			} else if grid[i][j] {
// 				if num == 2 || num == 3 {
// 					newGrid[i][j] = true
// 				} else {
// 					newGrid[i][j] = false
// 				}
// 			}
// 		}
// 	}

// 	grid = newGrid
// }

// func countNeighbors(row, col int) int {
// 	count := 0
// 	for i := -1; i <= 1; i++ {
// 		for j := -1; j <= 1; j++ {
// 			tRow := row + i
// 			tCol := col + j
// 			if tRow == row && tCol == col {
// 				continue
// 			}
// 			if tRow >= 0 && tRow < gridSize && tCol >= 0 && tCol < gridSize {
// 				if grid[tRow][tCol] {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
