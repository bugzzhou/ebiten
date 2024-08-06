package utils

import (
	"time"
)

func changeClock() {
	d := Digits
	timeNow := getCurrentTimeInSixDigits()
	clockMap = ConcatArrays(d[timeNow[0]], d[timeNow[1]], d[10], d[timeNow[2]], d[timeNow[3]], d[10], d[timeNow[4]], d[timeNow[5]])
}

func getCurrentTimeInSixDigits() []int {
	currentTime := time.Now()
	hour, minute, second := currentTime.Clock()

	h1 := hour / 10 // 小时的十位
	h2 := hour % 10 // 小时的个位

	m1 := minute / 10 // 分钟的十位
	m2 := minute % 10 // 分钟的个位

	s1 := second / 10 // 秒的十位
	s2 := second % 10 // 秒的个位

	return []int{h1, h2, m1, m2, s1, s2}
}
