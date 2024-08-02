package main

import (
	"clock/utils.go"
	"fmt"
)

func main() {
	d := utils.Digits

	result := utils.ConcatArrays(d[1], d[2], d[10], d[4], d[5], d[10], d[7], d[8])

	for _, row := range result {
		fmt.Println(row)
	}
}
