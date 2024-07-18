package models

import "github.com/hajimehoshi/ebiten/v2"

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
	action  []Act
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
