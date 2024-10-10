package main

import (
	"fmt"
	"math"
)

func main() {
	var k, fk float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	fk = math.Pow(4*k+2, 2) / ((4*k + 1) * (4*k + 3))
	fmt.Printf("nilai f(k) = %.10f", fk)
}
