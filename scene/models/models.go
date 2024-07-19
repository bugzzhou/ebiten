package models

import "github.com/hajimehoshi/ebiten/v2"

//combatScene 的结构体
type Character struct {
	Image   *ebiten.Image
	Hp      int
	Hplimit int
	Energy  int
}

type Enemy struct {
	Image   *ebiten.Image
	Hp      int
	Hplimit int
	Action  []Act
	Buffs   []Buff
}

type Act struct {
	Id          int
	Name        string
	Description string
}

type Buff struct {
	Id          int
	Name        string
	Description string

	Layers     int
	StartRound int
	EndRound   int
}

//1
//
//1
//
//1
//mapScene 的结构体
type Node struct {
	NodeType int //表示房间的 类型
	X, Y     int
	AdjList  []int
}

type Graph struct {
	Nodes []Node
}
