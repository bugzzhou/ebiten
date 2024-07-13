package game

import (
	"fmt"
	"testing"
)

func TestGetHandXY(t *testing.T) {
	res := getHandcardXYs(5)
	fmt.Printf("res are: %v\n", res)
}
