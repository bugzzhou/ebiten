package utils

import (
	"time"

	"math/rand"
)

// 用于存放各种常量、变量、路径

var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	CardDir = "./pic/cards"

	Lieren = "./pic/lieren.jpg"
	Kaka   = "./pic/kaka.jpg"
	Rough  = "./pic/rough.jpg"

	TestEnemy = "./pic/test_enemy.jpg"
	Boss1     = "./pic/test_enemy.jpg" //TODO bugzzhou 替换boss图片

	MapIconDir = "./pic/mapScene"

	CampfirePic = "./pic/campfire.jpg"
)

const (
	KakaId  = iota
	Boss1Id = 50

	TestEnemyId = 9999
)

const (
	CampFileFlag = iota
)
