package main

import (
	"fmt"
	"math"
	"testing"
)

func TestRaizQuadrada(t *testing.T) {

	x := 0.0001
	for i := 0; i <= 100; i++ {
		x = x + math.Sqrt(x)
		fmt.Printf("%v", x)

		if i > 100 {
			t.Error("Resultado invalido")
		}
	}
}
