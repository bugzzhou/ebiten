package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	GridSize  = 100
	SharkSize = 2
	FishSize  = 1
	NumSharks = 5
	NumFishes = 50
)

type Position struct {
	X, Y int
}

type Fish struct {
	Pos Position
}

type Shark struct {
	Pos Position
}

type Ocean struct {
	Grid   [GridSize][GridSize]int
	Fishes []Fish
	Sharks []Shark
}

func (o *Ocean) Initialize() {
	// 初始化鲨鱼位置
	for i := 0; i < NumSharks; i++ {
		x := rand.Intn(GridSize - SharkSize)
		y := rand.Intn(GridSize - SharkSize)
		o.Sharks = append(o.Sharks, Shark{Position{x, y}})
		// 占用2x2格
		for j := 0; j < SharkSize; j++ {
			for k := 0; k < SharkSize; k++ {
				o.Grid[x+j][y+k] = 2 // 2 表示鲨鱼
			}
		}
	}

	// 初始化鱼位置
	for i := 0; i < NumFishes; i++ {
		var x, y int
		for {
			x = rand.Intn(GridSize)
			y = rand.Intn(GridSize)
			if o.Grid[x][y] == 0 {
				break
			}
		}
		o.Fishes = append(o.Fishes, Fish{Position{x, y}})
		o.Grid[x][y] = 1 // 1 表示鱼
	}
}

func (o *Ocean) UpdateFishMovement() {
	for i := range o.Fishes {
		// 更新鱼的位置
		fish := &o.Fishes[i]
		o.Grid[fish.Pos.X][fish.Pos.Y] = 0 // 清除当前位置

		// 简单规则：随机移动
		newX := fish.Pos.X + rand.Intn(3) - 1 // -1, 0, 1
		newY := fish.Pos.Y + rand.Intn(3) - 1

		// 边界检查
		if newX < 0 {
			newX = 0
		} else if newX >= GridSize {
			newX = GridSize - 1
		}
		if newY < 0 {
			newY = 0
		} else if newY >= GridSize {
			newY = GridSize - 1
		}

		// 更新位置
		fish.Pos.X = newX
		fish.Pos.Y = newY
		o.Grid[newX][newY] = 1 // 更新新位置
	}
}

func (o *Ocean) UpdateSharkMovement() {
	for i := range o.Sharks {
		// 更新鲨鱼的位置
		shark := &o.Sharks[i]
		// 清除当前2x2方格
		for j := 0; j < SharkSize; j++ {
			for k := 0; k < SharkSize; k++ {
				o.Grid[shark.Pos.X+j][shark.Pos.Y+k] = 0
			}
		}

		// 简单规则：随机移动
		newX := shark.Pos.X + rand.Intn(3) - 1 // -1, 0, 1
		newY := shark.Pos.Y + rand.Intn(3) - 1

		// 边界检查
		if newX < 0 {
			newX = 0
		} else if newX+SharkSize >= GridSize {
			newX = GridSize - SharkSize
		}
		if newY < 0 {
			newY = 0
		} else if newY+SharkSize >= GridSize {
			newY = GridSize - SharkSize
		}

		// 更新位置
		shark.Pos.X = newX
		shark.Pos.Y = newY
		for j := 0; j < SharkSize; j++ {
			for k := 0; k < SharkSize; k++ {
				o.Grid[newX+j][newY+k] = 2
			}
		}
	}
}

func (o *Ocean) PrintGrid() {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if o.Grid[i][j] != 0 {
				fmt.Printf("⚪ ")
			} else {
				fmt.Printf("   ")
			}
			// fmt.Printf("%d ", o.Grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func fishRun() {
	rand.Seed(time.Now().UnixNano())

	ocean := Ocean{}
	ocean.Initialize()

	for t := 0; t < 10; t++ {
		fmt.Printf("Step %d:\n", t)
		ocean.PrintGrid()

		ocean.UpdateFishMovement()
		ocean.UpdateSharkMovement()

		time.Sleep(1 * time.Second)
	}
}
