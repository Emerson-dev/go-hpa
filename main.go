package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x = x + math.Sqrt(x)
		fmt.Printf("Code.education Rocks!")
	}
}

