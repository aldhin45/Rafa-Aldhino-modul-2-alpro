package main

import (
    "fmt"
    "math"
)

func main() {
    var radius int

    fmt.Print("Jari-jari: ")
    fmt.Scan(&radius)

    volume := (4 / 3) * math.Pi * math.Pow(float64(radius), 3)
    surfaceArea := 4 * math.Pi * math.Pow(float64(radius), 2)

    fmt.Printf("Bola dengan jari-jari %d memiliki volume %.6f dan luas kulit %.6f\n", radius, volume, surfaceArea)
}