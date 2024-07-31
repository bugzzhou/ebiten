package campfirescene

import "ebiten/utils"

const (
	campfireWeight = 150
	campfireHeight = 150
)

type HandcardXY struct {
	X int
	Y int
}

func GetXY(flag int) (x, y float64) {
	switch flag {
	case utils.CampFileFlag:
		return 200, 200

	}

	return 0, 0

}

func GetXYRange(flag int) (x1, x2, y1, y2 float64) {
	switch flag {
	case utils.CampFileFlag:
		return 200, 200 + campfireWeight, 200, 200 + campfireHeight
	}

	return 0, 0, 0, 0
}

func GetXYRangeInt(flag int) (x1, x2, y1, y2 int) {
	switch flag {
	case utils.CampFileFlag:
		return 200, 200 + campfireWeight, 200, 200 + campfireHeight
	}

	return 0, 0, 0, 0
}
