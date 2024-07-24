package mapscene

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	a := gensimplePathMap()

	for _, v := range a {
		fmt.Printf("v is: %v\n", v)
	}

}
