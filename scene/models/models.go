package models

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
