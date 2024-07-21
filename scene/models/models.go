package models

//1
//
//1
//
//1
//mapScene 的结构体
type Node struct {
	Id       int
	NodeType string //表示房间的 类型
	X, Y     int
	AdjList  []int
	Explored bool
}

type Graph struct {
	Nodes []Node
}
