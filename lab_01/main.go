package main

import (
	"fmt"
	"math/rand"
)

func f1(x float64) float64 {
	return 0.5*x + x*x
}

func f2(x float64) float64 {
	return 12/-x + x*x
}

func withIf(x float64) {
	var y float64
	if x > 100 {
		y = f1(x)
	} else {
		y = f2(x)
	}

	fmt.Printf("f(%v) = %v\n", x, y)
}

func withSwitch(x float64) {
	var y float64
	switch {
	case x > 100:
		y = f1(x)
	default:
		y = f2(x)
	}

	fmt.Printf("f(%v) = %v\n", x, y)
}

func main() {
	x := rand.Float64() * 200

	withSwitch(x)
	withIf(x)
}
