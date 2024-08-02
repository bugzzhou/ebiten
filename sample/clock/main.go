package main

import (
	"clock/utils.go"
	"fmt"
)

// ConcatArrays takes eight 2D arrays (A1, A2, B1, A3, A4, B2, A5, A6) and concatenates them into one large 2D array
//
//	12 : 34 : 56 //用于展示一个时间

func main() {
	d := utils.Digits

	result := utils.ConcatArrays(d[1], d[2], d[3], d[4], d[5], d[6], d[7], d[8])

	for _, row := range result {
		fmt.Println(row)
	}
}
