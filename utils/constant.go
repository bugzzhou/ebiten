package utils

import (
	"math/rand"
	"time"
)

// 用于存放各种常量、变量、路径
var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	CardDir = "./pic/cards"

	CardInfoPath = "./utils/cards.csv"

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

const (
	StartNode    = "start"
	CombatNode   = "combat"
	CampfireNode = "fire"
	RandomNode   = "random"
	EndNode      = "end"
	BossNode     = "boss"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)
